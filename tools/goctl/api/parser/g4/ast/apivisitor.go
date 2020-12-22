package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const handlerKey = "handler"

type (
	ApiVisitor struct {
		parser.BaseApiParserVisitor
		apiSpec spec.ApiSpec
	}
	kv struct {
		key   string
		value string
	}
	serviceBody struct {
		name   string
		routes []spec.Route
	}
	Route struct {
		method   string
		path     string
		request  string
		response string
	}
)

func NewApiVisitor() *ApiVisitor {
	return &ApiVisitor{}
}

func (v *ApiVisitor) VisitApi(ctx *parser.ApiContext) interface{} {
	v.VisitChildren(ctx)
	return v.apiSpec
}

func (v *ApiVisitor) VisitBody(ctx *parser.BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitSyntaxLit(ctx *parser.SyntaxLitContext) interface{} {
	version := v.getTokenText(ctx.GetVersion(), true)

	return &spec.ApiSyntax{Version: version}
}

func (v *ApiVisitor) VisitImportSpec(ctx *parser.ImportSpecContext) interface{} {
	iImportLitContext := ctx.ImportLit()
	iImportLitGroupContext := ctx.ImportLitGroup()
	var list []string
	if iImportLitContext != nil {
		importLitContext, ok := iImportLitContext.(*parser.ImportLitContext)
		if ok {
			result := v.VisitImportLit(importLitContext)
			importValue, ok := result.(*spec.ApiImport)
			if ok {
				list = append(list, importValue.List...)
			}
		}
	}

	if iImportLitGroupContext != nil {
		importGroupContext, ok := iImportLitGroupContext.(*parser.ImportLitGroupContext)
		if ok {
			result := v.VisitImportLitGroup(importGroupContext)
			importValue, ok := result.(*spec.ApiImport)
			if ok {
				list = append(list, importValue.List...)
			}
		}
	}
	return &spec.ApiImport{List: list}
}

func (v *ApiVisitor) VisitImportLit(ctx *parser.ImportLitContext) interface{} {
	importPath := v.getTokenText(ctx.GetImportPath(), true)

	return &spec.ApiImport{
		List: []string{importPath},
	}
}

func (v *ApiVisitor) VisitImportLitGroup(ctx *parser.ImportLitGroupContext) interface{} {
	nodes := ctx.AllIMPORT_PATH()
	var list []string
	for _, node := range nodes {
		importPath := v.getNodeText(node, true)

		list = append(list, importPath)
	}
	return &spec.ApiImport{List: list}
}

func (v *ApiVisitor) VisitInfoBlock(ctx *parser.InfoBlockContext) interface{} {
	var info spec.Info
	info.Proterties = make(map[string]string)
	iKvLitContexts := ctx.AllKvLit()
	for _, each := range iKvLitContexts {
		kvLitContext, ok := each.(*parser.KvLitContext)
		if !ok {
			continue
		}

		r := v.VisitKvLit(kvLitContext)
		kv, ok := r.(*kv)
		if !ok {
			continue
		}
		info.Proterties[kv.key] = kv.value

	}
	return &info
}

func (v *ApiVisitor) VisitTypeBlock(ctx *parser.TypeBlockContext) interface{} {
	iTypeLitContext := ctx.TypeLit()
	iTypeGroupContext := ctx.TypeGroup()
	var list []*spec.Type
	if iTypeLitContext != nil {
		typeLitContext, ok := iTypeLitContext.(*parser.TypeLitContext)
		if !ok {
			return list
		}

		typeLit := v.VisitTypeLit(typeLitContext)
		tp := typeLit.(*spec.Type)
		list = append(list, tp)
	} else if iTypeGroupContext != nil {
		typeGroupContext, ok := iTypeGroupContext.(*parser.TypeGroupContext)
		if !ok {
			return list
		}

		typeGroup := v.VisitTypeGroup(typeGroupContext)
		types, ok := typeGroup.([]*spec.Type)
		if !ok {
			return list
		}

		list = append(list, types...)
	}

	return list
}

func (v *ApiVisitor) VisitTypeLit(ctx *parser.TypeLitContext) interface{} {
	var tp spec.Type
	iTypeSpecContext := ctx.TypeSpec()
	typeSpecContext, ok := iTypeSpecContext.(*parser.TypeSpecContext)
	if !ok {
		return tp
	}

	typeSpec := v.VisitTypeSpec(typeSpecContext)
	result, ok := typeSpec.(*spec.Type)
	if !ok {
		return tp
	}
	return result
}

func (v *ApiVisitor) VisitTypeGroup(ctx *parser.TypeGroupContext) interface{} {
	iTypeSpecContexts := ctx.AllTypeSpec()
	var list []*spec.Type
	for _, each := range iTypeSpecContexts {
		typeSpecContext, ok := each.(*parser.TypeSpecContext)
		if !ok {
			continue
		}

		typeSpec := v.VisitTypeSpec(typeSpecContext)
		tp, ok := typeSpec.(*spec.Type)
		if !ok {
			continue
		}

		list = append(list, tp)
	}
	return list
}

func (v *ApiVisitor) VisitTypeSpec(ctx *parser.TypeSpecContext) interface{} {
	var tp spec.Type
	iTypeAliasContext := ctx.TypeAlias()
	iTypeStructContext := ctx.TypeStruct()
	if iTypeAliasContext != nil {
		typeAliasContext, ok := iTypeAliasContext.(*parser.TypeAliasContext)
		if !ok {
			return &tp
		}

		return v.VisitTypeAlias(typeAliasContext)
	} else if iTypeStructContext != nil {
		structContext, ok := iTypeStructContext.(*parser.TypeStructContext)
		if !ok {
			return &tp
		}

		return v.VisitTypeStruct(structContext)
	}
	return &tp
}

func (v *ApiVisitor) VisitTypeAlias(ctx *parser.TypeAliasContext) interface{} {
	line := ctx.GetAlias().GetLine()
	column := ctx.GetAlias().GetColumn()
	// todo: to support the alias types in the feature
	panic(fmt.Errorf("line %d:%d unsupport alias", line, column))
}

func (v *ApiVisitor) VisitTypeStruct(ctx *parser.TypeStructContext) interface{} {
	var tp spec.Type
	tp.Name = v.getTokenText(ctx.GetName(), false)
	iTypeFieldContexts := ctx.AllTypeField()
	for _, each := range iTypeFieldContexts {
		fieldContext, ok := each.(*parser.TypeFieldContext)
		if !ok {
			continue
		}

		field := v.VisitTypeField(fieldContext)
		member, ok := field.(*spec.Member)
		if !ok {
			continue
		}

		tp.Members = append(tp.Members, *member)
	}
	return &tp
}

func (v *ApiVisitor) VisitTypeField(ctx *parser.TypeFieldContext) interface{} {
	var member spec.Member
	iFiledContext := ctx.Filed()
	if iFiledContext != nil {
		filedContext, ok := iFiledContext.(*parser.FiledContext)
		if ok {
			fieldResult := v.VisitFiled(filedContext)
			m, ok := fieldResult.(*spec.Member)
			if ok {
				member = *m
			}
		}
	} else { // anonymousType
		member.IsInline = true
	}
	member.Name = v.getTokenText(ctx.GetName(), false)
	return &member
}

func (v *ApiVisitor) VisitFiled(ctx *parser.FiledContext) interface{} {
	iDataTypeContext := ctx.DataType()
	iInnerStructContext := ctx.InnerStruct()
	tag := v.getTokenText(ctx.GetTag(), false)
	tag = strings.ReplaceAll(tag, "`", "")
	// todo: tag valid?
	var tp interface{}
	if iDataTypeContext != nil {
		dataTypeContext, ok := iDataTypeContext.(*parser.DataTypeContext)
		if !ok {
			return tp
		}

		dataTypeResult := v.VisitDataType(dataTypeContext)
		filed := &spec.Member{}
		switch v := dataTypeResult.(type) {
		case *spec.BasicType:
			filed.Type = v.Name
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case *spec.PointerType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case *spec.MapType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case *spec.ArrayType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		case *spec.InterfaceType:
			filed.Type = v.StringExpr
			filed.Expr = dataTypeResult
			filed.Tag = tag
			return filed
		default:
			return tp
		}
	} else if iInnerStructContext != nil {
		innerStructContext, ok := iInnerStructContext.(*parser.InnerStructContext)
		if !ok {
			return tp
		}
		symbol := innerStructContext.LBRACE().GetSymbol()
		line := symbol.GetLine()
		column := symbol.GetColumn()
		panic(fmt.Errorf("line %d:%d nested type not supported", line, column))
	}
	return tp
}

func (v *ApiVisitor) VisitInnerStruct(ctx *parser.InnerStructContext) interface{} {
	var list []*spec.Member
	iTypeFieldContexts := ctx.AllTypeField()
	for _, each := range iTypeFieldContexts {
		typeFieldContext, ok := each.(*parser.TypeFieldContext)
		if !ok {
			continue
		}

		result := v.VisitTypeField(typeFieldContext)
		field, ok := result.(*spec.Member)
		if !ok {
			continue
		}

		list = append(list, field)
	}
	return list
}

func (v *ApiVisitor) VisitDataType(ctx *parser.DataTypeContext) interface{} {
	iPointerContext := ctx.Pointer()
	iMapTypeContext := ctx.MapType()
	iArrayTypeContext := ctx.ArrayType()
	interfaceNode := ctx.INTERFACE()
	var tp interface{}
	if iPointerContext != nil {
		pointerContext, ok := iPointerContext.(*parser.PointerContext)
		if !ok {
			return tp
		}

		return v.VisitPointer(pointerContext)
	} else if iMapTypeContext != nil {
		mapTypeContext, ok := iMapTypeContext.(*parser.MapTypeContext)
		if !ok {
			return tp
		}

		return v.VisitMapType(mapTypeContext)
	} else if iArrayTypeContext != nil {
		arrayTypeContext, ok := iArrayTypeContext.(*parser.ArrayTypeContext)
		if !ok {
			return tp
		}

		return v.VisitArrayType(arrayTypeContext)
	} else if interfaceNode != nil {
		return &spec.InterfaceType{StringExpr: ctx.GetText()}
	}
	return tp
}

func (v *ApiVisitor) VisitMapType(ctx *parser.MapTypeContext) interface{} {
	tp := &spec.MapType{}
	key := v.getTokenText(ctx.GetKey(), false)
	iDataTypeContext := ctx.DataType()
	dataTypeContext, ok := iDataTypeContext.(*parser.DataTypeContext)
	if ok {
		dt := v.VisitDataType(dataTypeContext)
		tp.Key = key
		tp.Value = dt
		tp.StringExpr = ctx.GetText()
	}
	return tp
}

func (v *ApiVisitor) VisitArrayType(ctx *parser.ArrayTypeContext) interface{} {
	tp := &spec.ArrayType{}
	iDataTypeContext := ctx.DataType()
	dataTypeContext, ok := iDataTypeContext.(*parser.DataTypeContext)
	if ok {
		dt := v.VisitDataType(dataTypeContext)
		tp.ArrayType = dt
		tp.StringExpr = ctx.GetText()
	}
	return tp
}

func (v *ApiVisitor) VisitPointer(ctx *parser.PointerContext) interface{} {
	if len(ctx.AllSTAR()) == 0 { // basic type
		if ctx.GOTYPE() != nil {
			tp := &spec.BasicType{}
			tp.StringExpr = ctx.GetText()
			tp.Name = v.getNodeText(ctx.GOTYPE(), false)
			return tp
		} else if ctx.ID() != nil {
			tp := &spec.Type{}
			tp.Name = v.getNodeText(ctx.ID(), false)
			if tp.Name == "interface" {
				symbol := ctx.ID().GetSymbol()
				panic(fmt.Errorf("line %d:%d expected '{'", symbol.GetLine(), symbol.GetColumn()))
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
		tp := &spec.BasicType{}
		tp.StringExpr = v.getNodeText(ctx.GOTYPE(), false)
		tp.Name = v.getNodeText(ctx.GOTYPE(), false)
		tmp.Star = tp
	} else if ctx.ID() != nil {
		tp := &spec.Type{}
		tp.Name = v.getNodeText(ctx.ID(), false)
		if tp.Name == "interface" {
			symbol := ctx.ID().GetSymbol()
			panic(fmt.Errorf("line %d:%d unexpected interface", symbol.GetLine(), symbol.GetColumn()))
		}
		tmp.Star = tp
	}
	return parent
}

func (v *ApiVisitor) VisitServiceBlock(ctx *parser.ServiceBlockContext) interface{} {
	var serviceGroup spec.Group
	if ctx.ServerMeta() != nil {
		serviceGroup.Annotation = ctx.ServerMeta().Accept(v).(spec.Annotation)
	}
	body := ctx.ServiceBody().Accept(v).(serviceBody)
	serviceGroup.Routes = body.routes
	if len(v.apiSpec.Service.Name) > 0 && body.name != v.apiSpec.Service.Name {
		v.apiSpec.Service.Name = body.name
		panic(fmt.Sprintf("multi service name [%s, %s] should name equal", v.apiSpec.Service.Name, body.name))
	}

	v.apiSpec.Service.Groups = append(v.apiSpec.Service.Groups, serviceGroup)
	return serviceGroup
}

func (v *ApiVisitor) VisitServerMeta(ctx *parser.ServerMetaContext) interface{} {
	var annotation spec.Annotation
	annotation.Properties = make(map[string]string, 0)
	annos := ctx.AllAnnotation()
	for _, anno := range annos {
		kv := anno.Accept(v).(kv)
		annotation.Properties[kv.key] = kv.value
	}

	return annotation
}

func (v *ApiVisitor) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	key := v.getTokenText(ctx.GetKey(), true)

	if len(key) == 0 || ctx.GetValue() == nil {
		panic(errors.New("empty annotation key or value"))
	}

	return kv{key: key, value: ctx.GetValue().GetText()}
}

func (v *ApiVisitor) VisitAnnotationKeyValue(ctx *parser.AnnotationKeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitServiceBody(ctx *parser.ServiceBodyContext) interface{} {
	var body serviceBody
	name := strings.TrimSpace(ctx.ServiceName().GetText())
	if len(name) == 0 {
		panic("service name should notnull")
	}

	body.name = strings.TrimSpace(ctx.ServiceName().GetText())
	for _, item := range ctx.AllServiceRoute() {
		body.routes = append(body.routes, item.Accept(v).(spec.Route))
	}
	return body
}

func (v *ApiVisitor) VisitServiceName(ctx *parser.ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitServiceRoute(ctx *parser.ServiceRouteContext) interface{} {
	var route spec.Route
	if ctx.RouteHandler() != nil {
		route.Handler = ctx.RouteHandler().Accept(v).(string)
	}
	if ctx.RouteDoc() != nil {
		route.Docs = ctx.RouteDoc().Accept(v).(spec.Doc)
	}
	if ctx.ServerMeta() != nil {
		route.Annotation = ctx.ServerMeta().Accept(v).(spec.Annotation)
	}
	if len(route.Handler) == 0 {
		route.Handler = route.Annotation.Properties[handlerKey]
	}
	if len(route.Handler) == 0 {
		panic(fmt.Sprintf("route [%s] handler missing", route.Path))
	}

	var routePath = ctx.RoutePath().Accept(v).(Route)
	route.Method = routePath.method
	route.Path = routePath.path
	for _, ty := range v.apiSpec.Types {
		if ty.Name == routePath.request {
			route.RequestType = ty
		}
		if ty.Name == routePath.response {
			route.ResponseType = ty
		}
	}
	return route
}

func (v *ApiVisitor) VisitRouteDoc(ctx *parser.RouteDocContext) interface{} {
	var doc spec.Doc
	if ctx.LineDoc() != nil {
		doc = append(doc, ctx.LineDoc().Accept(v).(string))
	}
	if ctx.Doc() != nil {
		doc = append(doc, ctx.Doc().Accept(v).(spec.Doc)...)
	}
	return doc
}

func (v *ApiVisitor) VisitDoc(ctx *parser.DocContext) interface{} {
	var doc spec.Doc
	for _, item := range ctx.AllKvLit() {
		doc = append(doc, item.GetValue().GetText())
	}
	return doc
}

func (v *ApiVisitor) VisitLineDoc(ctx *parser.LineDocContext) interface{} {
	return ctx.STRING_LIT().GetText()
}

func (v *ApiVisitor) VisitRouteHandler(ctx *parser.RouteHandlerContext) interface{} {
	return ctx.ID().GetText()
}

func (v *ApiVisitor) VisitRoutePath(ctx *parser.RoutePathContext) interface{} {
	var routePath Route
	routePath.method = ctx.HTTPMETHOD().GetText()
	routePath.path = ctx.Path().GetText()
	if ctx.Request() != nil {
		routePath.request = ctx.Request().GetText()
	}
	if ctx.Reply() != nil {
		routePath.request = ctx.Reply().GetText()
	}
	return routePath
}
func (v *ApiVisitor) VisitPath(ctx *parser.PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitRequest(ctx *parser.RequestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitReply(ctx *parser.ReplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ApiVisitor) VisitKvLit(ctx *parser.KvLitContext) interface{} {
	key := v.getTokenText(ctx.GetKey(), false)
	value := v.getTokenText(ctx.GetValue(), true)

	return &kv{
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
