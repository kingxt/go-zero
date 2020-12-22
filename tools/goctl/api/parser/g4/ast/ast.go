package ast

import (
	"fmt"

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
	options = append(options, WithErrorCallback(nil))
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
	visitor := NewApiVisitor()
	if len(p.options) > 0 {
		for _, opt := range p.options {
			opt(apiParser)
		}
	}

	api = fn(apiParser, visitor)
	return
}

// Parse parse the api file from root
func (p *Parser) Parse(content string) (*spec.ApiSpec, error) {
	// todo: Recursive parse if there are some api imports
	return p.parse(content)
}

func (p *Parser) parse(content string) (api *spec.ApiSpec, err error) {
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
	visitor := NewApiVisitor()
	if len(p.options) > 0 {
		for _, opt := range p.options {
			opt(apiParser)
		}
	}

	api = apiParser.Api().Accept(visitor).(*spec.ApiSpec)
	return
}

func WithErrorCallback(callback ErrCallback) option {
	return func(p *parser.ApiParser) {
		p.RemoveErrorListeners()
		errListener := NewErrorListener(callback)
		p.AddErrorListener(errListener)
	}
}
