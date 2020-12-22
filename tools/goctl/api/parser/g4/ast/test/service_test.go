package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestServiceAnnotation(t *testing.T) {
	testServiceAnnotation(t, `@server(
		jwt: Foo
		group: foo/bar
		anotherKey: anotherValue
	)
	service example-api {
	}
	`, "jwt", "Foo")
}

func TestServiceBody(t *testing.T) {
	p := ast.NewParser(ast.WithErrorCallback("", func(err error) {
		assert.Nil(t, err)
	}))
	result, err := p.Accept(`
		service example-api {
			@doc(
				summary: "foo1"
			)
			@server(
				handler: fooHandler1
				anotherKey: anotherValue
			)
    		post /api/foo1 (SingleExample)

			@handler fooHandler3
    		post /api/foo3/:id returns (SingleExample2)
		}
	`, func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServiceBlock().Accept(visitor)
	})
	assert.Nil(t, err)

	group, ok := result.(spec.Group)
	assert.True(t, ok)
	assert.Equal(t, len(group.Routes), 2)
	assert.Equal(t, group.Routes[0].Path, "/api/foo1")
	assert.Equal(t, group.Routes[0].Method, "post")

	assert.Equal(t, group.Routes[1].Handler, "fooHandler3")
}

func testServiceAnnotation(t *testing.T, content, key, value string) {
	p := ast.NewParser(ast.WithErrorCallback("", func(err error) {
		assert.Nil(t, err)
	}))
	result, err := p.Accept(content, func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServiceBlock().Accept(visitor)
	})
	assert.Nil(t, err)
	group, ok := result.(spec.Group)
	assert.True(t, ok)
	assert.Equal(t, group.Annotation.Properties[key], value)
}

func TestServerMeta(t *testing.T) {
	do := func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServerMeta().Accept(visitor)
	}
	test(t, do, spec.Annotation{
		Properties: map[string]string{
			"jwt":   "Foo",
			"group": "foo/bar",
			"key":   "value",
		},
	}, false, `@server(
		jwt: Foo
		group: foo/bar
		key: value
	)
	service user-api{}
	`)

	test(t, do, nil, true, `@server(
		jwt: Foo
		jwt: Bar
		group: foo/bar
		key: value
	)
	service foo-api{}
	`)

}

func TestServiceBlock(t *testing.T) {
	do := func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServiceBlock().Accept(visitor)
	}

	test(t, do, nil, true, `@server(
		jwt: Foo
		group: foo/bar
		key: value
	)
	service foo-api{
		@handler foo
		post /foo
		
		@handler foo
		post /bar
	}
	`)

	test(t, do, spec.Service{
		Name: "foo-api",
		Groups: []spec.Group{
			{
				Annotation: spec.Annotation{Properties: map[string]string{
					"jwt":   "Foo",
					"group": "foo/bar",
					"key":   "value",
				}},
				Routes: []spec.Route{
					{
						Method:  "post",
						Path:    "/foo",
						Handler: "foo",
					},
					{
						Method:  "get",
						Path:    "/foo",
						Handler: "bar",
					},
				},
			},
		},
	}, false, `@server(
		jwt: Foo
		group: foo/bar
		key: value
	)
	service foo-api{
		@handler foo
		post /foo
		
		@handler bar
		get /foo
	}
	`)

	test(t, do, nil, true, `@server(
		jwt: Foo
		group: foo/bar
		key: value
	)
	service foo-api{
		@handler foo
		post /foo
		
		@handler bar
		post /foo
	}
	`)
}
