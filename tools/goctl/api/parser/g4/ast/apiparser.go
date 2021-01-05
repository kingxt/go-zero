package ast

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
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
	var visitorOptions []VisitorOption
	visitorOptions = append(visitorOptions, WithVisitorPrefix(p.prefix))
	if p.debug {
		visitorOptions = append(visitorOptions, WithVisitorDebug())
	}
	visitor := NewApiVisitor(visitorOptions...)
	v = fn(apiParser, visitor)
	return
}

// Parse is used to parse the api from the specified file name
func (p *Parser) Parse(filename string) (*Api, error) {
	data, err := p.readContent(filename)
	if err != nil {
		return nil, err
	}

	return p.parse(filename, data)
}

// ParseContent is used to parse the api from the specified content
func (p *Parser) ParseContent(content string) (*Api, error) {
	return p.parse("", content)
}

// parse is used to parse api from the content
// filename is only used to mark the file where the error is located
func (p *Parser) parse(filename, content string) (*Api, error) {
	api, err := p.invoke(filename, content)
	if err != nil {
		return nil, err
	}

	var apiAstList []*Api
	apiAstList = append(apiAstList, api)
	for _, imp := range api.Import {
		path := imp.Value.Text()
		data, err := p.readContent(path)
		if err != nil {
			return nil, err
		}

		nestedApi, err := p.invoke(path, data)
		if err != nil {
			return nil, err
		}

		err = p.valid(api, path, nestedApi)
		if err != nil {
			return nil, err
		}

		apiAstList = append(apiAstList, nestedApi)
	}

	err = p.checkTypeDeclaration(apiAstList)
	if err != nil {
		return nil, err
	}

	allApi := p.memberFill(apiAstList)
	return allApi, nil
}

func (p *Parser) invoke(filename, content string) (v *Api, err error) {
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
	var visitorOptions []VisitorOption
	visitorOptions = append(visitorOptions, WithVisitorPrefix(filepath.Base(filename)))
	if p.debug {
		visitorOptions = append(visitorOptions, WithVisitorDebug())
	}
	visitor := NewApiVisitor(visitorOptions...)
	v = apiParser.Api().Accept(visitor).(*Api)
	v.Filename = filename
	return
}

func (p *Parser) valid(mainApi *Api, filename string, nestedApi *Api) error {
	if len(nestedApi.Import) > 0 {
		importToken := nestedApi.Import[0].Import
		return fmt.Errorf("%s line %d:%d the nested api does not support import",
			filename, importToken.Line(), importToken.Column())
	}

	if mainApi.Syntax.Version != nestedApi.Syntax.Version {
		syntaxToken := nestedApi.Syntax.Syntax
		return fmt.Errorf("%s line %d:%d multiple syntax declaration, expecting syntax '%s', but found '%s'",
			filename, syntaxToken.Line(), syntaxToken.Column(), mainApi.Syntax.Version, nestedApi.Syntax.Version)
	}

	if len(mainApi.Service) > 0 {
		mainService := mainApi.Service[0]
		for _, service := range nestedApi.Service {
			if mainService.ServiceApi.Name.Text() != service.ServiceApi.Name.Text() {
				return fmt.Errorf("%s multiple service name declaration, expecting service name '%s', but found '%s'",
					filename, mainService.ServiceApi.Name.Text(), service.ServiceApi.Name.Text())
			}
		}
	}

	mainHandlerMap := make(map[string]PlaceHolder)
	mainRouteMap := make(map[string]PlaceHolder)
	mainTypeMap := make(map[string]PlaceHolder)

	routeMap := func(list []*ServiceRoute) (map[string]PlaceHolder, map[string]PlaceHolder) {
		handlerMap := make(map[string]PlaceHolder)
		routeMap := make(map[string]PlaceHolder)

		for _, g := range list {
			var handlerName = g.GetHandler().Text()
			handlerMap[handlerName] = Holder
			path := fmt.Sprintf("%s://%s", g.Route.Method.Text(), g.Route.Path.Text())
			routeMap[path] = Holder
		}

		return handlerMap, routeMap
	}

	for _, each := range mainApi.Service {
		h, r := routeMap(each.ServiceApi.ServiceRoute)

		for k, v := range h {
			mainHandlerMap[k] = v
		}

		for k, v := range r {
			mainRouteMap[k] = v
		}
	}

	for _, each := range mainApi.Type {
		mainTypeMap[each.NameExpr().Text()] = Holder
	}

	// duplicate route check
	for _, each := range nestedApi.Service {
		for _, r := range each.ServiceApi.ServiceRoute {
			handler := r.GetHandler()
			if _, ok := mainHandlerMap[handler.Text()]; ok {
				return fmt.Errorf("%s line %d:%d duplicate handler '%s'",
					filename, handler.Line(), handler.Column(), handler.Text())
			}

			path := fmt.Sprintf("%s://%s", r.Route.Method.Text(), r.Route.Path.Text())
			if _, ok := mainRouteMap[path]; ok {
				return fmt.Errorf("%s line %d:%d duplicate route '%s'",
					filename, r.Route.Method.Line(), r.Route.Method.Column(), r.Route.Method.Text()+" "+r.Route.Path.Text())
			}
		}
	}

	// duplicate type check
	for _, each := range nestedApi.Type {
		if _, ok := mainTypeMap[each.NameExpr().Text()]; ok {
			return fmt.Errorf("%s line %d:%d duplicate type declaration '%s'",
				filename, each.NameExpr().Line(), each.NameExpr().Column(), each.NameExpr().Text())
		}
	}
	return nil
}

func (p *Parser) memberFill(apiList []*Api) *Api {
	var api Api

	for index, each := range apiList {
		if index == 0 {
			api.Syntax = each.Syntax
			api.Info = each.Info
			api.Import = each.Import
		}

		api.Type = append(api.Type, each.Type...)
		api.Service = append(api.Service, each.Service...)
	}

	return &api
}

// checkTypeDeclaration checks whether a struct type has been declared in context
func (p *Parser) checkTypeDeclaration(apiList []*Api) error {
	types := make(map[string]TypeExpr)

	for _, api := range apiList {
		for _, each := range api.Type {
			types[each.NameExpr().Text()] = each
		}
	}

	for _, api := range apiList {
		filename := api.Filename
		prefix := filepath.Base(filename)
		for _, each := range api.Type {
			tp, ok := each.(*TypeStruct)
			if !ok {
				continue
			}
			for _, member := range tp.Fields {
				err := p.checkType(prefix, types, member.DataType)
				if err != nil {
					return err
				}
			}
		}

		for _, service := range api.Service {
			for _, each := range service.ServiceApi.ServiceRoute {
				route := each.Route
				if route.Req != nil {
					_, ok := types[route.Req.Name.Text()]
					if !ok {
						return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
							prefix, route.Req.Name.Line(), route.Req.Name.Column(), route.Req.Name.Text())
					}
				}

				if route.Reply != nil {
					_, ok := types[route.Reply.Name.Text()]
					if !ok {
						return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
							prefix, route.Reply.Name.Line(), route.Reply.Name.Column(), route.Reply.Name.Text())
					}
				}
			}
		}
	}
	return nil
}

func (p *Parser) checkType(prefix string, types map[string]TypeExpr, expr DataType) error {
	if expr == nil {
		return nil
	}

	switch v := expr.(type) {
	case *Literal:
		name := v.Literal.Text()
		_, ok := types[name]
		if !ok {
			return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
				prefix, v.Literal.Line(), v.Literal.Column(), name)
		}

	case *Pointer:
		name := v.Name.Text()
		_, ok := types[name]
		if !ok {
			return fmt.Errorf("%s line %d:%d can not found declaration '%s' in context",
				prefix, v.Name.Line(), v.Name.Column(), name)
		}
	case *Map:
		return p.checkType(prefix, types, v.Value)
	case *Array:
		return p.checkType(prefix, types, v.Literal)
	default:
		return nil
	}
	return nil
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
