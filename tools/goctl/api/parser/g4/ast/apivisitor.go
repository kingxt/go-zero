package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const (
	handlerKey   = "handler"
	syntaxToken  = "syntax"
	serviceToken = "service"
	summaryToken = "summary"
	returnToken  = "returns"
)

var (
	goTypeToken = map[string]PlaceHolder{
		"bool":       hodler,
		"uint8":      hodler,
		"uint16":     hodler,
		"uint32":     hodler,
		"uint64":     hodler,
		"int8":       hodler,
		"int16":      hodler,
		"int32":      hodler,
		"int64":      hodler,
		"int":        hodler,
		"float32":    hodler,
		"float64":    hodler,
		"complex64":  hodler,
		"complex128": hodler,
		"string":     hodler,
		"uint":       hodler,
		"uintptr":    hodler,
		"byte":       hodler,
		"rune":       hodler,
		"time.Time":  hodler,
	}

	httpMethodToken = map[string]PlaceHolder{
		"get":     hodler,
		"head":    hodler,
		"post":    hodler,
		"put":     hodler,
		"patch":   hodler,
		"delete":  hodler,
		"connect": hodler,
		"options": hodler,
		"trace":   hodler,
	}
)

type (
	ApiVisitor struct {
		api.BaseApiParserVisitor
		apiSpec    spec.ApiSpec
		importSet  map[string]importAst
		infoAst    infoAst
		typeMap    map[string]structAst
		serviceAst serviceAst
		filename   string
	}

	ast struct {
		line   int
		column int
	}

	importAst struct {
		ast
		v string
	}

	infoAst struct {
		ast
		flag bool
	}

	structAst struct {
		ast
		v spec.Type
	}

	serviceAst struct {
		ast
		name       string
		handlerMap map[string]struct{}
		routeMap   map[string]struct{}
	}

	kv struct {
		ast
		key   string
		value string
	}

	serviceBody struct {
		name string
		ast
		routes []spec.Route
	}

	ReqReply struct {
		ast
		name string
	}

	Route struct {
		method   string
		path     string
		request  ReqReply
		response ReqReply
	}
)

func NewApiVisitor(filename string) *ApiVisitor {
	return &ApiVisitor{
		importSet: make(map[string]importAst),
		typeMap:   make(map[string]structAst),
		serviceAst: serviceAst{
			handlerMap: map[string]struct{}{},
			routeMap:   map[string]struct{}{},
		},
		filename: filename,
	}
}

func (v *ApiVisitor) VisitApi(ctx *api.ApiContext) interface{} {
	var api spec.ApiSpec
	iSyntaxLitContext := ctx.SyntaxLit()
	iBodyContexts := ctx.AllBody()
	if iSyntaxLitContext != nil {
		api.Syntax = iSyntaxLitContext.Accept(v).(spec.ApiSyntax)
	}

	for _, item := range iBodyContexts {
		body := item.Accept(v)
		if body == nil {
			continue
		}

		switch v := body.(type) {
		case spec.ApiImport:
			api.Import.List = append(api.Import.List, v.List...)
		case spec.Info:
			api.Info = v
		case []spec.Type:
			api.Types = append(api.Types, v...)
		case spec.Service:
			api.Service.Name = v.Name
			api.Service.Groups = append(api.Service.Groups, v.Groups...)
		default:
			continue
		}
	}

	return &api
}

func (v *ApiVisitor) VisitBody(ctx *api.BodyContext) interface{} {
	iImportSpecContext := ctx.ImportSpec()
	iInfoBlockContext := ctx.InfoBlock()
	iTypeBlockContext := ctx.TypeBlock()
	iServiceBlockContext := ctx.ServiceBlock()
	if iImportSpecContext != nil {
		return iImportSpecContext.Accept(v)
	} else if iInfoBlockContext != nil {
		return iInfoBlockContext.Accept(v)
	} else if iTypeBlockContext != nil {
		return iTypeBlockContext.Accept(v)
	} else if iServiceBlockContext != nil {
		return iServiceBlockContext.Accept(v)
	} else {
		return nil
	}
}

func (v *ApiVisitor) VisitSyntaxLit(ctx *api.SyntaxLitContext) interface{} {
	v.checkToken(ctx.GetSyntaxToken(), syntaxToken)
	version := v.getTokenText(ctx.GetVersion(), true)
	return spec.ApiSyntax{Version: version}
}

func (v *ApiVisitor) VisitImportSpec(ctx *api.ImportSpecContext) interface{} {
	iImportLitContext := ctx.ImportLit()
	iImportLitGroupContext := ctx.ImportLitGroup()
	var list []spec.Import
	if iImportLitContext != nil {
		importValue := iImportLitContext.Accept(v).(spec.ApiImport)
		list = append(list, importValue.List...)
	}

	if iImportLitGroupContext != nil {
		importValue := iImportLitGroupContext.Accept(v).(spec.ApiImport)
		list = append(list, importValue.List...)
	}

	return spec.ApiImport{List: list}
}

func (v *ApiVisitor) VisitImportLit(ctx *api.ImportLitContext) interface{} {
	importPath := v.getTokenText(ctx.GetImportPath(), true)
	line := ctx.GetImportPath().GetLine()
	column := ctx.GetImportPath().GetColumn()
	if _, ok := v.importSet[importPath]; ok {
		panic(v.wrapError(
			ast{line: line, column: column},
			`duplicate import '%s'`, importPath),
		)
	}

	v.importSet[importPath] = importAst{
		ast: ast{
			line:   line,
			column: column,
		},
		v: importPath,
	}

	return spec.ApiImport{
		List: []spec.Import{
			{
				LineColumn: spec.LineColumn{
					Line:   line,
					Column: column,
				},
				Value: importPath,
			},
		},
	}
}

func (v *ApiVisitor) VisitImportLitGroup(ctx *api.ImportLitGroupContext) interface{} {
	nodes := ctx.AllIMPORT_PATH()
	var list []spec.Import
	for _, node := range nodes {
		importPath := v.getNodeText(node, true)
		line := node.GetSymbol().GetLine()
		column := node.GetSymbol().GetColumn()
		if _, ok := v.importSet[importPath]; ok {
			panic(v.wrapError(ast{
				line:   line,
				column: column,
			}, fmt.Sprintf(`duplicate import '%s'`, importPath)))
		}

		v.importSet[importPath] = importAst{
			ast: ast{
				line:   line,
				column: column,
			},
			v: importPath,
		}

		list = append(list, spec.Import{
			LineColumn: spec.LineColumn{
				Line:   line,
				Column: column,
			},
			Value: importPath,
		})
	}

	return spec.ApiImport{List: list}
}

func (v *ApiVisitor) VisitInfoBlock(ctx *api.InfoBlockContext) interface{} {
	content := ctx.INFO_BLOCK().GetText()
	line := ctx.INFO_BLOCK().GetSymbol().GetLine()
	column := ctx.INFO_BLOCK().GetSymbol().GetColumn()
	kvParser := NewKVParser()
	content = strings.TrimPrefix(content, "info")
	kvSpec, err := kvParser.Parse(line-1, v.filename, content)
	if err != nil {
		panic(err)
	}

	if v.infoAst.flag {
		panic(v.wrapError(ast{
			line:   line,
			column: column,
		}, "duplicate info block"))
	}

	properties := make(map[string]string)
	for _, each := range kvSpec.List {
		key := each.Key
		if _, ok := properties[key.Text]; ok {
			panic(v.wrapError(ast{
				line:   key.Line,
				column: key.Column,
			}, "duplicate key '%s' in info block", key.Text))
		}

		properties[each.Key.Text] = each.Value.Text
	}

	v.infoAst.flag = true
	v.infoAst.ast = ast{
		line:   line,
		column: column,
	}

	return spec.Info{
		LineColumn: spec.LineColumn{
			Line:   line,
			Column: column,
		},
		Proterties: properties,
	}
}

func (v *ApiVisitor) VisitTypeBlock(ctx *api.TypeBlockContext) interface{} {
	iTypeLitContext := ctx.TypeLit()
	iTypeGroupContext := ctx.TypeGroup()
	var list []spec.Type
	if iTypeLitContext != nil {
		tp := iTypeLitContext.Accept(v).(spec.Type)
		list = append(list, tp)
	} else if iTypeGroupContext != nil {
		types := iTypeGroupContext.Accept(v).([]spec.Type)
		list = append(list, types...)
	}

	return list
}

func (v *ApiVisitor) VisitTypeLit(ctx *api.TypeLitContext) interface{} {
	iTypeSpecContext := ctx.TypeSpec()
	return iTypeSpecContext.Accept(v).(spec.Type)
}

func (v *ApiVisitor) VisitTypeGroup(ctx *api.TypeGroupContext) interface{} {
	iTypeSpecContexts := ctx.AllTypeSpec()
	var list []spec.Type
	for _, each := range iTypeSpecContexts {
		tp := each.Accept(v).(spec.Type)
		list = append(list, tp)
	}
	return list
}

func (v *ApiVisitor) VisitTypeSpec(ctx *api.TypeSpecContext) interface{} {
	var tp spec.Type
	iTypeAliasContext := ctx.TypeAlias()
	iTypeStructContext := ctx.TypeStruct()
	if iTypeAliasContext != nil {
		return iTypeAliasContext.Accept(v)
	} else if iTypeStructContext != nil {
		return iTypeStructContext.Accept(v)
	}
	return tp
}

func (v *ApiVisitor) VisitTypeAlias(ctx *api.TypeAliasContext) interface{} {
	line := ctx.GetAlias().GetLine()
	column := ctx.GetAlias().GetColumn()
	// todo: to support the alias types in the feature
	panic(v.wrapError(ast{
		line:   line,
		column: column,
	}, "unexpecting alias"))
}

func (v *ApiVisitor) VisitTypeStruct(ctx *api.TypeStructContext) interface{} {
	var tp spec.Type
	tp.Name = v.getTokenText(ctx.GetName(), false)
	line := ctx.GetName().GetLine()
	column := ctx.GetName().GetColumn()
	tp.Line = line
	tp.Column = column

	if _, ok := v.typeMap[tp.Name]; ok {
		panic(v.wrapError(ast{
			line:   line,
			column: column,
		}, `duplicate type "%s"`, tp.Name))
	}

	iTypeFieldContexts := ctx.AllTypeField()
	set := make(map[string]struct{})
	for _, each := range iTypeFieldContexts {
		member := each.Accept(v).(spec.Member)
		typeFieldContext := each.(*api.TypeFieldContext)
		line := typeFieldContext.GetName().GetLine()
		column := typeFieldContext.GetName().GetColumn()
		member.Line = line
		member.Column = column
		if _, ok := set[member.Name]; ok {
			panic(v.wrapError(ast{
				line:   line,
				column: column,
			}, `duplicate filed "%s"`, member.Name))
		}

		set[member.Name] = struct{}{}
		tp.Members = append(tp.Members, member)
	}

	v.typeMap[tp.Name] = structAst{
		ast: ast{
			line:   line,
			column: column,
		},
		v: tp,
	}

	return tp
}

func (v *ApiVisitor) VisitTypeField(ctx *api.TypeFieldContext) interface{} {
	var member spec.Member
	iFiledContext := ctx.Filed()
	if iFiledContext != nil {
		member = iFiledContext.Accept(v).(spec.Member)
	} else { // anonymousType
		member.IsInline = true
	}
	member.Name = v.getTokenText(ctx.GetName(), false)
	return member
}

func (v *ApiVisitor) VisitFiled(ctx *api.FiledContext) interface{} {
	iDataTypeContext := ctx.DataType()
	iInnerStructContext := ctx.InnerStruct()
	tag := v.getTokenText(ctx.GetTag(), false)
	tag = strings.ReplaceAll(tag, "`", "")
	// todo: tag valid?
	var tp interface{}
	if iDataTypeContext != nil {
		dataTypeResult := iDataTypeContext.Accept(v)
		filed := spec.Member{}
		switch v := dataTypeResult.(type) {
		case spec.BasicType:
			filed.Type = v.Name
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.PointerType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.MapType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.ArrayType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.InterfaceType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.TimeType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case spec.Type:
			filed.Type = v.Name
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		default:
			return tp
		}
	} else if iInnerStructContext != nil {
		iInnerStructContext.Accept(v)
	}
	return tp
}

func (v *ApiVisitor) VisitInnerStruct(ctx *api.InnerStructContext) interface{} {
	symbol := ctx.LBRACE().GetSymbol()
	line := symbol.GetLine()
	column := symbol.GetColumn()
	panic(v.wrapError(ast{
		line:   line,
		column: column,
	}, "nested type is not supported"))
}

func (v *ApiVisitor) VisitDataType(ctx *api.DataTypeContext) interface{} {
	iPointerContext := ctx.Pointer()
	iMapTypeContext := ctx.MapType()
	iArrayTypeContext := ctx.ArrayType()
	interfaceNode := ctx.INTERFACE()
	var tp interface{}
	if iPointerContext != nil {
		return iPointerContext.Accept(v)
	} else if iMapTypeContext != nil {
		return iMapTypeContext.Accept(v)
	} else if iArrayTypeContext != nil {
		return iArrayTypeContext.Accept(v)
	} else if interfaceNode != nil {
		return spec.InterfaceType{StringExpr: ctx.GetText()}
	}
	return tp
}

func (v *ApiVisitor) VisitMapType(ctx *api.MapTypeContext) interface{} {
	tp := spec.MapType{}
	tp.Key = v.getTokenText(ctx.GetKey(), false)
	iDataTypeContext := ctx.DataType()
	tp.Value = iDataTypeContext.Accept(v)
	tp.StringExpr = ctx.GetText()
	tp.Line = ctx.MAP().GetSymbol().GetLine()
	tp.Column = ctx.MAP().GetSymbol().GetColumn()
	return tp
}

func (v *ApiVisitor) VisitArrayType(ctx *api.ArrayTypeContext) interface{} {
	tp := spec.ArrayType{}
	iDataTypeContext := ctx.DataType()
	tp.ArrayType = iDataTypeContext.Accept(v)
	tp.StringExpr = ctx.GetText()
	tp.Line = ctx.LBRACK().GetSymbol().GetLine()
	tp.Column = ctx.LBRACK().GetSymbol().GetColumn()
	return tp
}

func (v *ApiVisitor) VisitPointer(ctx *api.PointerContext) interface{} {
	if len(ctx.AllSTAR()) == 0 { // basic type
		if ctx.GOTYPE() != nil {
			line := ctx.GOTYPE().GetSymbol().GetLine()
			column := ctx.GOTYPE().GetSymbol().GetColumn()
			text := v.getNodeText(ctx.GOTYPE(), false)
			if text == "time.Time" {
				tp := spec.TimeType{}
				tp.StringExpr = text
				tp.Line = line
				tp.Column = column
				return tp
			} else {
				tp := spec.BasicType{}
				tp.StringExpr = ctx.GetText()
				tp.Name = text
				tp.Line = line
				tp.Column = column
				return tp
			}
		} else if ctx.ID() != nil {
			tp := spec.Type{}
			line := ctx.ID().GetSymbol().GetLine()
			column := ctx.ID().GetSymbol().GetColumn()
			tp.Name = v.getNodeText(ctx.ID(), false)
			tp.Line = line
			tp.Column = column
			if tp.Name == "interface" {
				symbol := ctx.ID().GetSymbol()
				panic(v.wrapError(ast{
					line:   symbol.GetLine(),
					column: symbol.GetColumn(),
				}, "expecting '{'"))
			}
			return tp
		}
	}

	// pointer
	text := ctx.GetText()
	parent := &spec.PointerType{
		StringExpr: text,
	}
	tmp := parent
	for index := 1; index < len(ctx.AllSTAR()); index++ {
		p := &spec.PointerType{
			StringExpr: text[index:],
			Star:       nil,
		}
		tmp.Star = p
		tmp = p
	}

	if ctx.GOTYPE() != nil {
		line := ctx.GOTYPE().GetSymbol().GetLine()
		column := ctx.GOTYPE().GetSymbol().GetColumn()
		text := v.getNodeText(ctx.GOTYPE(), false)
		if text == "time.Time" {
			tp := spec.TimeType{}
			tp.StringExpr = text
			tp.Line = line
			tp.Column = column
			tmp.Star = tp
		} else {
			tp := spec.BasicType{}
			tp.StringExpr = text
			tp.Name = text
			tp.Line = line
			tp.Column = column
			tmp.Star = tp
		}
	} else if ctx.ID() != nil {
		tp := spec.Type{}
		line := ctx.ID().GetSymbol().GetLine()
		column := ctx.ID().GetSymbol().GetColumn()
		tp.Line = line
		tp.Column = column
		tp.Name = v.getNodeText(ctx.ID(), false)
		if tp.Name == "interface" {
			symbol := ctx.ID().GetSymbol()
			panic(v.wrapError(ast{
				line:   symbol.GetLine(),
				column: symbol.GetColumn(),
			}, "unexpecting interface"))
		}
		tmp.Star = tp
	}
	return *parent
}

func (v *ApiVisitor) VisitServiceBlock(ctx *api.ServiceBlockContext) interface{} {
	var serviceGroup spec.Group
	if ctx.ServerMeta() != nil {
		kv := ctx.ServerMeta().Accept(v).(*KVSpec)
		if kv != nil {
			properties := make(map[string]string)
			for _, each := range kv.List {
				properties[each.Key.Text] = each.Value.Text
			}

			annotation := spec.Annotation{Properties: properties}
			serviceGroup.Annotation = annotation
		}
	}

	body := ctx.ServiceBody().Accept(v).(serviceBody)
	serviceGroup.Routes = body.routes
	if v.serviceAst.name != "" && body.name != v.serviceAst.name {
		panic(v.wrapError(body.ast, `multiple service name "%s"`, body.name))
	}

	v.serviceAst.name = body.name
	v.serviceAst.line = body.line
	v.serviceAst.column = body.column
	return spec.Service{
		Name:   body.name,
		Groups: []spec.Group{serviceGroup},
	}
}

func (v *ApiVisitor) VisitServerMeta(ctx *api.ServerMetaContext) interface{} {
	serverMeta := ctx.SERVER_META_STRING()
	if serverMeta == nil {
		return nil
	}

	line := serverMeta.GetSymbol().GetLine()
	kvParser := NewKVParser()
	content := serverMeta.GetText()
	content = strings.TrimPrefix(content, "@server")
	kv, err := kvParser.Parse(line-1, v.filename, content)
	if err != nil {
		panic(err)
	}

	return kv
}

func (v *ApiVisitor) VisitServiceBody(ctx *api.ServiceBodyContext) interface{} {
	v.checkToken(ctx.GetServiceToken(), serviceToken)
	var body serviceBody
	name := strings.TrimSpace(ctx.ServiceName().GetText())
	if len(name) == 0 {
		panic(v.filename + " service name should not null")
	}

	serviceNameContext := ctx.ServiceName().(*api.ServiceNameContext)
	line := serviceNameContext.ID(0).GetSymbol().GetLine()
	column := serviceNameContext.ID(0).GetSymbol().GetColumn()

	body.line = line
	body.column = column
	body.name = name
	for _, item := range ctx.AllServiceRoute() {
		r := item.Accept(v).(spec.Route)
		body.routes = append(body.routes, r)
	}

	return body
}

func (v *ApiVisitor) VisitServiceName(ctx *api.ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitServiceRoute(ctx *api.ServiceRouteContext) interface{} {
	var route spec.Route
	iRouteDocContext := ctx.RouteDoc()
	iServerMetaContext := ctx.ServerMeta()
	iRouteHandlerContext := ctx.RouteHandler()
	iRoutePathContext := ctx.RoutePath()
	if iRouteDocContext != nil {
		route.Docs = []string{iRouteDocContext.Accept(v).(string)}
	}

	if iServerMetaContext != nil {
		kvSpec := iServerMetaContext.Accept(v).(*KVSpec)
		if kvSpec != nil {
			properties := make(map[string]string)
			for _, each := range kvSpec.List {
				key := each.Key.Text
				properties[key] = each.Value.Text
				if key == handlerKey {
					route.Handler = each.Value.Text
					route.HandlerLineColumn.Line = each.Key.Line
					route.HandlerLineColumn.Column = each.Key.Column
				}
			}

			route.Annotation = spec.Annotation{Properties: properties}
		}
	} else if iRouteHandlerContext != nil {
		handlerContext := iRouteHandlerContext.(*api.RouteHandlerContext)
		route.HandlerLineColumn.Line = handlerContext.ID().GetSymbol().GetLine()
		route.HandlerLineColumn.Column = handlerContext.ID().GetSymbol().GetColumn()
		route.Handler = iRouteHandlerContext.Accept(v).(string)
	}

	if iRoutePathContext != nil {
		routePathContext := iRoutePathContext.(*api.RoutePathContext)
		route.Line = routePathContext.GetHttpMethodToken().GetLine()
		route.Column = routePathContext.GetHttpMethodToken().GetColumn()
		r := iRoutePathContext.Accept(v).(Route)
		route.Method = r.method
		route.Path = r.path
		route.RequestType = spec.Type{
			Name: r.request.name,
			LineColumn: spec.LineColumn{
				Line:   r.request.line,
				Column: r.request.column,
			},
		}
		route.ResponseType = spec.Type{
			Name: r.response.name,
			LineColumn: spec.LineColumn{
				Line:   r.response.line,
				Column: r.response.column,
			},
		}
	}

	return route
}

func (v *ApiVisitor) VisitRouteDoc(ctx *api.RouteDocContext) interface{} {
	iLineDocContext := ctx.LineDoc()
	iDocContext := ctx.Doc()
	if iLineDocContext != nil {
		return iLineDocContext.Accept(v).(string)
	} else if iDocContext != nil {
		return iDocContext.Accept(v).(string)
	}

	return ""
}

func (v *ApiVisitor) VisitDoc(ctx *api.DocContext) interface{} {
	v.checkToken(ctx.GetSummaryToken(), summaryToken)
	return v.getNodeText(ctx.STRING_LIT(), true)
}

func (v *ApiVisitor) VisitLineDoc(ctx *api.LineDocContext) interface{} {
	return v.getNodeText(ctx.STRING_LIT(), true)
}

func (v *ApiVisitor) VisitRouteHandler(ctx *api.RouteHandlerContext) interface{} {
	text := v.getNodeText(ctx.ID(), false)
	if _, ok := v.serviceAst.handlerMap[text]; ok {
		panic(v.wrapError(ast{
			line:   ctx.ID().GetSymbol().GetLine(),
			column: ctx.ID().GetSymbol().GetColumn(),
		}, `duplicate handler "%s"`, text))
	}

	v.serviceAst.handlerMap[text] = struct{}{}
	return text
}

func (v *ApiVisitor) VisitRoutePath(ctx *api.RoutePathContext) interface{} {
	v.checkHttpMethod(ctx.GetHttpMethodToken())
	var routePath Route
	routePath.method = v.getTokenText(ctx.GetHttpMethodToken(), false)
	line := ctx.GetHttpMethodToken().GetLine()
	column := ctx.GetHttpMethodToken().GetColumn()

	if ctx.Path() != nil {
		path := ctx.Path().GetText()
		if _, ok := v.serviceAst.routeMap[routePath.method+path]; ok {
			panic(v.wrapError(ast{
				line:   line,
				column: column,
			}, `duplicate route path "%s"`, routePath.method+" "+path))
		}

		v.serviceAst.routeMap[routePath.method+path] = struct{}{}
		routePath.path = path
	}

	iRequestContext := ctx.GetReq()
	iReplyContext := ctx.GetReply()
	if iRequestContext != nil {
		req := iRequestContext.Accept(v).(*ReqReply)
		if len(req.name) != 0 {
			routePath.request = *req
		}
	}

	if iReplyContext != nil {
		reply := iReplyContext.Accept(v).(*ReqReply)
		if len(reply.name) != 0 {
			routePath.response = *reply
		}
	}

	return routePath
}

func (v *ApiVisitor) VisitHttpBody(ctx *api.HttpBodyContext) interface{} {
	obj := ctx.GetObj()
	var ret ReqReply
	if obj != nil {
		ret.name = v.getTokenText(obj, false)
		ret.line = obj.GetLine()
		ret.column = obj.GetColumn()
	}
	return &ret
}

func (v *ApiVisitor) getTokenInt(token antlr.Token) (int64, error) {
	text := v.getTokenText(token, true)
	if len(text) == 0 {
		return 0, nil
	}

	vInt, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, err
	}

	return vInt, nil
}

func (v *ApiVisitor) getTokenText(token antlr.Token, trimQuote bool) string {
	if token == nil {
		return ""
	}

	text := token.GetText()
	if trimQuote {
		text = v.trimQuote(text)
	}
	return text
}

func (v *ApiVisitor) getNodeInt(node antlr.TerminalNode) (int64, error) {
	text := v.getNodeText(node, true)
	if len(text) == 0 {
		return 0, nil
	}

	vInt, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, err
	}

	return vInt, nil
}

func (v *ApiVisitor) getNodeText(node antlr.TerminalNode, trimQuote bool) string {
	if node == nil {
		return ""
	}

	text := node.GetText()
	if trimQuote {
		text = v.trimQuote(text)
	}
	return text
}

func (v *ApiVisitor) trimQuote(text string) string {
	text = strings.ReplaceAll(text, `"`, "")
	text = strings.ReplaceAll(text, `'`, "")
	text = strings.ReplaceAll(text, "`", "")
	return text
}

func (v *ApiVisitor) wrapError(ast ast, format string, a ...interface{}) error {
	if v.filename != "" {
		v.filename = v.filename + " "
	}
	return fmt.Errorf("%s line %d:%d %s", v.filename, ast.line, ast.column, fmt.Sprintf(format, a...))
}

func (v *ApiVisitor) checkToken(token antlr.Token, text string) {
	if token == nil {
		return
	}
	tokenText := v.getTokenText(token, false)
	if tokenText != text {
		if len(v.filename) > 0 {
			v.filename = v.filename + " "
		}
		panic(fmt.Errorf("%sline %d:%d expecting %s, but found %s",
			v.filename, token.GetLine(), token.GetColumn(), text, tokenText))
	}
}

func (v *ApiVisitor) isGoType(text string) bool {
	_, ok := goTypeToken[text]
	return ok
}

func (v *ApiVisitor) checkHttpMethod(token antlr.Token) {
	if token == nil {
		return
	}
	text := v.getTokenText(token, false)
	_, ok := httpMethodToken[text]
	if !ok {
		if len(v.filename) > 0 {
			v.filename = v.filename + " "
		}
		panic(fmt.Errorf("%sline %d:%d expecting http method, but found %s",
			v.filename, token.GetLine(), token.GetColumn(), text))
	}
}
