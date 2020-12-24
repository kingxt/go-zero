package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const duplicateInfoBlock = `
syntax = "v1"

info()
info()
`

const duplicateImport = `
syntax = "v1"

import "foo.api"
import "foo.api"
`

const duplicateTypeLit = `
syntax = "v1"

type Foo {}
type Foo {}
`

const duplicateTypeInGroup = `
syntax = "v1"

type (
	Bar {}
	Bar {}
)
`

const duplicateServiceBlock = `
syntax = "v1"

service foo-api{}
service bar-api{}
`

const oldApi = `
info()

type Foo {}
`

const duplicateInfoKeyBlock = `
syntax = "v1"
info (
	title: this is title
	desc: "new line"
	desc: 你好
)
`

func TestApi(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.Api().Accept(visitor)
	}

	test(t, do, nil, true, duplicateInfoBlock)
	test(t, do, nil, true, duplicateImport)
	test(t, do, nil, true, duplicateTypeLit)
	test(t, do, nil, true, duplicateTypeInGroup)
	test(t, do, nil, true, duplicateServiceBlock)
	test(t, do, nil, true, duplicateInfoKeyBlock)
	test(t, do, spec.ApiSpec{
		Info: spec.Info{
			Proterties: map[string]string{},
		},
		Types: []spec.Type{
			{
				Name: "Foo",
			},
		},
		Service: spec.Service{},
	}, false, oldApi)
}

func TestToken(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.Api().Accept(visitor)
	}
	test(t, do, nil, true, `infos ()`)
}
