package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestSyntaxLit(t *testing.T) {
	do := func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.SyntaxLit().Accept(visitor)
	}

	test(t, do, spec.ApiSyntax{Version: "v1"}, false, `syntax = "v1"`)
	test(t, do, nil, true, `syntax = "v 1"`)
	test(t, do, nil, true, `syntax = "1"`)
	test(t, do, nil, true, `syntax = ""`)
	test(t, do, nil, true, `syntax1 = "v1"`)
	test(t, do, nil, true, `syntax`)
	test(t, do, nil, true, `syntax=`)
	test(t, do, nil, true, `syntax "v1"`)
	test(t, do, nil, true, `syntax = "v0"`)
}
