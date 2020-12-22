package ast

import (
	"fmt"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

type (
	Parser struct {
		options []option
	}

	option func(p *parser.ApiParser)
)

func NewParser(options ...option) *Parser {
	instance := &Parser{
		options: options,
	}
	return instance
}

// Accept can parse any terminalNode of api tree by fn.
func (p *Parser) Accept(content string, fn func(p *parser.ApiParser, visitor *ApiVisitor) interface{}) (api interface{}, err error) {
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
	lexer := parser.NewApiLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	apiParser := parser.NewApiParser(tokens)
	visitor := NewApiVisitor("")
	p.options = append(p.options, WithErrorCallback("", nil))
	if len(p.options) > 0 {
		for _, opt := range p.options {
			opt(apiParser)
		}
	}

	api = fn(apiParser, visitor)
	return
}

// Parse parse the api file from root
func (p *Parser) Parse(filename string) (*spec.ApiSpec, error) {
	return p.parse(filename)
}

func (p *Parser) parse(filename string) (api *spec.ApiSpec, err error) {
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

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	inputStream := antlr.NewInputStream(string(data))
	lexer := parser.NewApiLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	apiParser := parser.NewApiParser(tokens)
	visitor := NewApiVisitor(filename)
	p.options = append(p.options, WithErrorCallback(filename, nil))
	if len(p.options) > 0 {
		for _, opt := range p.options {
			opt(apiParser)
		}
	}

	api = apiParser.Api().Accept(visitor).(*spec.ApiSpec)
	return
}

func WithErrorCallback(filename string, callback ErrCallback) option {
	return func(p *parser.ApiParser) {
		p.RemoveErrorListeners()
		errListener := NewErrorListener(filename, callback)
		p.AddErrorListener(errListener)
	}
}
