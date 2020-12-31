package ast

import (
	"sort"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

type (
	// TypeAlias, TypeStruct
	TypeExpr interface {
		Doc() Expr
		Comment() Expr
	}

	TypeAlias struct {
		Name        Expr
		Assign      Expr
		DataType    Expr
		DocExpr     Expr
		CommentExpr Expr
	}

	TypeStruct struct {
		Name        Expr
		Struct      Expr
		LBrace      Expr
		RBrace      Expr
		DocExpr     Expr
		CommentExpr Expr
		Fields      []*TypeField
	}

	TypeField struct {
		IsAnonymous bool
		// Name is nil if IsAnonymous
		Name        Expr
		DataType    DataType
		Tag         Expr
		DocExpr     Expr
		CommentExpr Expr
	}

	// Literal, Interface, Map, Array, Time, Pointer
	DataType interface {
		Expr() Expr
		Equal(dt DataType) bool
	}

	// int, bool, Foo,...
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
	iAnonymousFiled := ctx.AnonymousFiled()
	iNormalFieldContext := ctx.NormalField()
	if iAnonymousFiled != nil {
		return iAnonymousFiled.Accept(v).(*TypeField)
	}
	if iNormalFieldContext != nil {
		return iNormalFieldContext.Accept(v).(*TypeField)
	}
	return nil
}

func (v *ApiVisitor) VisitNormalField(ctx *api.NormalFieldContext) interface{} {
	var field TypeField
	field.Name = v.newExprWithToken(ctx.GetFieldName())
	iDataTypeContext := ctx.DataType()
	if iDataTypeContext != nil {
		field.DataType = iDataTypeContext.Accept(v).(DataType)
	}
	field.Tag = v.newExprWithToken(ctx.GetTag())
	field.DocExpr = v.getDoc(ctx.GetDoc())
	field.CommentExpr = v.getDoc(ctx.GetComment())
	return &field
}

func (v *ApiVisitor) VisitAnonymousFiled(ctx *api.AnonymousFiledContext) interface{} {
	start := ctx.GetStart()
	stop := ctx.GetStop()
	var field TypeField
	field.IsAnonymous = true
	if ctx.GetStar() != nil {
		field.DataType = &Pointer{
			PointerExpr: v.newExprWithText(ctx.GetStar().GetText()+ctx.ID().GetText(), start.GetLine(), start.GetColumn(), start.GetStart(), stop.GetStop()),
			Star:        v.newExprWithToken(ctx.GetStar()),
			Name:        v.newExprWithTerminalNode(ctx.ID()),
		}
	} else {
		field.DataType = &Literal{Literal: v.newExprWithTerminalNode(ctx.ID())}
	}
	field.DocExpr = v.getDoc(ctx.GetDoc())
	field.CommentExpr = v.getDoc(ctx.GetComment())
	return &field
}

func (v *ApiVisitor) VisitDataType(ctx *api.DataTypeContext) interface{} {
	if ctx.ID() != nil {
		return &Literal{Literal: v.newExprWithTerminalNode(ctx.ID())}
	}
	if ctx.MapType() != nil {
		return ctx.MapType().Accept(v)
	}
	if ctx.ArrayType() != nil {
		return ctx.ArrayType().Accept(v)
	}
	if ctx.GetInter() != nil {
		return Interface{Literal: v.newExprWithToken(ctx.GetInter())}
	}
	if ctx.GetTime() != nil {
		return Time{Literal: v.newExprWithToken(ctx.GetTime())}
	}
	if ctx.PointerType() != nil {
		return ctx.PointerType().Accept(v)
	}
	return ctx.TypeStruct().Accept(v)
}

func (v *ApiVisitor) VisitPointerType(ctx *api.PointerTypeContext) interface{} {
	return &Pointer{
		PointerExpr: v.newExprWithText(ctx.GetText(), ctx.GetStar().GetLine(), ctx.GetStar().GetColumn(), ctx.GetStar().GetStart(), ctx.ID().GetSymbol().GetStop()),
		Star:        v.newExprWithToken(ctx.GetStar()),
		Name:        v.newExprWithTerminalNode(ctx.ID()),
	}
}

func (v *ApiVisitor) VisitMapType(ctx *api.MapTypeContext) interface{} {
	return Map{
		MapExpr: v.newExprWithText(ctx.GetText(), ctx.GetMapToken().GetLine(), ctx.GetMapToken().GetColumn(),
			ctx.GetMapToken().GetStart(), ctx.GetValue().GetStop().GetStop()),
		Map:    v.newExprWithToken(ctx.GetMapToken()),
		LBrack: v.newExprWithToken(ctx.GetLbrack()),
		RBrack: v.newExprWithToken(ctx.GetRbrack()),
		Key:    v.newExprWithToken(ctx.GetKey()),
		Value:  ctx.GetValue().Accept(v).(DataType),
	}
}

func (v *ApiVisitor) VisitArrayType(ctx *api.ArrayTypeContext) interface{} {
	return Array{
		ArrayExpr: v.newExprWithText(ctx.GetText(), ctx.GetLbrack().GetLine(), ctx.GetLbrack().GetColumn(), ctx.GetLbrack().GetStart(), ctx.DataType().GetStop().GetStop()),
		LBrack:    v.newExprWithToken(ctx.GetLbrack()),
		RBrack:    v.newExprWithToken(ctx.GetRbrack()),
		Literal:   ctx.DataType().Accept(v).(DataType),
	}
}

func (a *TypeAlias) Doc() Expr {
	return a.DocExpr
}

func (a *TypeAlias) Comment() Expr {
	return a.CommentExpr
}

func (l *Literal) Expr() Expr {
	return l.Literal
}

func (l *Literal) Equal(dt DataType) bool {
	if l == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Literal)
	if !ok {
		return false
	}

	return l.Literal.Equal(v.Literal)
}

func (i *Interface) Expr() Expr {
	return i.Literal
}

func (i *Interface) Equal(dt DataType) bool {
	if i == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Interface)
	if !ok {
		return false
	}

	return i.Literal.Equal(v.Literal)
}

func (m *Map) Expr() Expr {
	return m.MapExpr
}

func (m *Map) Equal(dt DataType) bool {
	if m == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Map)
	if !ok {
		return false
	}

	if !m.Key.Equal(v.Key) {
		return false
	}

	if !m.Value.Equal(v.Value) {
		return false
	}

	if !m.MapExpr.Equal(v.MapExpr) {
		return false
	}

	return m.Map.Equal(v.Map)
}

func (a *Array) Expr() Expr {
	return a.ArrayExpr
}

func (a *Array) Equal(dt DataType) bool {
	if a == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Array)
	if !ok {
		return false
	}

	if !a.ArrayExpr.Equal(v.ArrayExpr) {
		return false
	}

	return a.Literal.Equal(v.Literal)
}

func (t *Time) Expr() Expr {
	return t.Literal
}

func (t *Time) Equal(dt DataType) bool {
	if t == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Time)
	if !ok {
		return false
	}

	return t.Literal.Equal(v.Literal)
}

func (p *Pointer) Expr() Expr {
	return p.PointerExpr
}

func (p *Pointer) Equal(dt DataType) bool {
	if p == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*Pointer)
	if !ok {
		return false
	}

	if !p.PointerExpr.Equal(v.PointerExpr) {
		return false
	}

	if !p.Star.Equal(v.Star) {
		return false
	}

	return p.Name.Equal(v.Name)
}

func (s *TypeStruct) Expr() Expr {
	return s.Name
}

func (s *TypeStruct) Equal(dt interface{}) bool {
	if s == nil {
		if dt != nil {
			return false
		}
		return true
	}

	if dt == nil {
		return false
	}

	v, ok := dt.(*TypeStruct)
	if !ok {
		return false
	}

	if !s.Name.Equal(v.Name) {
		return false
	}

	if EqualDoc(s, v) {
		return false
	}

	if s.Struct.Equal(v.Struct) {
		return false
	}

	if len(s.Fields) != len(v.Fields) {
		return false
	}

	var expected, acual []*TypeField
	expected = append(expected, s.Fields...)
	acual = append(acual, v.Fields...)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].Name.Text() < expected[j].Name.Text()
	})
	sort.Slice(acual, func(i, j int) bool {
		return acual[i].Name.Text() < acual[j].Name.Text()
	})

	for index, each := range expected {
		ac := acual[index]
		if !each.Equal(ac) {
			return false
		}
	}

	return true
}

func (s *TypeStruct) Doc() Expr {
	return s.DocExpr
}

func (s *TypeStruct) Comment() Expr {
	return s.CommentExpr
}

func (s *TypeStruct) Format() error {
	// todo
	return nil
}

func (t *TypeField) Equal(v interface{}) bool {
	if t == nil {
		if v != nil {
			return false
		}

		return true
	}

	if v == nil {
		return false
	}

	f, ok := v.(*TypeField)
	if !ok {
		return false
	}

	if t.IsAnonymous != f.IsAnonymous {
		return false
	}

	if !t.DataType.Equal(f.DataType) {
		return false
	}

	if !t.IsAnonymous {
		if !t.Name.Equal(f.Name) {
			return false
		}
		if !t.Tag.Equal(f.Tag) {
			return false
		}
	}

	return EqualDoc(t, f)
}

func (t *TypeField) Doc() Expr {
	return t.DocExpr
}

func (t *TypeField) Comment() Expr {
	return t.CommentExpr
}

func (t *TypeField) Format() error {
	// todo
	return nil
}
