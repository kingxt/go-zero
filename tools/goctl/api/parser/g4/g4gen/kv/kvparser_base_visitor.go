// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/KVParser.g4 by ANTLR 4.9. DO NOT EDIT.

package kv // KVParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseKVParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKVParserVisitor) VisitKv(ctx *KvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKVParserVisitor) VisitKvLit(ctx *KvLitContext) interface{} {
	return v.VisitChildren(ctx)
}
