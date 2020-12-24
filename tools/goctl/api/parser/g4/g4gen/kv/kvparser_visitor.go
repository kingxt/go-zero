// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/KVParser.g4 by ANTLR 4.9. DO NOT EDIT.

package kv // KVParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by KVParser.
type KVParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KVParser#kv.
	VisitKv(ctx *KvContext) interface{}

	// Visit a parse tree produced by KVParser#kvLit.
	VisitKvLit(ctx *KvLitContext) interface{}
}
