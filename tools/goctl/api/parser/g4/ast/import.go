package ast

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"

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
			DocExpr:     v.getDoc(ctx.GetDoc()),
			CommentExpr: v.getDoc(ctx.GetComment()),
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
	doc := v.getDoc(ctx.GetDoc())
	comment := v.getDoc(ctx.GetComment())
	return &ImportExpr{
		Value:       value,
		DocExpr:     doc,
		CommentExpr: comment,
	}
}

func (v *ApiVisitor) VisitImportValue(ctx *api.ImportValueContext) interface{} {
	return v.newExprWithTerminalNode(ctx.STRING())
}

func (v *ApiVisitor) getDoc(ctx api.ICommentSpecContext) Expr {
	if ctx == nil {
		return nil
	}

	ret := ctx.Accept(v)
	if ret == nil {
		return nil
	}

	return ret.(Expr)
}

func (i *ImportExpr) Format() error {
	// todo
	return nil
}

func (i *ImportExpr) Equal(v interface{}) bool {
	imp, ok := v.(*ImportExpr)
	if !ok {
		return false
	}

	if !EqualDoc(i, imp) {
		return false
	}

	return ExprEqual(i.Import, imp.Import) &&
		ExprEqual(i.Value, imp.Value)
}

func (i *ImportExpr) Doc() Expr {
	return i.DocExpr
}

func (i *ImportExpr) Comment() Expr {
	return i.CommentExpr
}
