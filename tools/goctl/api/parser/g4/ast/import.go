package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

type ImportExpr struct {
	Import      Expr
	Value       Expr
	DocExpr     Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitImportSpec(ctx *api.ImportSpecContext) interface{} {
	var list []*ImportExpr
	if ctx.ImportLit() != nil {
		lits := ctx.ImportLit().Accept(v).([]*ImportExpr)
		list = append(list, lits...)
	}
	if ctx.ImportBlock() != nil {
		blocks := ctx.ImportBlock().Accept(v).([]*ImportExpr)
		list = append(list, blocks...)
	}

	return list
}

func (v *ApiVisitor) VisitImportLit(ctx *api.ImportLitContext) interface{} {
	importToken := v.newExprWithToken(ctx.GetImportToken())
	valueExpr := ctx.ImportValue().Accept(v).(Expr)
	return []*ImportExpr{
		{
			Import:      importToken,
			Value:       valueExpr,
			DocExpr:     v.getDoc(ctx.GetDoc(), true, ctx.BaseParserRuleContext),
			CommentExpr: v.getDoc(ctx.GetComment(), false, ctx.BaseParserRuleContext),
		},
	}
}

func (v *ApiVisitor) VisitImportBlock(ctx *api.ImportBlockContext) interface{} {
	importToken := v.newExprWithToken(ctx.GetImportToken())
	values := ctx.AllImportBlockValue()
	var list []*ImportExpr
	for _, value := range values {
		importExpr := value.Accept(v).(*ImportExpr)
		importExpr.Import = importToken
		list = append(list, importExpr)
	}

	return list
}

func (v *ApiVisitor) VisitImportBlockValue(ctx *api.ImportBlockValueContext) interface{} {
	value := ctx.ImportValue().Accept(v).(Expr)
	doc := v.getDoc(ctx.GetDoc(), true, ctx.BaseParserRuleContext)
	comment := v.getDoc(ctx.GetComment(), false, ctx.BaseParserRuleContext)
	return &ImportExpr{
		Value:       value,
		DocExpr:     doc,
		CommentExpr: comment,
	}
}

func (v *ApiVisitor) VisitImportValue(ctx *api.ImportValueContext) interface{} {
	return v.newExprWithTerminalNode(ctx.STRING())
}

func (v *ApiVisitor) getDoc(ctx api.ICommentSpecContext, doc bool, current ...*antlr.BaseParserRuleContext) Expr {
	if ctx == nil {
		return nil
	}

	ret := ctx.Accept(v)
	if ret == nil {
		return nil
	}

	docExpr := ret.(Expr)
	if len(current) > 0 {
		line := current[0].GetStart().GetLine()
		if doc {
			line = line - 1
		}
		if docExpr.Line() != line {
			return nil
		}
	}
	return docExpr
}

func (i *ImportExpr) Format() error {
	// todo
	return nil
}

func (i *ImportExpr) Equal(v interface{}) bool {
	if v == nil {
		return false
	}

	imp, ok := v.(*ImportExpr)
	if !ok {
		return false
	}

	if !EqualDoc(i, imp) {
		return false
	}

	return i.Import.Equal(imp.Import) && i.Value.Equal(imp.Value)
}

func (i *ImportExpr) Doc() Expr {
	return i.DocExpr
}

func (i *ImportExpr) Comment() Expr {
	return i.CommentExpr
}
