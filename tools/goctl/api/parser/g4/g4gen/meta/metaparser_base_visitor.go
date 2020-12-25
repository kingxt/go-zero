// Code generated from /Users/anqiansong/goland/go/go-zero_kingxt/tools/goctl/api/parser/g4/MetaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package meta // MetaParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMetaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMetaParserVisitor) VisitKv(ctx *KvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMetaParserVisitor) VisitKvLit(ctx *KvLitContext) interface{} {
	return v.VisitChildren(ctx)
}
