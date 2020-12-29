package ast

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"

type (
	// TypeAlias, TypeStruct
	TypeExpr interface {
		Start() Expr
		Stop() Expr
		Doc() Expr
		Comment() Expr
	}

	TypeAlias struct {
		Name        Expr
		Assign      Expr
		DataType    Expr
		StartExpr   Expr
		StopExpr    Expr
		DocExpr     Expr
		CommentExpr Expr
	}

	TypeStruct struct {
		Name        Expr
		Struct      Expr
		LBrace      Expr
		RBrace      Expr
		StartExpr   Expr
		StopExpr    Expr
		DocExpr     Expr
		CommentExpr Expr
		Fields      []*TypeField
	}

	TypeField struct {
		isAnonymous bool
		Name        Expr
		DataType    DataType
		Tag         Expr
		DocExpr     Expr
		CommentExpr Expr
	}

	// Literal, Interface, Map, Array, Time, Pointer, TypeStruct
	DataType interface {
		Expr() Expr
	}

	// int, bool,Foo
	Literal struct {
		Literal Expr
	}

	Interface struct {
		Literal Expr
	}

	Map struct {
		MapExpr Expr
		Map     Expr
		LBrack  Expr
		RBrack  Expr
		Key     Expr
		Value   DataType
	}

	Array struct {
		ArrayExpr Expr
		LBrack    Expr
		RBrack    Expr
		Literal   DataType
	}

	Time struct {
		Literal Expr
	}

	Pointer struct {
		PointerExpr Expr
		Star        Expr
		Name        Expr
	}
)

func (v *ApiVisitor) VisitTypeSpec(ctx *api.TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitTypeLit(ctx *api.TypeLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitTypeBlock(ctx *api.TypeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitTypeBody(ctx *api.TypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitTypeStruct(ctx *api.TypeStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitTypeAlias(ctx *api.TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitField(ctx *api.FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitDataType(ctx *api.DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitPointerType(ctx *api.PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitMapType(ctx *api.MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitArrayType(ctx *api.ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (a *TypeAlias) Start() Expr {
	return a.StartExpr
}

func (a *TypeAlias) Stop() Expr {
	return a.StopExpr
}

func (a *TypeAlias) Doc() Expr {
	return a.DocExpr
}

func (a *TypeAlias) Comment() Expr {
	return a.CommentExpr
}

func (s *TypeStruct) Start() Expr {
	return s.StartExpr
}

func (s *TypeStruct) Stop() Expr {
	return s.StopExpr
}

func (s *TypeStruct) Doc() Expr {
	return s.DocExpr
}

func (s *TypeStruct) Comment() Expr {
	return s.CommentExpr
}

func (l *Literal) Expr() Expr {
	return l.Literal
}

func (i *Interface) Expr() Expr {
	return i.Literal
}

func (m *Map) Expr() Expr {
	return m.MapExpr
}

func (a *Array) Expr() Expr {
	return a.ArrayExpr
}

func (t *Time) Expr() Expr {
	return t.Literal
}

func (p *Pointer) Expr() Expr {
	return p.PointerExpr
}

func (s *TypeStruct) Expr() Expr {
	return s.Name
}
