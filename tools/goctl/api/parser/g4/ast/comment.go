package ast

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"

func (v *ApiVisitor) VisitCommentSpec(ctx *api.CommentSpecContext) interface{} {
	if ctx.LINE_COMMENT() != nil {
		return v.newExprWithTerminalNode(ctx.LINE_COMMENT())
	}
	return v.newExprWithTerminalNode(ctx.COMMENT())
}
