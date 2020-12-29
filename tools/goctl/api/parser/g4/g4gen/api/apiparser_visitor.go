// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/ApiParser.g4 by ANTLR 4.9. DO NOT EDIT.

package api // ApiParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ApiParserParser.
type ApiParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ApiParserParser#api.
	VisitApi(ctx *ApiContext) interface{}

	// Visit a parse tree produced by ApiParserParser#spec.
	VisitSpec(ctx *SpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#syntaxLit.
	VisitSyntaxLit(ctx *SyntaxLitContext) interface{}

	// Visit a parse tree produced by ApiParserParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#importLit.
	VisitImportLit(ctx *ImportLitContext) interface{}

	// Visit a parse tree produced by ApiParserParser#importBlock.
	VisitImportBlock(ctx *ImportBlockContext) interface{}

	// Visit a parse tree produced by ApiParserParser#importBlockValue.
	VisitImportBlockValue(ctx *ImportBlockValueContext) interface{}

	// Visit a parse tree produced by ApiParserParser#importValue.
	VisitImportValue(ctx *ImportValueContext) interface{}

	// Visit a parse tree produced by ApiParserParser#infoSpec.
	VisitInfoSpec(ctx *InfoSpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#kvLit.
	VisitKvLit(ctx *KvLitContext) interface{}

	// Visit a parse tree produced by ApiParserParser#commentSpec.
	VisitCommentSpec(ctx *CommentSpecContext) interface{}
}
