package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/meta"
)

type (
	MetaParser struct {
		options []metaOption
	}

	metaOption func(p *parser.MetaParser)
)

func NewMetaParser() *MetaParser {
	return &MetaParser{}
}

func (p *MetaParser) Accept(baseLine int, content string, fn func(p *parser.MetaParser, visitor *MetaVisitor) interface{}) (kv interface{}, err error) {
	defer func() {
		p := recover()
		if p != nil {
			switch e := p.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("%+v", p)
			}
		}
	}()

	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewMetaLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	metaParser := parser.NewMetaParser(tokens)
	visitor := NewMetaVisitor("", baseLine)

	p.options = append(p.options, WithMetaErrorCallback(baseLine, "", nil))
	for _, opt := range p.options {
		opt(metaParser)
	}

	kv = fn(metaParser, visitor)
	return
}

func (p *MetaParser) Parse(baseLine int, filename string, content string) (kv *KVSpec, err error) {
	defer func() {
		p := recover()
		if p != nil {
			switch e := p.(type) {
			case error:
				err = e
			default:
				err = fmt.Errorf("%+v", p)
			}
		}
	}()

	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewMetaLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	MetaParser := parser.NewMetaParser(tokens)
	visitor := NewMetaVisitor(filename, baseLine)
	p.options = append(p.options, WithMetaErrorCallback(baseLine, filename, nil))
	for _, opt := range p.options {
		opt(MetaParser)
	}

	kv = MetaParser.Kv().Accept(visitor).(*KVSpec)
	kv.Filename = filename
	return
}

func WithMetaErrorCallback(baseLine int, filename string, callback ErrCallback) metaOption {
	return func(p *parser.MetaParser) {
		p.RemoveErrorListeners()
		errListener := NewErrorListener(baseLine, filename, callback)
		p.AddErrorListener(errListener)
	}
}
