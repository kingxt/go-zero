package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

var parser = NewParser(WithParserPrefix("test.api"), WithParserDebug())

var normal = `
	syntax="v1"
	
	info (
		foo: bar
	)

	type Foo {
		Bar int
	}

	@server(
		foo: bar
	)
	service foo-api{
		@doc("foo")
		@handler foo
		post /foo (Foo) returns (Foo)
	}
`

func TestApiParser(t *testing.T) {
	_, err := parser.Accept(func(p *api.ApiParserParser, visitor *ApiVisitor) interface{} {
		return p.Api().Accept(visitor)
	}, normal)
	assert.Nil(t, err)
}
