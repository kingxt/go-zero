package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestServiceAnnotation(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServiceBlock().Accept(visitor)
	}
	test(t, do, spec.Service{
		Name: "example-api",
		Groups: []spec.Group{
			{
				Annotation: spec.Annotation{
					Properties: map[string]string{
						"jwt":        "Foo",
						"group":      "foo/bar",
						"anotherKey": "anotherValue",
					},
				},
			},
		},
	}, false, `@server(
		jwt: Foo
		group: foo/bar
		anotherKey: anotherValue
	)
	service example-api {
	}
	`)
}

func TestServiceBody(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServiceBlock().Accept(visitor)
	}
	test(t, do, spec.Service{
		Name: "example-api",
		Groups: []spec.Group{
			{
				Routes: []spec.Route{
					{
						Annotation: spec.Annotation{
							Properties: map[string]string{
								"handler":    "fooHandler1",
								"anotherKey": "anotherValue",
							},
						},
						Method: "post",
						Path:   "/api/foo1",
						RequestType: spec.Type{
							Name: "SingleExample",
						},
						Docs:    []string{"foo1"},
						Handler: "fooHandler1",
					},
					{
						Method: "post",
						Path:   "/api/foo3/:id",
						ResponseType: spec.Type{
							Name: "SingleExample2",
						},
						Handler: "fooHandler3",
					},
					{
						Method: "post",
						Path:   "/api/foo-bar",
						ResponseType: spec.Type{
							Name: "SingleExample3",
						},
						Handler: "fooHandler4",
					},
				},
			},
		},
	}, false, `
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

			@handler fooHandler4
    		post /api/foo-bar returns (SingleExample3)
		}
	`)
}

func TestServerMeta(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ServerMeta().Accept(visitor)
	}
	test(t, do, ast.KVSpec{
		List: []ast.KV{
			{
				Key: ast.Token{
					Text: "jwt",
				},
				Value: ast.Token{
					Text: "Foo",
				},
			},
			{
				Key: ast.Token{
					Text: "group",
				},
				Value: ast.Token{
					Text: "foo/bar",
				},
			},
			{
				Key: ast.Token{
					Text: "key",
				},
				Value: ast.Token{
					Text: "value",
				},
			},
		},
	}, false, `@server(
		jwt: Foo
		group: foo/bar
		key: value
	)
	`)

}

func TestServiceBlock(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
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
