package ast

import (
	"strings"
	"unicode"

	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/kv"
)

type Token struct {
	Text   string
	Line   int `json:"-"`
	Column int `json:"-"`
}

type KV struct {
	Key   Token
	Value Token
}

type KVSpec struct {
	List     []KV
	Filename string
}

type KVVisitor struct {
	parser.BaseKVParserVisitor
	filename string
	baseLine int
}

func NewKVVisitor(filename string, baseLine int) *KVVisitor {
	return &KVVisitor{
		filename: filename,
		baseLine: baseLine,
	}
}

func (v *KVVisitor) VisitKv(ctx *parser.KvContext) interface{} {
	iKvLitContexts := ctx.AllKvLit()
	var spec KVSpec
	for _, kvLitContext := range iKvLitContexts {
		kv := kvLitContext.Accept(v).(KV)
		spec.List = append(spec.List, kv)
	}
	return &spec
}

func (v *KVVisitor) VisitKvLit(ctx *parser.KvLitContext) interface{} {
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
		text = v.trimSpaceAndQuote(text)
		kv.Value = Token{
			Text:   text,
			Line:   v.baseLine + value.GetLine(),
			Column: value.GetColumn(),
		}
	}
	return kv
}

// trimSpaceAndQuote trim the prefix and the suffix character which is space or quote
func (v *KVVisitor) trimSpaceAndQuote(text string) string {
	text = strings.TrimFunc(text, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	text = strings.TrimPrefix(text, `"`)
	text = strings.TrimSuffix(text, `"`)
	return text
}
