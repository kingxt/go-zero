package ast

import (
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/meta"
)

type MetaVisitor struct {
	parser.BaseMetaParserVisitor
	filename string
	baseLine int
}

func NewMetaVisitor(filename string, baseLine int) *MetaVisitor {
	return &MetaVisitor{
		filename: filename,
		baseLine: baseLine,
	}
}

func (v *MetaVisitor) VisitKv(ctx *parser.KvContext) interface{} {
	iKvLitContexts := ctx.AllKvLit()
	var spec KVSpec
	for _, kvLitContext := range iKvLitContexts {
		kv := kvLitContext.Accept(v).(KV)
		spec.List = append(spec.List, kv)
	}
	return &spec
}

func (v *MetaVisitor) VisitKvLit(ctx *parser.KvLitContext) interface{} {
	var kv KV
	key := ctx.GetKey()
	value := ctx.GetValue()
	if key != nil {
		kv.Key = Token{
			Text:   key.GetText(),
			Line:   v.baseLine + key.GetLine(),
			Column: key.GetColumn(),
		}
	}

	if value != nil {
		text := value.GetText()
		kv.Value = Token{
			Text:   text,
			Line:   v.baseLine + value.GetLine(),
			Column: value.GetColumn(),
		}
	}
	return kv
}
