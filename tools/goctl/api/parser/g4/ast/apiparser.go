package ast

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

type (
	Parser struct {
		prefix string
		debug  bool
		antlr.DefaultErrorListener
	}

	ParserOption func(p *Parser)
)

func NewParser(options ...ParserOption) *Parser {
	p := &Parser{}
	for _, opt := range options {
		opt(p)
	}

	return p
}

// Accept can parse any terminalNode of api tree by fn.
func (p *Parser) Accept(fn func(p *api.ApiParserParser, visitor *ApiVisitor) interface{}, content string) (v interface{}, err error) {
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
	lexer := api.NewApiParserLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	apiParser := api.NewApiParserParser(tokens)
	apiParser.RemoveErrorListeners()
	apiParser.AddErrorListener(p)
	visitor := NewApiVisitor(WithVisitorPrefix(p.prefix))
	v = fn(apiParser, visitor)
	return
}

// Parse is used to parse the api from the specified file name
func (p *Parser) Parse(filename string) (*spec.ApiSpec, error) {
	data, err := p.readContent(filename)
	if err != nil {
		return nil, err
	}

	return p.parse(filename, data)
}

// ParseContent is used to parse the api from the specified content
func (p *Parser) ParseContent(content string) (*spec.ApiSpec, error) {
	return p.parse("", content)
}

// parse is used to parse api from the content
// filename is only used to mark the file where the error is located
func (p *Parser) parse(filename, content string) (*spec.ApiSpec, error) {
	api, err := p.invoke(filename, content)
	if err != nil {
		return nil, err
	}

	var apiSpecs []*spec.ApiSpec
	imports := api.Import.List
	apiSpecs = append(apiSpecs, api)
	for _, imp := range imports {
		data, err := p.readContent(imp.Value)
		if err != nil {
			return nil, err
		}

		nestedApi, err := p.invoke(imp.Value, data)
		if err != nil {
			return nil, err
		}

		err = p.valid(api, imp.Value, nestedApi)
		if err != nil {
			return nil, err
		}

		apiSpecs = append(apiSpecs, nestedApi)
	}

	err = p.fillTypeMember(apiSpecs)
	if err != nil {
		return nil, err
	}

	allApi := p.memberFill(apiSpecs)
	return allApi, nil
}

func (p *Parser) invoke(filename, content string) (v *spec.ApiSpec, err error) {
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
	lexer := api.NewApiParserLexer(inputStream)
	lexer.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	apiParser := api.NewApiParserParser(tokens)
	apiParser.RemoveErrorListeners()
	apiParser.AddErrorListener(p)
	visitor := NewApiVisitor(WithVisitorPrefix(filepath.Base(filename)))

	v = apiParser.Api().Accept(visitor).(*spec.ApiSpec)
	return
}

func (p *Parser) valid(mainApi *spec.ApiSpec, filename string, nestedApi *spec.ApiSpec) error {
	if len(nestedApi.Import.List) > 0 {
		return fmt.Errorf("%s line %d:%d the nested api does not support import",
			filename, nestedApi.Import.List[0].Line, nestedApi.Import.List[0].Column)
	}

	if mainApi.Syntax.Version != nestedApi.Syntax.Version {
		return fmt.Errorf("%s line %d:%d multiple syntax declaration, expecting syntax '%s', but found '%s'",
			filename, nestedApi.Syntax.Line, nestedApi.Syntax.Column, mainApi.Syntax.Version, nestedApi.Syntax.Version)
	}

	if len(mainApi.Service.Name) != 0 && len(nestedApi.Service.Name) != 0 && mainApi.Service.Name != nestedApi.Service.Name {
		return fmt.Errorf("%s multiple service name declaration, expecting service name '%s', but found '%s'",
			filename, mainApi.Service.Name, nestedApi.Service.Name)
	}

	mainHandlerMap := make(map[string]PlaceHolder)
	mainRouteMap := make(map[string]PlaceHolder)
	mainTypeMap := make(map[string]PlaceHolder)

	routeMap := func(list []spec.Route) (map[string]PlaceHolder, map[string]PlaceHolder) {
		handlerMap := make(map[string]PlaceHolder)
		routeMap := make(map[string]PlaceHolder)

		for _, g := range list {
			handlerMap[g.Handler] = holder
			routeMap[g.Method+g.Path] = holder
		}

		return handlerMap, routeMap
	}

	h, r := routeMap(mainApi.Service.Routes())

	for k, v := range h {
		mainHandlerMap[k] = v
	}

	for k, v := range r {
		mainRouteMap[k] = v
	}

	for _, each := range mainApi.Types {
		mainTypeMap[each.Name] = holder
	}

	// duplicate route check
	for _, r := range nestedApi.Service.Routes() {
		if _, ok := mainHandlerMap[r.Handler]; ok {
			return fmt.Errorf("%s line %d:%d duplicate handler '%s'",
				filename, r.HandlerLineColumn.Line, r.HandlerLineColumn.Column, r.Handler)
		}

		if _, ok := mainRouteMap[r.Method+r.Path]; ok {
			return fmt.Errorf("%s line %d:%d duplicate route '%s'",
				filename, r.Line, r.Column, r.Method+" "+r.Path)
		}
	}

	// duplicate type check
	for _, each := range nestedApi.Types {
		if _, ok := mainTypeMap[each.Name]; ok {
			return fmt.Errorf("%s line %d:%d duplicate type declaration '%s'",
				filename, each.Line, each.Column, each.Name)
		}
	}
	return nil
}

func (p *Parser) memberFill(apiList []*spec.ApiSpec) *spec.ApiSpec {
	var api spec.ApiSpec

	for index, each := range apiList {
		if index == 0 {
			api.Syntax = each.Syntax
			api.Filename = each.Filename
			api.Info = each.Info
			api.Import = each.Import
			api.Service.Name = each.Service.Name
		}

		api.Types = append(api.Types, each.Types...)
		api.Service.Groups = append(api.Service.Groups, each.Service.Groups...)
	}

	return &api
}

func (p *Parser) fillTypeMember(apiList []*spec.ApiSpec) error {
	types := make(map[string]spec.Type)

	for _, api := range apiList {
		for _, each := range api.Types {
			types[each.Name] = each
		}
	}

	for _, api := range apiList {
		filename := api.Filename
		prefix := filepath.Base(filename)

		for _, each := range api.Types {
			for _, member := range each.Members {
				expr, err := p.fillType(prefix, types, member.Expr)
				if err != nil {
					return err
				}
				member.Expr = expr
			}
		}

		for _, each := range api.Service.Routes() {
			if len(each.RequestType.Name) > 0 {
				r, ok := types[each.RequestType.Name]
				if !ok {
					return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
						prefix, each.RequestType.Line, each.RequestType.Column, each.RequestType.Name)
				}
				each.RequestType.Members = r.Members
			}

			if len(each.ResponseType.Name) > 0 {
				r, ok := types[each.ResponseType.Name]
				if !ok {
					return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
						prefix, each.ResponseType.Line, each.ResponseType.Column, each.ResponseType.Name)
				}
				each.ResponseType.Members = r.Members
			}
		}
	}
	return nil
}

func (p *Parser) fillType(prefix string, types map[string]spec.Type, expr interface{}) (interface{}, error) {
	if expr == nil {
		return expr, nil
	}

	switch v := expr.(type) {
	case spec.Type:
		name := v.Name
		r, ok := types[name]
		if !ok {
			return nil, fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
				prefix, v.Line, v.Column, name)
		}

		v.Members = r.Members
		return v, nil
	case spec.PointerType:
		pointerExpr, err := p.fillType(prefix, types, v.Star)
		if err != nil {
			return nil, err
		}

		v.Star = pointerExpr
		return v, nil
	case spec.MapType:
		value, err := p.fillType(prefix, types, v.Value)
		if err != nil {
			return nil, err
		}

		v.Value = value
		return v, nil
	case spec.ArrayType:
		arrayType, err := p.fillType(prefix, types, v.ArrayType)
		if err != nil {
			return nil, err
		}

		v.ArrayType = arrayType
		return v, nil
	default:
		return expr, nil
	}
}

func (p *Parser) readContent(filename string) (string, error) {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(abs)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p *Parser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	str := fmt.Sprintf(`%s line %d:%d  %s`, p.prefix, line, column, msg)
	if p.debug {
		fmt.Println("[debug]", str)
	}
	panic(str)
}

var ParserDebug = WithParserDebug()

func WithParserDebug() ParserOption {
	return func(p *Parser) {
		p.debug = true
	}
}

func WithParserPrefix(prefix string) ParserOption {
	return func(p *Parser) {
		p.prefix = prefix
	}
}
