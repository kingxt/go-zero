package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/kv"
)

type (
	KVParser struct {
		options []kvOption
	}

	kvOption func(p *parser.KVParser)
)

func NewKVParser() *KVParser {
	return &KVParser{}
}

func (p *KVParser) Accept(baseLine int, content string, fn func(p *parser.KVParser, visitor *KVVisitor) interface{}) (kv interface{}, err error) {
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
	lexer := parser.NewKVLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	kvParser := parser.NewKVParser(tokens)
	visitor := NewKVVisitor("", baseLine)

	p.options = append(p.options, WithKVErrorCallback(baseLine, "", nil))
	for _, opt := range p.options {
		opt(kvParser)
	}

	kv = fn(kvParser, visitor)
	return
}

func (p *KVParser) Parse(baseLine int, filename string, content string) (kv *KVSpec, err error) {
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
	lexer := parser.NewKVLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	kvParser := parser.NewKVParser(tokens)
	visitor := NewKVVisitor(filename, baseLine)
	p.options = append(p.options, WithKVErrorCallback(baseLine, filename, nil))
	for _, opt := range p.options {
		opt(kvParser)
	}

	kv = kvParser.Kv().Accept(visitor).(*KVSpec)
	kv.Filename = filename
	return
}

func WithKVErrorCallback(baseLine int, filename string, callback ErrCallback) kvOption {
	return func(p *parser.KVParser) {
		p.RemoveErrorListeners()
		errListener := NewErrorListener(baseLine, filename, callback)
		p.AddErrorListener(errListener)
	}
}
