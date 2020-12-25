// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/MetaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package meta // MetaParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MetaParser.
type MetaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MetaParser#kv.
	VisitKv(ctx *KvContext) interface{}

	// Visit a parse tree produced by MetaParser#kvLit.
	VisitKvLit(ctx *KvLitContext) interface{}
}
