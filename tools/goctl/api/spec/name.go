package internal

import (
	"strings"

	spec "github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func (t spec.PrimitiveType) Name() string {
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
