package test

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"

var parser = ast.NewParser(ast.WithParserPrefix("test.api"), ast.ParserDebug)
