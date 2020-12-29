// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/ApiParser.g4 by ANTLR 4.9. DO NOT EDIT.

package api // ApiParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseApiParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseApiParserVisitor) VisitApi(ctx *ApiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitSpec(ctx *SpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitSyntaxLit(ctx *SyntaxLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitImportLit(ctx *ImportLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitImportBlock(ctx *ImportBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitImportBlockValue(ctx *ImportBlockValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitImportValue(ctx *ImportValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitInfoSpec(ctx *InfoSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitKvLit(ctx *KvLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseApiParserVisitor) VisitCommentSpec(ctx *CommentSpecContext) interface{} {
	return v.VisitChildren(ctx)
}
