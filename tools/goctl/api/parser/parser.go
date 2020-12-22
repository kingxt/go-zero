package parser

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
	"github.com/tal-tech/go-zero/tools/goctl/util"
)

func Parser(filename string) (*spec.ApiSpec, error) {
	apiAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	api, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parser := ast.NewParser(string(api))
	visitor := ast.NewApiVisitor()
	apiSpec := parser.Api().Accept(visitor).(spec.ApiSpec)

	for _, item := range apiSpec.Import.List {
		importFile := strings.TrimSpace(item)
		if len(importFile) > 0 {
			item = strings.TrimSpace(item)
			var path = item
			if !util.FileExists(item) {
				path = filepath.Join(filepath.Dir(apiAbsPath), item)
			}
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, errors.New("import api file not exist: " + item)
			}

			parser := ast.NewParser(string(content))
			visitor := ast.NewApiVisitor()
			importApiSpec := parser.Api().Accept(visitor).(spec.ApiSpec)

			if len(importApiSpec.Import.List) > 0 {
				return nil, errors.New("import api should not import another api file recursive")
			}

			apiSpec.Types = append(apiSpec.Types, importApiSpec.Types...)
			apiSpec.Service.Groups = append(apiSpec.Service.Groups, importApiSpec.Service.Groups...)
		}
	}

	return &apiSpec, nil
}
