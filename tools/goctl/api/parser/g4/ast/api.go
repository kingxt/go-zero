package ast

import (
	"fmt"
	"sort"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/gen/api"
)

type Api struct {
	LinePrefix string
	Syntax     *SyntaxExpr
	Import     []*ImportExpr
	importM    map[string]PlaceHolder
	Info       *InfoExpr
	Type       []TypeExpr
	typeM      map[string]PlaceHolder
	Service    []*Service
	serviceM   map[string]PlaceHolder
	handlerM   map[string]PlaceHolder
	routeM     map[string]PlaceHolder
}

func (v *ApiVisitor) VisitApi(ctx *api.ApiContext) interface{} {
	defer func() {
		if p := recover(); p != nil {
			panic(fmt.Errorf("%+v", p))
		}
	}()
	var final Api
	final.importM = map[string]PlaceHolder{}
	final.typeM = map[string]PlaceHolder{}
	final.serviceM = map[string]PlaceHolder{}
	final.handlerM = map[string]PlaceHolder{}
	final.routeM = map[string]PlaceHolder{}
	for _, each := range ctx.AllSpec() {
		api := each.Accept(v).(*Api)
		if api.Syntax != nil {
			if final.Syntax != nil {
				v.panic(api.Syntax.Syntax, fmt.Sprintf("mutiple syntax declaration"))
			}
			final.Syntax = api.Syntax
		}

		for _, imp := range api.Import {
			if _, ok := final.importM[imp.Value.Text()]; ok {
				v.panic(imp.Import, fmt.Sprintf("duplicate import '%s'", imp.Value.Text()))
			}
			final.importM[imp.Value.Text()] = Holder
			final.Import = append(final.Import, imp)
		}

		if api.Info != nil {
			infoM := map[string]PlaceHolder{}
			if final.Info != nil {
				v.panic(api.Info.Info, fmt.Sprintf("mutiple info declaration"))
			}
			for _, value := range api.Info.Kvs {
				if _, ok := infoM[value.Key.Text()]; ok {
					v.panic(value.Key, fmt.Sprintf("duplicate key '%s'", value.Key.Text()))
				}
				infoM[value.Key.Text()] = Holder
			}
			final.Info = api.Info
		}

		for _, tp := range api.Type {
			if _, ok := final.typeM[tp.NameExpr().Text()]; ok {
				v.panic(tp.NameExpr(), fmt.Sprintf("duplicate type '%s'", tp.NameExpr().Text()))
			}
			final.typeM[tp.NameExpr().Text()] = Holder
			final.Type = append(final.Type, tp)
		}

		for _, service := range api.Service {
			if _, ok := final.serviceM[service.ServiceApi.Name.Text()]; !ok && len(final.serviceM) > 0 {
				v.panic(service.ServiceApi.Name, fmt.Sprintf("mutiple service declaration"))
			}
			if service.AtServer != nil {
				atServerM := map[string]PlaceHolder{}
				for _, kv := range service.AtServer.Kv {
					if _, ok := atServerM[kv.Key.Text()]; ok {
						v.panic(kv.Key, fmt.Sprintf("duplicate key '%s'", kv.Key.Text()))
					}
					atServerM[kv.Key.Text()] = Holder
				}
			}
			for _, route := range service.ServiceApi.ServiceRoute {
				uniqueRoute := fmt.Sprintf("%s %s", route.Route.Method.Text(), route.Route.Path.Text())
				if _, ok := final.routeM[uniqueRoute]; ok {
					v.panic(route.Route.Method, fmt.Sprintf("duplicate route '%s'", uniqueRoute))
				}
				final.routeM[uniqueRoute] = Holder

				var handlerExpr Expr
				if route.AtServer != nil {
					atServerM := map[string]PlaceHolder{}
					for _, kv := range route.AtServer.Kv {
						if _, ok := atServerM[kv.Key.Text()]; ok {
							v.panic(kv.Key, fmt.Sprintf("duplicate key '%s'", kv.Key.Text()))
						}
						atServerM[kv.Key.Text()] = Holder
						if kv.Key.Text() == "handler" {
							handlerExpr = kv.Value
						}
					}
				}

				if route.AtHandler != nil {
					handlerExpr = route.AtHandler.Name
				}

				if handlerExpr == nil {
					v.panic(route.Route.Method, fmt.Sprintf("mismtached handler"))
				}

				if handlerExpr.Text() == "" {
					v.panic(handlerExpr, fmt.Sprintf("mismtached handler"))
				}

				if _, ok := final.handlerM[handlerExpr.Text()]; ok {
					v.panic(handlerExpr, fmt.Sprintf("duplicate handler '%s'", handlerExpr.Text()))
				}
				final.handlerM[handlerExpr.Text()] = Holder
			}
			final.Service = append(final.Service, service)
		}
	}

	return &final
}

func (v *ApiVisitor) VisitSpec(ctx *api.SpecContext) interface{} {
	var api Api
	if ctx.SyntaxLit() != nil {
		api.Syntax = ctx.SyntaxLit().Accept(v).(*SyntaxExpr)
	}
	if ctx.ImportSpec() != nil {
		api.Import = ctx.ImportSpec().Accept(v).([]*ImportExpr)
	}
	if ctx.InfoSpec() != nil {
		api.Info = ctx.InfoSpec().Accept(v).(*InfoExpr)
	}
	if ctx.TypeSpec() != nil {
		tp := ctx.TypeSpec().Accept(v)
		api.Type = tp.([]TypeExpr)
	}
	if ctx.ServiceSpec() != nil {
		api.Service = []*Service{ctx.ServiceSpec().Accept(v).(*Service)}
	}
	return &api
}

func (a *Api) Format() error {
	// todo
	return nil
}

func (a *Api) Equal(v interface{}) bool {
	if v == nil {
		return false
	}

	api, ok := v.(*Api)
	if !ok {
		return false
	}

	if !a.Syntax.Equal(api.Syntax) {
		return false
	}

	if len(a.Import) != len(api.Import) {
		return false
	}

	var expectingImport, actualImport []*ImportExpr
	expectingImport = append(expectingImport, a.Import...)
	actualImport = append(actualImport, api.Import...)

	sort.Slice(expectingImport, func(i, j int) bool {
		return expectingImport[i].Value.Text() < expectingImport[j].Value.Text()
	})

	sort.Slice(actualImport, func(i, j int) bool {
		return actualImport[i].Value.Text() < actualImport[j].Value.Text()
	})

	for index, each := range expectingImport {
		ac := actualImport[index]
		if !each.Equal(ac) {
			return false
		}
	}

	if !a.Info.Equal(api.Info) {
		return false
	}

	if len(a.Type) != len(api.Type) {
		return false
	}

	var expectingType, actualType []TypeExpr
	expectingType = append(expectingType, a.Type...)
	actualType = append(actualType, api.Type...)

	sort.Slice(expectingType, func(i, j int) bool {
		return expectingType[i].NameExpr().Text() < expectingType[j].NameExpr().Text()
	})
	sort.Slice(actualType, func(i, j int) bool {
		return actualType[i].NameExpr().Text() < actualType[j].NameExpr().Text()
	})

	for index, each := range expectingType {
		ac := actualType[index]
		if !each.Equal(ac) {
			return false
		}
	}

	if len(a.Service) != len(api.Service) {
		return false
	}

	var expectingService, actualService []*Service
	expectingService = append(expectingService, a.Service...)
	actualService = append(actualService, api.Service...)
	for index, each := range expectingService {
		ac := actualService[index]
		if !each.Equal(ac) {
			return false
		}
	}

	return true
}
