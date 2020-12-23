package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const (
	handlerKey   = "handler"
	infoToken    = "info"
	syntaxToken  = "syntax"
	serviceToken = "service"
	summaryToken = "summary"
	returnToken  = "returns"
)

var (
	goTypeToken = map[string]struct{}{
		"bool":       struct{}{},
		"uint8":      struct{}{},
		"uint16":     struct{}{},
		"uint32":     struct{}{},
		"uint64":     struct{}{},
		"int8":       struct{}{},
		"int16":      struct{}{},
		"int32":      struct{}{},
		"int64":      struct{}{},
		"int":        struct{}{},
		"float32":    struct{}{},
		"float64":    struct{}{},
		"complex64":  struct{}{},
		"complex128": struct{}{},
		"string":     struct{}{},
		"uint":       struct{}{},
		"uintptr":    struct{}{},
		"byte":       struct{}{},
		"rune":       struct{}{},
		"time.Time":  struct{}{},
	}
	httpMethodToken = map[string]struct{}{
		"get":     struct{}{},
		"head":    struct{}{},
		"post":    struct{}{},
		"put":     struct{}{},
		"patch":   struct{}{},
		"delete":  struct{}{},
		"connect": struct{}{},
		"options": struct{}{},
		"trace":   struct{}{},
	}
)

type (
	ApiVisitor struct {
		parser.BaseApiParserVisitor
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
		v    spec.Info
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

func (v *ApiVisitor) VisitApi(ctx *parser.ApiContext) interface{} {
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

func (v *ApiVisitor) VisitBody(ctx *parser.BodyContext) interface{} {
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

func (v *ApiVisitor) VisitSyntaxLit(ctx *parser.SyntaxLitContext) interface{} {
	v.checkToken(ctx.GetSyntaxToken(), syntaxToken)
	version := v.getTokenText(ctx.GetVersion(), true)
	return spec.ApiSyntax{Version: version}
}

func (v *ApiVisitor) VisitImportSpec(ctx *parser.ImportSpecContext) interface{} {
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

func (v *ApiVisitor) VisitImportLit(ctx *parser.ImportLitContext) interface{} {
	importPath := v.getTokenText(ctx.GetImportPath(), true)
	line := ctx.GetImportPath().GetLine()
	column := ctx.GetImportPath().GetColumn()
	if _, ok := v.importSet[importPath]; ok {
		panic(v.wrapError(
			ast{line: line, column: column},
			`duplicate import "%s"`, importPath),
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

func (v *ApiVisitor) VisitImportLitGroup(ctx *parser.ImportLitGroupContext) interface{} {
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
			}, fmt.Sprintf(`duplicate import "%s"`, importPath)))
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

func (v *ApiVisitor) VisitInfoBlock(ctx *parser.InfoBlockContext) interface{} {
	v.checkToken(ctx.GetInfoToken(), infoToken)
	var info spec.Info
	info.Proterties = make(map[string]string)
	iKvLitContexts := ctx.AllKvLit()
	for _, each := range iKvLitContexts {
		kv := each.Accept(v).(kv)
		if _, ok := info.Proterties[kv.key]; ok {
			panic(v.wrapError(kv.ast, fmt.Sprintf(`duplicate info key "%s"`, kv.key)))
		}

		info.Proterties[kv.key] = kv.value
	}

	line := ctx.GetInfoToken().GetLine()
	column := ctx.GetInfoToken().GetColumn()
	info.Line = line
	info.Column = column

	if v.infoAst.flag {
		panic(v.wrapError(ast{
			line:   line,
			column: column,
		}, "duplicate info block"))
	}

	v.infoAst.flag = true
	v.infoAst.v = info
	v.infoAst.ast = ast{
		line:   line,
		column: column,
	}
	return info
}

func (v *ApiVisitor) VisitTypeBlock(ctx *parser.TypeBlockContext) interface{} {
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

func (v *ApiVisitor) VisitTypeLit(ctx *parser.TypeLitContext) interface{} {
	iTypeSpecContext := ctx.TypeSpec()
	return iTypeSpecContext.Accept(v).(spec.Type)
}

func (v *ApiVisitor) VisitTypeGroup(ctx *parser.TypeGroupContext) interface{} {
	iTypeSpecContexts := ctx.AllTypeSpec()
	var list []spec.Type
	for _, each := range iTypeSpecContexts {
		tp := each.Accept(v).(spec.Type)
		list = append(list, tp)
	}
	return list
}

func (v *ApiVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext) interface{} {
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

func (v *ApiVisitor) VisitTypeAlias(ctx *parser.TypeAliasContext) interface{} {
	line := ctx.GetAlias().GetLine()
	column := ctx.GetAlias().GetColumn()
	// todo: to support the alias types in the feature
	panic(v.wrapError(ast{
		line:   line,
		column: column,
	}, "unsupport alias"))
}

func (v *ApiVisitor) VisitTypeStruct(ctx *parser.TypeStructContext) interface{} {
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
		typeFieldContext := each.(*parser.TypeFieldContext)
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

func (v *ApiVisitor) VisitTypeField(ctx *parser.TypeFieldContext) interface{} {
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

func (v *ApiVisitor) VisitFiled(ctx *parser.FiledContext) interface{} {
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

func (v *ApiVisitor) VisitInnerStruct(ctx *parser.InnerStructContext) interface{} {
	symbol := ctx.LBRACE().GetSymbol()
	line := symbol.GetLine()
	column := symbol.GetColumn()
	panic(v.wrapError(ast{
		line:   line,
		column: column,
	}, "nested type is not supported"))
}

func (v *ApiVisitor) VisitDataType(ctx *parser.DataTypeContext) interface{} {
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

func (v *ApiVisitor) VisitMapType(ctx *parser.MapTypeContext) interface{} {
	tp := spec.MapType{}
	tp.Key = v.getTokenText(ctx.GetKey(), false)
	iDataTypeContext := ctx.DataType()
	tp.Value = iDataTypeContext.Accept(v)
	tp.StringExpr = ctx.GetText()
	tp.Line = ctx.MAP().GetSymbol().GetLine()
	tp.Column = ctx.MAP().GetSymbol().GetColumn()
	return tp
}

func (v *ApiVisitor) VisitArrayType(ctx *parser.ArrayTypeContext) interface{} {
	tp := spec.ArrayType{}
	iDataTypeContext := ctx.DataType()
	tp.ArrayType = iDataTypeContext.Accept(v)
	tp.StringExpr = ctx.GetText()
	tp.Line = ctx.LBRACK().GetSymbol().GetLine()
	tp.Column = ctx.LBRACK().GetSymbol().GetColumn()
	return tp
}

func (v *ApiVisitor) VisitPointer(ctx *parser.PointerContext) interface{} {
	if len(ctx.AllSTAR()) == 0 { // basic type
		if ctx.GOTYPE() != nil {
			text := v.getNodeText(ctx.GOTYPE(), false)
			if text == "time.Time" {
				tp := spec.TimeType{}
				tp.StringExpr = text
				return tp
			} else {
				tp := spec.BasicType{}
				tp.StringExpr = ctx.GetText()
				tp.Name = text
				return tp
			}
		} else if ctx.ID() != nil {
			tp := spec.Type{}
			tp.Name = v.getNodeText(ctx.ID(), false)
			if tp.Name == "interface" {
				symbol := ctx.ID().GetSymbol()
				panic(v.wrapError(ast{
					line:   symbol.GetLine(),
					column: symbol.GetColumn(),
				}, "expected '{'"))
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
		text := v.getNodeText(ctx.GOTYPE(), false)
		if text == "time.Time" {
			tp := spec.TimeType{}
			tp.StringExpr = text
			tmp.Star = tp
		} else {
			tp := spec.BasicType{}
			tp.StringExpr = text
			tp.Name = text
			tmp.Star = tp
		}
	} else if ctx.ID() != nil {
		tp := spec.Type{}
		tp.Name = v.getNodeText(ctx.ID(), false)
		if tp.Name == "interface" {
			symbol := ctx.ID().GetSymbol()
			panic(v.wrapError(ast{
				line:   symbol.GetLine(),
				column: symbol.GetColumn(),
			}, "unexpected interface"))
		}
		tmp.Star = tp
	}
	return *parent
}

func (v *ApiVisitor) VisitServiceBlock(ctx *parser.ServiceBlockContext) interface{} {
	var serviceGroup spec.Group
	if ctx.ServerMeta() != nil {
		annotation := ctx.ServerMeta().Accept(v).(spec.Annotation)
		serviceGroup.Annotation = annotation
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

func (v *ApiVisitor) VisitServerMeta(ctx *parser.ServerMetaContext) interface{} {
	var annotation spec.Annotation
	annotation.Properties = make(map[string]string, 0)
	annos := ctx.AllAnnotation()

	duplicate := make(map[string]struct{})
	for _, anno := range annos {
		kv := anno.Accept(v).(kv)
		if _, ok := duplicate[kv.key]; ok {
			panic(v.wrapError(kv.ast, `duplicate key "%s"`, kv.key))
		}

		duplicate[kv.key] = struct{}{}
		annotation.Properties[kv.key] = kv.value
	}

	return annotation
}

func (v *ApiVisitor) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	key := v.getTokenText(ctx.GetKey(), true)

	if len(key) == 0 || ctx.GetValue() == nil {
		panic(v.wrapError(ast{
			line:   ctx.GetKey().GetLine(),
			column: ctx.GetKey().GetColumn(),
		}, "empty annotation key or value"))
	}

	line := ctx.GetKey().GetLine()
	column := ctx.GetKey().GetColumn()

	return kv{
		key:   key,
		value: ctx.GetValue().GetText(),
		ast:   ast{line: line, column: column},
	}
}

func (v *ApiVisitor) VisitAnnotationKeyValue(ctx *parser.AnnotationKeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitServiceBody(ctx *parser.ServiceBodyContext) interface{} {
	v.checkToken(ctx.GetServiceToken(), serviceToken)
	var body serviceBody
	name := strings.TrimSpace(ctx.ServiceName().GetText())
	if len(name) == 0 {
		panic(v.filename + " service name should not null")
	}

	serviceNameContext := ctx.ServiceName().(*parser.ServiceNameContext)
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

func (v *ApiVisitor) VisitServiceName(ctx *parser.ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitServiceRoute(ctx *parser.ServiceRouteContext) interface{} {
	var route spec.Route
	iRouteDocContext := ctx.RouteDoc()
	iServerMetaContext := ctx.ServerMeta()
	iRouteHandlerContext := ctx.RouteHandler()
	iRoutePathContext := ctx.RoutePath()
	if iRouteDocContext != nil {
		route.Docs = []string{iRouteDocContext.Accept(v).(string)}
	}

	if iServerMetaContext != nil {
		metaContext := iServerMetaContext.(*parser.ServerMetaContext)
		for _, each := range metaContext.AllAnnotation() {
			key := v.getTokenText(each.GetKey(), false)
			if key == handlerKey {
				route.HandlerLineColumn.Line = each.GetKey().GetLine()
				route.HandlerLineColumn.Column = each.GetKey().GetColumn()
			}
		}
		annotation := iServerMetaContext.Accept(v).(spec.Annotation)
		route.Annotation = annotation
		route.Handler = annotation.Properties[handlerKey]
	} else if iRouteHandlerContext != nil {
		handlerContext := iRouteHandlerContext.(*parser.RouteHandlerContext)
		route.HandlerLineColumn.Line = handlerContext.ID().GetSymbol().GetLine()
		route.HandlerLineColumn.Column = handlerContext.ID().GetSymbol().GetColumn()
		route.Handler = iRouteHandlerContext.Accept(v).(string)
	}

	if iRoutePathContext != nil {
		routePathContext := iRoutePathContext.(*parser.RoutePathContext)
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

func (v *ApiVisitor) VisitRouteDoc(ctx *parser.RouteDocContext) interface{} {
	iLineDocContext := ctx.LineDoc()
	iDocContext := ctx.Doc()
	if iLineDocContext != nil {
		return iLineDocContext.Accept(v).(string)
	} else if iDocContext != nil {
		return iDocContext.Accept(v).(string)
	}

	return ""
}

func (v *ApiVisitor) VisitDoc(ctx *parser.DocContext) interface{} {
	v.checkToken(ctx.GetSummaryToken(), summaryToken)
	return v.getNodeText(ctx.STRING_LIT(), true)
}

func (v *ApiVisitor) VisitLineDoc(ctx *parser.LineDocContext) interface{} {
	return v.getNodeText(ctx.STRING_LIT(), true)
}

func (v *ApiVisitor) VisitRouteHandler(ctx *parser.RouteHandlerContext) interface{} {
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

func (v *ApiVisitor) VisitRoutePath(ctx *parser.RoutePathContext) interface{} {
	v.checkHttpMethod(ctx.GetHttpMethodToken())
	var routePath Route
	routePath.method = v.getTokenText(ctx.GetHttpMethodToken(), false)
	if ctx.Path() != nil {
		path := ctx.Path().GetText()
		pathContext := ctx.Path().(*parser.PathContext)
		line := pathContext.ID(0).GetSymbol().GetLine()
		column := pathContext.ID(0).GetSymbol().GetColumn()
		if _, ok := v.serviceAst.routeMap[routePath.method+path]; ok {
			panic(v.wrapError(ast{
				line:   line,
				column: column,
			}, `duplicate route path "%s"`, routePath.method+" "+path))
		}

		v.serviceAst.routeMap[routePath.method+path] = struct{}{}
		routePath.path = path
	}

	iRequestContext := ctx.Request()
	iReplyContext := ctx.Reply()
	if iRequestContext != nil {
		routePath.request = iRequestContext.Accept(v).(ReqReply)
	}

	if iReplyContext != nil {
		routePath.response = iReplyContext.Accept(v).(ReqReply)
	}

	return routePath
}
func (v *ApiVisitor) VisitPath(ctx *parser.PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitRequest(ctx *parser.RequestContext) interface{} {
	var ret ReqReply
	ret.name = v.getNodeText(ctx.ID(), false)
	ret.line = ctx.ID().GetSymbol().GetLine()
	ret.column = ctx.ID().GetSymbol().GetColumn()
	return ret
}

func (v *ApiVisitor) VisitReply(ctx *parser.ReplyContext) interface{} {
	v.checkToken(ctx.GetReturnToken(), returnToken)
	var ret ReqReply
	ret.name = v.getTokenText(ctx.GetObj(), false)
	ret.line = ctx.GetObj().GetLine()
	ret.column = ctx.GetObj().GetColumn()
	return ret
}

func (v *ApiVisitor) VisitKvLit(ctx *parser.KvLitContext) interface{} {
	key := v.getTokenText(ctx.GetKey(), false)
	value := v.getTokenText(ctx.GetValue(), true)
	line := ctx.GetKey().GetLine()
	column := ctx.GetKey().GetColumn()
	return kv{
		ast: ast{
			line:   line,
			column: column,
		},
		key:   key,
		value: value,
	}
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
		panic(fmt.Errorf("%sline %d:%d expected %s, but found %s",
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
		panic(fmt.Errorf("%s line %d:%d expected http method, but found %s",
			v.filename, token.GetLine(), token.GetColumn(), text))
	}
}
