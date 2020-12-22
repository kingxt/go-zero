package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const structLit = "Foo {\n        VString string `json:\"vString\"`\n        VBool bool `json:\"vBool\"`\n        VInt8 int8 `json:\"vInt8\"`\n        VInt16 int16 `json:\"vInt16\"`\n        VInt32 int32 `json:\"vInt32\"`\n        VInt64 int64 `json:\"vInt64\"`\n        VInt int `json:\"vInt\"`\n        VUInt8 uint8 `json:\"vUInt8\"`\n        VUInt16 uint16 `json:\"vUInt16\"`\n        VUInt32 uint32 `json:\"vUInt32\"`\n        VUInt64 uint64 `json:\"vUInt64\"`\n        VFloat32 float32 `json:\"vFloat32\"`\n        VFloat64 float64 `json:\"vFloat64\"`\n        VByte byte `json:\"vByte\"`\n        VRune rune `json:\"vRune\"`\n        VMap map[string]int `json:\"vMap\"`\n        VArray []int `json:\"vArray\"`\n        VStruct Foo `json:\"vStruct\"`\n        VStructPointer *Foo `json:\"vStructPointer\"`\n        VInterface interface{} `json:\"vInterface\"`\n        T time.Time\n    }"

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
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}

	test(t, do, spec.InterfaceType{StringExpr: "interface{}"}, false, `interface{}`)
	test(t, do, nil, true, `interface`)
	test(t, do, nil, true, `*interface`)
	test(t, do, nil, true, `interface{ }`)
	test(t, do, nil, true, `interface{`)
}

func TestDataType(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.DataType().Accept(visitor)
	}

	test(t, do, spec.PointerType{
		StringExpr: "*int",
		Star: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `*int`)

	test(t, do, spec.BasicType{
		StringExpr: "int",
		Name:       "int",
	}, false, `int`)

	test(t, do, spec.Type{
		Name: "User",
	}, false, `User`)

	test(t, do, spec.MapType{
		StringExpr: "map[string]int",
		Key:        "string",
		Value: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `map[string]int`)

	test(t, do, spec.ArrayType{
		StringExpr: "[]int",
		ArrayType: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
	}, false, `[]int`)

	test(t, do, spec.InterfaceType{StringExpr: "interface{}"}, false, `interface{}`)
}

func TestField(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeField().Accept(visitor)
	}

	test(t, do, spec.Member{
		Name: "Name",
		Type: "string",
		Expr: spec.BasicType{
			StringExpr: "string",
			Name:       "string",
		},
		Tag: "",
	}, false, `Name string`)

	test(t, do, spec.Member{
		Name: "Name",
		Type: "int",
		Expr: spec.BasicType{
			StringExpr: "int",
			Name:       "int",
		},
		Tag: "",
	}, false, `Name int`)

	test(t, do, spec.Member{
		Name: "Name",
		Type: "map[string]int",
		Expr: spec.MapType{
			StringExpr: "map[string]int",
			Key:        "string",
			Value: spec.BasicType{
				StringExpr: "int",
				Name:       "int",
			},
		},
		Tag: "",
	}, false, `Name map[string]int`)

	test(t, do, spec.Member{
		Name:     "User",
		IsInline: true,
	}, false, `User`)

	test(t, do, nil, true, `User{}`)
	test(t, do, nil, true, `User struct{}`)
}

func TestStruct(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeStruct().Accept(visitor)
	}

	test(t, do, spec.Type{
		Name: "User",
	}, false, `User {}`)

	test(t, do, spec.Type{
		Name: "Foo",
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
	}, false, structLit)
}

func TestAlias(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeAlias().Accept(visitor)
	}

	test(t, do, nil, true, `Gender int`)
}

func TestTypeSpec(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeSpec().Accept(visitor)
	}

	test(t, do, spec.Type{
		Name: "User",
	}, false, `User {}`)

	test(t, do, nil, true, `Gender int`)
}

func TestTypeLit(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeLit().Accept(visitor)
	}

	test(t, do, spec.Type{
		Name: "User",
	}, false, `type User {}`)

	test(t, do, nil, true, `type Gender int`)
}

func TestTypeGroup(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeGroup().Accept(visitor)
	}

	test(t, do, []spec.Type{
		{
			Name: "User",
		},
	}, false, `type (
		User {}
	)`)

	test(t, do, nil, true, `type (
		Gender int
	)`)
}

func TestTypeBlock(t *testing.T) {
	do := func(p *ast.Parser, visitor *ast.ApiVisitor) interface{} {
		return p.TypeBlock().Accept(visitor)
	}

	test(t, do, []spec.Type{
		{
			Name: "User",
		},
	}, false, `type (
		User {}
	)`)

	test(t, do, nil, true, `type (
		Gender int
	)`)
}
