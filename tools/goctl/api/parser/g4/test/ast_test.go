package test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

var parser = ast.NewParser(ast.WithParserPrefix("test.api"), ast.ParserDebug)

func TestApi(t *testing.T) {
	fn := func(p *api.ApiParserParser, visitor *ast.ApiVisitor) interface{} {
		return p.Api().Accept(visitor)
	}
	content, err := ioutil.ReadFile("./test.api")
	assert.Nil(t, err)

	v, err := parser.Accept(fn, string(content))
	assert.Nil(t, err)
	api := v.(*ast.Api)
	api.Equal(&ast.Api{
		Syntax:  nil,
		Import:  nil,
		Info:    nil,
		Type:    nil,
		Service: nil,
	})
}
