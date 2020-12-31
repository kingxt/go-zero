// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/AnonymousParser.g4 by ANTLR 4.9. DO NOT EDIT.

package anonymous // AnonymousParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by AnonymousParserParser.
type AnonymousParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AnonymousParserParser#anonymousFiled.
	VisitAnonymousFiled(ctx *AnonymousFiledContext) interface{}

	// Visit a parse tree produced by AnonymousParserParser#commentSpec.
	VisitCommentSpec(ctx *CommentSpecContext) interface{}
}
