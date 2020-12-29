package ast

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"

type SyntaxExpr struct {
	Syntax      Expr
	Assign      Expr
	Version     Expr
	DocExpr     Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitSyntaxLit(ctx *api.SyntaxLitContext) interface{} {
	syntax := v.newExprWithToken(ctx.GetSyntaxToken())
	assign := v.newExprWithToken(ctx.GetAssign())
	version := v.newExprWithToken(ctx.GetVersion())
	docExpr := v.getDoc(ctx.GetDoc())
	commentExpr := v.getDoc(ctx.GetComment())
	return &SyntaxExpr{
		Syntax:      syntax,
		Assign:      assign,
		Version:     version,
		DocExpr:     docExpr,
		CommentExpr: commentExpr,
	}
}

func (s *SyntaxExpr) Format() error {
	// todo
	return nil
}

func (s *SyntaxExpr) Equal(v interface{}) bool {
	syntax, ok := v.(*SyntaxExpr)
	if !ok {
		return false
	}

	if !EqualDoc(s, syntax) {
		return false
	}

	return ExprEqual(s.Syntax, syntax.Syntax) &&
		ExprEqual(s.Assign, syntax.Assign) &&
		ExprEqual(s.Version, syntax.Version)
}

func (s *SyntaxExpr) Doc() Expr {
	return s.DocExpr
}

func (s *SyntaxExpr) Comment() Expr {
	return s.CommentExpr
}
