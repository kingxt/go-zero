package spec

import (
	"fmt"
	"strings"
)

type (
	Doc []string

	Annotation struct {
		Properties map[string]string
	}

	ApiSyntax struct {
		Version string
	}

	ApiSpec struct {
		Info    Info
		Syntax  ApiSyntax
		Imports []Import
		Types   []Type
		Service Service
	}

	Import struct {
		Value string
	}

	Group struct {
		Annotation Annotation
		Routes     []Route
	}

	Info struct {
		// Deprecated: use Properties instead
		Title string
		// Deprecated: use Properties instead
		Desc string
		// Deprecated: use Properties instead
		Version string
		// Deprecated: use Properties instead
		Author string
		// Deprecated: use Properties instead
		Email      string
		Properties map[string]string
	}

	Member struct {
		Name string
		// 数据类型字面值，如：string、map[int]string、[]int64、[]*User
		Type    Type
		Tag     string
		Comment string
		// 成员头顶注释说明
		Docs     Doc
		IsInline bool
	}

	Route struct {
		Annotation   Annotation
		Method       string
		Path         string
		RequestType  Type
		ResponseType Type
		Docs         Doc
		Handler      string
	}

	Service struct {
		Name   string
		Groups []Group
	}

	Type interface {
		Name() string
		GolangExpr(pkg ...string) string
	}

	DefineStruct struct {
		RawName string
		Members []Member
		Docs    Doc
	}

	// 系统预设基本数据类型 bool int32 int64 float32
	PrimitiveType struct {
		RawName string
	}

	MapType struct {
		RawName string
		// only support the PrimitiveType
		Key string
		// it can be asserted as PrimitiveType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${PrimitiveType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		Value Type
	}

	ArrayType struct {
		RawName string
		Value   Type
	}

	InterfaceType struct {
		RawName string
	}

	PointerType struct {
		RawName string
		Type    Type
	}
)

func (t PrimitiveType) Name() string {
	return t.RawName
}

func (t PrimitiveType) GolangExpr(_ ...string) string {
	return t.RawName
}

func (t DefineStruct) Name() string {
	return t.RawName
}

func (t DefineStruct) GolangExpr(pkg ...string) string {
	if len(pkg) > 1 {
		panic("package cannot be more than 1")
	}

	if len(pkg) == 0 {
		return t.RawName
	}

	return fmt.Sprintf("%s.%s", pkg[0], strings.Title(t.RawName))
}

func (t MapType) Name() string {
	return t.RawName
}

func (t MapType) GolangExpr(pkg ...string) string {
	if len(pkg) > 1 {
		panic("package cannot be more than 1")
	}

	if len(pkg) == 0 {
		return t.RawName
	}

	return fmt.Sprintf("map[%s]%s", t.Key, t.Value.GolangExpr(pkg...))
}

func (t ArrayType) Name() string {
	return t.RawName
}

func (t ArrayType) GolangExpr(pkg ...string) string {
	if len(pkg) > 1 {
		panic("package cannot be more than 1")
	}

	if len(pkg) == 0 {
		return t.RawName
	}

	return fmt.Sprintf("[]%s", t.Value.GolangExpr(pkg...))
}

func (t PointerType) Name() string {
	return t.RawName
}

func (t PointerType) GolangExpr(pkg ...string) string {
	if len(pkg) > 1 {
		panic("package cannot be more than 1")
	}

	if len(pkg) == 0 {
		return t.RawName
	}

	return fmt.Sprintf("*%s", t.Type.GolangExpr(pkg...))
}

func (t InterfaceType) Name() string {
	return t.RawName
}

func (t InterfaceType) GolangExpr(_ ...string) string {
	return t.RawName
}
