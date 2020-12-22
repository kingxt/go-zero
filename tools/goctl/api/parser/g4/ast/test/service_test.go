package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestServiceAnnotation(t *testing.T) {
	testServiceAnnotation(t, `@server(
		jwt: Foo
		group: foo/bar
		anotherKey: anotherValue
	)
	`, "jwt", "Foo")
}

func TestServiceBody(t *testing.T) {
	p := ast.NewParser(`
		service example-api {
			@doc(
				summary: "foo1"
			)
			@server(
				handler: fooHandler1
				anotherKey: anotherValue
			)
    		post /api/foo1 (SingleExample)
		}
	`, ast.WithErrorCallback(func(err error) {
		assert.Nil(t, err)
	}))
	visitor := ast.NewApiVisitor()
	result := p.ServiceBlock().Accept(visitor)
	group, ok := result.(spec.Group)
	assert.True(t, ok)
	assert.Equal(t, len(group.Routes), 1)
}

func testServiceAnnotation(t *testing.T, content, key, value string) {
	p := ast.NewParser(content, ast.WithErrorCallback(func(err error) {
		assert.Nil(t, err)
	}))
	visitor := ast.NewApiVisitor()
	result := p.ServerMeta().Accept(visitor)
	anno, ok := result.(spec.Annotation)
	assert.True(t, ok)
	assert.Equal(t, anno.Properties[key], value)
}
