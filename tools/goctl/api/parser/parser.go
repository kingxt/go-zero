package parser

import (
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func Parser(filename string) (*spec.ApiSpec, error) {
	parser := ast.NewParser(ast.WithParserPrefix(filepath.Base(filename)))
	ast, err := parser.Parse(filename)
	if err != nil {
		return nil, err
	}

	return convert2Spec(ast)
}

func ParserContent(content string) (*spec.ApiSpec, error) {
	parser := ast.NewParser()
	ast, err := parser.ParseContent(content)
	if err != nil {
		return nil, err
	}
	return convert2Spec(ast)
}

// todo
func convert2Spec(in *ast.Api) (*spec.ApiSpec, error) {
	return nil, nil
}
