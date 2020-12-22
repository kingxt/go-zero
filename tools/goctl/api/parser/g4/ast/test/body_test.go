package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
)

func TestBody(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.Body().Accept(visitor)
	}
	test(t, do, nil, false, `info()`)
}
