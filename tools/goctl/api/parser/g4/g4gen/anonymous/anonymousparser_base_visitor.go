// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/AnonymousParser.g4 by ANTLR 4.9. DO NOT EDIT.

package anonymous // AnonymousParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseAnonymousParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAnonymousParserVisitor) VisitAnonymousFiled(ctx *AnonymousFiledContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAnonymousParserVisitor) VisitCommentSpec(ctx *CommentSpecContext) interface{} {
	return v.VisitChildren(ctx)
}
