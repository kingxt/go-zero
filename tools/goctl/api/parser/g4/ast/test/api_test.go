package test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestApiFile(t *testing.T) {
	abs, err := filepath.Abs("./foo.api")
	assert.Nil(t, err)

	data, err := ioutil.ReadFile(abs)
	assert.Nil(t, err)

	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.Api().Accept(visitor)
	}

	//annotation := spec.Annotation{Properties: map[string]string{}}

	test(t, do, spec.ApiSpec{
		Syntax: spec.ApiSyntax{
			Version: "v1",
		},
		Import: spec.ApiImport{
			List: []spec.Import{
				{
					Value: "foo.api",
				},
				{
					Value: "bar.api",
				},
				{
					Value: "foo/bar.api",
				},
			},
		},
		Info: spec.Info{
			Proterties: map[string]string{
				"title":   "foo title",
				"desc":    "foo\n    desc",
				"author":  "foo author",
				"email":   "foo email",
				"version": "foo version",
			},
		},
		Types: []spec.Type{
			{
				Name: "Foo",
				Members: []spec.Member{
					{
						Name: "Bar",
						Type: "string",
						Expr: spec.BasicType{
							StringExpr: "string",
							Name:       "string",
						},
						Tag:      "",
						Comment:  "",
						Comments: nil,
						Docs:     nil,
						IsInline: false,
					},
				},
			},
			{
				Name: "GoFoo",
				Members: []spec.Member{
					{
						Name: "Bar",
						Type: "string",
						Expr: spec.BasicType{
							StringExpr: "string",
							Name:       "string",
						},
						Tag: `json:"bar"`,
					},
				},
			},
			{
				Name: "GroupFoo",
				Members: []spec.Member{
					{
						Name: "Bar",
						Type: "string",
						Expr: spec.BasicType{
							StringExpr: "string",
							Name:       "string",
						},
						Tag: `json:"bar"`,
					},
				},
			},
			{
				Name: "GroupBar",
				Members: []spec.Member{
					{
						Name: "VString",
						Type: "string",
						Expr: spec.BasicType{
							StringExpr: "string",
							Name:       "string",
						},
						Tag: `json:"vString"`,
					},
					{
						Name: "VBool",
						Type: "bool",
						Expr: spec.BasicType{
							StringExpr: "bool",
							Name:       "bool",
						},
						Tag: `json:"vBool"`,
					},
					{
						Name: "VInt8",
						Type: "int8",
						Expr: spec.BasicType{
							StringExpr: "int8",
							Name:       "int8",
						},
						Tag: `json:"vInt8"`,
					},
					{
						Name: "VInt16",
						Type: "int16",
						Expr: spec.BasicType{
							StringExpr: "int16",
							Name:       "int16",
						},
						Tag: `json:"vInt16"`,
					},
					{
						Name: "VInt32",
						Type: "int32",
						Expr: spec.BasicType{
							StringExpr: "int32",
							Name:       "int32",
						},
						Tag: `json:"vInt32"`,
					},
					{
						Name: "VInt64",
						Type: "int64",
						Expr: spec.BasicType{
							StringExpr: "int64",
							Name:       "int64",
						},
						Tag: `json:"vInt64"`,
					},
					{
						Name: "VInt",
						Type: "int",
						Expr: spec.BasicType{
							StringExpr: "int",
							Name:       "int",
						},
						Tag: `json:"vInt"`,
					},
					{
						Name: "VUInt8",
						Type: "uint8",
						Expr: spec.BasicType{
							StringExpr: "uint8",
							Name:       "uint8",
						},
						Tag: `json:"vUInt8"`,
					},
					{
						Name: "VUInt16",
						Type: "uint16",
						Expr: spec.BasicType{
							StringExpr: "uint16",
							Name:       "uint16",
						},
						Tag: `json:"vUInt16"`,
					},
					{
						Name: "VUInt32",
						Type: "uint32",
						Expr: spec.BasicType{
							StringExpr: "uint32",
							Name:       "uint32",
						},
						Tag: `json:"vUInt32"`,
					},
					{
						Name: "VUInt64",
						Type: "uint64",
						Expr: spec.BasicType{
							StringExpr: "uint64",
							Name:       "uint64",
						},
						Tag: `json:"vUInt64"`,
					},
					{
						Name: "VFloat32",
						Type: "float32",
						Expr: spec.BasicType{
							StringExpr: "float32",
							Name:       "float32",
						},
						Tag: `json:"vFloat32"`,
					},
					{
						Name: "VFloat64",
						Type: "float64",
						Expr: spec.BasicType{
							StringExpr: "float64",
							Name:       "float64",
						},
						Tag: `json:"vFloat64"`,
					},
					{
						Name: "VByte",
						Type: "byte",
						Expr: spec.BasicType{
							StringExpr: "byte",
							Name:       "byte",
						},
						Tag: `json:"vByte"`,
					},
					{
						Name: "VRune",
						Type: "rune",
						Expr: spec.BasicType{
							StringExpr: "rune",
							Name:       "rune",
						},
						Tag: `json:"vRune"`,
					},
					{
						Name: "VMap",
						Type: "map[string]int",
						Expr: spec.MapType{
							StringExpr: "map[string]int",
							Key:        "string",
							Value: spec.BasicType{
								StringExpr: "int",
								Name:       "int",
							},
						},
						Tag: `json:"vMap"`,
					},
					{
						Name: "VArray",
						Type: "[]int",
						Expr: spec.ArrayType{
							StringExpr: "[]int",
							ArrayType: spec.BasicType{
								StringExpr: "int",
								Name:       "int",
							},
						},
						Tag: `json:"vArray"`,
					},
					{
						Name: "VStruct",
						Type: "Foo",
						Expr: spec.Type{
							Name: "Foo",
						},
						Tag: `json:"vStruct"`,
					},
					{
						Name: "VStructPointer",
						Type: "*Foo",
						Expr: spec.PointerType{
							StringExpr: "*Foo",
							Star: spec.Type{
								Name: "Foo",
							},
						},
						Tag: `json:"vStructPointer"`,
					},
					{
						Name: "VInterface",
						Type: "interface{}",
						Expr: spec.InterfaceType{StringExpr: "interface{}"},
						Tag:  `json:"vInterface"`,
					},
					{
						Name: "T",
						Type: "time.Time",
						Expr: spec.TimeType{StringExpr: "time.Time"},
					},
				},
			},
		},
		Service: spec.Service{
			Name: "foo-api",
			Groups: []spec.Group{
				{
					Annotation: spec.Annotation{
						Properties: map[string]string{
							"jwt":        "Foo",
							"group":      "foo/bar",
							"anotherKey": "anotherValue",
						},
					},
					Routes: []spec.Route{
						{
							Annotation: spec.Annotation{
								Properties: map[string]string{
									"handler":    "fooHandler1",
									"anotherKey": "anotherValue",
								},
							},
							Method: "post",
							Path:   "/api/foo/bar",
							RequestType: spec.Type{
								Name: "Foo",
							},
							Docs:    []string{"foo1"},
							Handler: "fooHandler1",
						},
						{
							Method: "get",
							Path:   "/api/foo",
							RequestType: spec.Type{
								Name: "Foo",
							},
							ResponseType: spec.Type{
								Name: "Foo",
							},
							Docs:    []string{"foo2"},
							Handler: "fooHandler2",
						},
						{
							Method: "post",
							Path:   "/api/foo/:id",
							ResponseType: spec.Type{
								Name: "Foo",
							},
							Handler: "fooHandler3",
						},
						{
							Method:  "get",
							Path:    "/api/foo-bar",
							Handler: "fooHandler4",
						},
					},
				},
			},
		},
	}, false, string(data))
}
