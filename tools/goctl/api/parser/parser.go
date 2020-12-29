package parser

import (
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func Parser(filename string) (*spec.ApiSpec, error) {
	parser := ast.NewParser(ast.WithPrefix(filepath.Base(filename)))
	apiSpec, err := parser.Parse(filename)
	return apiSpec, err
}

func ParserContent(content string) (*spec.ApiSpec, error) {
	parser := ast.NewParser()
	apiSpec, err := parser.ParseContent(content)
	return apiSpec, err
}
