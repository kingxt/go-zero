package ast

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

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

// Parse parse the api file from the the root node
func (p *Parser) Parse(filename string) (*spec.ApiSpec, error) {
	api, err := p.parse(filename)
	if err != nil {
		return nil, err
	}

	imports := api.Import.List
	var apiSpecs []*spec.ApiSpec
	apiSpecs = append(apiSpecs, api)
	for _, imp := range imports {
		nestedApi, err := p.parse(imp.Value)
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
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(abs)
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
	api.Filename = filename
	return
}

func WithErrorCallback(filename string, callback ErrCallback) option {
	return func(p *parser.ApiParser) {
		p.RemoveErrorListeners()
		errListener := NewErrorListener(filename, callback)
		p.AddErrorListener(errListener)
	}
}

func (p *Parser) valid(mainApi *spec.ApiSpec, filename string, nestedApi *spec.ApiSpec) error {
	if len(nestedApi.Import.List) > 0 {
		return fmt.Errorf("%s %d:%d nested api does not support import",
			filename, nestedApi.Import.List[0].Line, nestedApi.Import.List[0].Column)
	}

	if mainApi.Syntax.Version != nestedApi.Syntax.Version {
		return fmt.Errorf("%s %d:%d multiple syntax, expected syntax %s, but found %s",
			filename, nestedApi.Syntax.Line, nestedApi.Syntax.Column, mainApi.Syntax.Version, nestedApi.Syntax.Version)
	}

	if mainApi.Service.Name != nestedApi.Service.Name {
		return fmt.Errorf("%s expected service name %s, but found %s",
			filename, mainApi.Service.Name, nestedApi.Service.Name)
	}

	mainHandlerMap := make(map[string]struct{})
	mainRouteMap := make(map[string]struct{})
	mainTypeMap := make(map[string]struct{})
	routeMap := func(list []spec.Route) (map[string]struct{}, map[string]struct{}) {
		handlerMap := make(map[string]struct{})
		routeMap := make(map[string]struct{})
		for _, g := range list {
			handlerMap[g.Handler] = struct{}{}
			routeMap[g.Method+g.Path] = struct{}{}
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
		mainTypeMap[each.Name] = struct{}{}
	}

	// duplicate route check
	for _, r := range nestedApi.Service.Routes() {
		if _, ok := mainHandlerMap[r.Handler]; ok {
			return fmt.Errorf("%s %d:%d duplicate handler %s",
				filename, r.HandlerLineColumn.Line, r.HandlerLineColumn.Column, r.Handler)
		}

		if _, ok := mainRouteMap[r.Method+r.Path]; ok {
			return fmt.Errorf("%s %d:%d duplicate route %s",
				filename, r.Line, r.Column, r.Method+" "+r.Path)
		}
	}

	// duplicate type check
	for _, each := range nestedApi.Types {
		if _, ok := mainTypeMap[each.Name]; ok {
			return fmt.Errorf("%s %d:%d duplicate type declaration %s",
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
					return fmt.Errorf("%s %d:%d can not found declaration %s in context",
						prefix, each.RequestType.Line, each.RequestType.Column, each.RequestType.Name)
				}

				each.RequestType.Members = r.Members

				if len(each.ResponseType.Name) > 0 {
					r, ok = types[each.ResponseType.Name]
					if !ok {
						return fmt.Errorf("%s %d:%d can not found declaration %s in context",
							prefix, each.ResponseType.Line, each.ResponseType.Column, each.ResponseType.Name)
					}

					each.ResponseType.Members = r.Members
				}
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
			return nil, fmt.Errorf("%s %d:%d can not found declaration %s in context",
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
