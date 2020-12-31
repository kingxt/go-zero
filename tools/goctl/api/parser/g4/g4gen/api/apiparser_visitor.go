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

	// Visit a parse tree produced by ApiParserParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeLit.
	VisitTypeLit(ctx *TypeLitContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeBlock.
	VisitTypeBlock(ctx *TypeBlockContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeBody.
	VisitTypeBody(ctx *TypeBodyContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeStruct.
	VisitTypeStruct(ctx *TypeStructContext) interface{}

	// Visit a parse tree produced by ApiParserParser#typeAlias.
	VisitTypeAlias(ctx *TypeAliasContext) interface{}

	// Visit a parse tree produced by ApiParserParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by ApiParserParser#normalField.
	VisitNormalField(ctx *NormalFieldContext) interface{}

	// Visit a parse tree produced by ApiParserParser#anonymousFiled.
	VisitAnonymousFiled(ctx *AnonymousFiledContext) interface{}

	// Visit a parse tree produced by ApiParserParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by ApiParserParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by ApiParserParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ApiParserParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by ApiParserParser#kvLit.
	VisitKvLit(ctx *KvLitContext) interface{}

	// Visit a parse tree produced by ApiParserParser#commentSpec.
	VisitCommentSpec(ctx *CommentSpecContext) interface{}
}
