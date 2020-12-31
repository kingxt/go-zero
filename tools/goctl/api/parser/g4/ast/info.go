package ast

import (
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

type InfoExpr struct {
	Info        Expr
	Lp          Expr
	Rp          Expr
	DocExpr     Expr
	CommentExpr Expr
	Kvs         []*KvExpr
}

func (v *ApiVisitor) VisitInfoSpec(ctx *api.InfoSpecContext) interface{} {
	var expr InfoExpr
	expr.DocExpr = v.getDoc(ctx.GetDoc())
	expr.CommentExpr = v.getDoc(ctx.GetComment())
	expr.Info = v.newExprWithToken(ctx.GetInfoToken())
	expr.Lp = v.newExprWithToken(ctx.GetLp())
	expr.Rp = v.newExprWithToken(ctx.GetRp())
	list := ctx.AllKvLit()
	for _, each := range list {
		kvExpr := each.Accept(v).(*KvExpr)
		expr.Kvs = append(expr.Kvs, kvExpr)
	}
	if v.infoFlag {
		v.fmtErrorf(expr.Info, "duplicate declaration 'info'")
	}
	return &expr
}

func (i *InfoExpr) Format() error {
	// todo
	return nil
}

func (i *InfoExpr) Equal(v interface{}) bool {
	if i == nil {
		if v != nil {
			return false
		}

		return true
	}

	if v == nil {
		return false
	}

	info, ok := v.(*InfoExpr)
	if !ok {
		return false
	}

	if !i.Info.Equal(info.Info) {
		return false
	}

	if !EqualDoc(i, info) {
		return false
	}

	var expected, actual []*KvExpr
	expected = append(expected, info.Kvs...)
	actual = append(actual, i.Kvs...)

	if len(expected) != len(actual) {
		return false
	}

	for index, each := range expected {
		ac := actual[index]
		if !each.Equal(ac) {
			return false
		}
	}

	return true
}

func (i *InfoExpr) Doc() Expr {
	return i.DocExpr
}

func (i *InfoExpr) Comment() Expr {
	return i.CommentExpr
}
