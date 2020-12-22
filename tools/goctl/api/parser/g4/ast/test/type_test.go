package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestPointer(t *testing.T) {
	do := func(p *ast.Parser, v *ast.ApiVisitor) interface{} {
		return p.Pointer().Accept(v)
	}
	test(t, do, spec.BasicType{
		StringExpr: "int",
		Name:       "int",
	}, false, `int`)

	test(t, do, spec.BasicType{
		StringExpr: "bool",
		Name:       "bool",
	}, false, `bool`)

	test(t, do, spec.Type{
		Name: "User",
	}, false, `User`)

	test(t, do, spec.PointerType{
		StringExpr: "*int",
		Star: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `*int`)

	test(t, do, spec.PointerType{
		StringExpr: "*User",
		Star: spec.Type{
			Name: "User",
		},
	}, false, `*User`)

	test(t, do, spec.PointerType{
		StringExpr: "**User",
		Star: spec.PointerType{
			StringExpr: "*User",
			Star: spec.Type{
				Name: "User",
			},
		},
	}, false, `**User`)

	test(t, do, nil, true, `map[string]string`)
}

func TestArrayType(t *testing.T) {
	do := func(p *ast.Parser, v *ast.ApiVisitor) interface{} {
		return p.ArrayType().Accept(v)
	}

	test(t, do, spec.ArrayType{
		StringExpr: "[]int",
		ArrayType: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `[]int`)

	test(t, do, spec.ArrayType{
		StringExpr: "[]*int",
		ArrayType: spec.PointerType{
			StringExpr: "*int",
			Star: spec.BasicType{
				StringExpr: "int",
				Name:       "int",
			},
		},
	}, false, `[]*int`)

	test(t, do, spec.ArrayType{
		StringExpr: "[][]int",
		ArrayType: spec.ArrayType{
			StringExpr: "[]int",
			ArrayType: spec.BasicType{
				StringExpr: "int",
				Name:       "int",
			},
		},
	}, false, `[][]int`)

	test(t, do, spec.ArrayType{
		StringExpr: "[]map[string]string",
		ArrayType: spec.MapType{
			StringExpr: "map[string]string",
			Key:        "string",
			Value: spec.BasicType{
				StringExpr: "string",
				Name:       "string",
			},
		},
	}, false, `[]map[string]string`)

	test(t, do, spec.ArrayType{
		StringExpr: "[]interface{}",
		ArrayType:  spec.InterfaceType{StringExpr: "interface{}"},
	}, false, `[]interface{}`)

	test(t, do, spec.ArrayType{
		StringExpr: "[][]*User",
		ArrayType: spec.ArrayType{
			StringExpr: "[]*User",
			ArrayType: spec.PointerType{
				StringExpr: "*User",
				Star: spec.Type{
					Name: "User",
				},
			},
		},
	}, false, `[][]*User`)

	test(t, do, nil, true, `[[]][]int`)
	test(t, do, nil, true, `[]{}int`)
	test(t, do, nil, true, `[]`)
	test(t, do, nil, true, `[int]`)
	test(t, do, nil, true, `[int]int`)
	test(t, do, nil, true, `[]map[User]int`)
}

func TestMapType(t *testing.T) {
	do := func(p *ast.Parser, v *ast.ApiVisitor) interface{} {
		return p.MapType().Accept(v)
	}

	test(t, do, spec.MapType{
		StringExpr: "map[string]int",
		Key:        "string",
		Value: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `map[string]int`)

	test(t, do, spec.MapType{
		StringExpr: "map[string]bool",
		Key:        "string",
		Value: spec.BasicType{
			StringExpr: "bool",
			Name:       "bool",
		},
	}, false, `map[string]bool`)

	test(t, do, spec.MapType{
		StringExpr: "map[string]User",
		Key:        "string",
		Value: spec.Type{
			Name: "User",
		},
	}, false, `map[string]User`)

	test(t, do, spec.MapType{
		StringExpr: "map[string][]int",
		Key:        "string",
		Value: spec.ArrayType{
			StringExpr: "[]int",
			ArrayType: spec.BasicType{
				StringExpr: "int",
				Name:       "int",
			},
		},
	}, false, `map[string][]int`)

	test(t, do, spec.MapType{
		StringExpr: "map[string]map[int]User",
		Key:        "string",
		Value: spec.MapType{
			StringExpr: "map[int]User",
			Key:        "int",
			Value: spec.Type{
				Name: "User",
			},
		},
	}, false, `map[string]map[int]User`)

	test(t, do, spec.MapType{
		StringExpr: "map[string]**User",
		Key:        "string",
		Value: spec.PointerType{
			StringExpr: "**User",
			Star: spec.PointerType{
				StringExpr: "*User",
				Star: spec.Type{
					Name: "User",
				},
			},
		},
	}, false, `map[string]**User`)
}

func TestInterface(t *testing.T) {
	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.InterfaceType{StringExpr: "interface{}"}, false, `interface{}`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, nil, true, `interface`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, nil, true, `*interface`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, nil, true, `interface{ }`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, nil, true, `interface{`)
}

func TestDataType(t *testing.T) {
	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.PointerType{
		StringExpr: "*int",
		Star: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `*int`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.BasicType{
		StringExpr: "int",
		Name:       "int",
	}, false, `int`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.Type{
		Name: "User",
	}, false, `User`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.MapType{
		StringExpr: "map[string]int",
		Key:        "string",
		Value: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `map[string]int`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.ArrayType{
		StringExpr: "[]int",
		ArrayType: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `[]int`)

	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}, spec.InterfaceType{StringExpr: "interface{}"}, false, `interface{}`)
}

func TestField(t *testing.T) {
	test(t, func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeField().Accept(visitor)
	}, spec.Member{
		Name: "Name",
		Type: "string",
		Expr: spec.BasicType{
			StringExpr: "string",
			Name:       "string",
		},
		Tag: "",
	}, false, `Name string`)
}

func test(t *testing.T, do func(p *ast.Parser, visitor *ast.ApiVisitor) interface{}, expected interface{}, expectedErr bool, content string) {
	defer func() {
		p := recover()
		if expectedErr {
			if logEnable {
				fmt.Printf("%+v\n", p)
			}
			assert.NotNil(t, p)
			return
		}
		assert.Nil(t, p)
	}()

	p := ast.NewParser(content)
	visitor := ast.NewApiVisitor()
	result := do(p, visitor)
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(result)
	assert.JSONEq(t, string(expectedJson), string(actualJson))
}
