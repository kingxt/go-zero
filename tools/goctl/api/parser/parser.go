package parser

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

type parser struct {
	ast  *ast.Api
	spec *spec.ApiSpec
}

func Parse(filename string) (*spec.ApiSpec, error) {
	astParser := ast.NewParser(ast.WithParserPrefix(filepath.Base(filename)))
	ast, err := astParser.Parse(filename)
	if err != nil {
		return nil, err
	}

	spec := new(spec.ApiSpec)
	p := parser{ast: ast, spec: spec}
	err = p.convert2Spec()
	if err != nil {
		return nil, err
	}

	return spec, nil
}

func ParseContent(content string) (*spec.ApiSpec, error) {
	astParser := ast.NewParser()
	ast, err := astParser.ParseContent(content)
	if err != nil {
		return nil, err
	}

	spec := new(spec.ApiSpec)
	p := parser{ast: ast, spec: spec}
	err = p.convert2Spec()
	if err != nil {
		return nil, err
	}

	return spec, nil
}

// todo
func (p parser) convert2Spec() error {
	p.fillInfo()
	p.fillSyntax()
	p.fillImport()
	return p.fillTypes()
}

func (p parser) fillInfo() {
	if p.ast.Info != nil {
		p.spec.Info = spec.Info{}
		for _, kv := range p.ast.Info.Kvs {
			p.spec.Info.Proterties[kv.Key.Text()] = kv.Value.Text()
		}
	}
}

func (p parser) fillSyntax() {
	if p.ast.Syntax != nil {
		p.spec.Syntax = spec.ApiSyntax{Version: p.ast.Syntax.Version.Text()}
	}
}

func (p parser) fillImport() {
	if len(p.ast.Import) > 0 {
		for _, item := range p.ast.Import {
			p.spec.Imports = append(p.spec.Imports, spec.Import{Value: item.Value.Text()})
		}
	}
}

func (p parser) fillTypes() error {
	for _, item := range p.ast.Type {
		switch v := (item).(type) {
		case *ast.TypeStruct:
			var members []spec.Member
			for _, item := range v.Fields {
				members = append(members, p.fieldToMember(item))
			}
			p.spec.Types = append(p.spec.Types, spec.DefineStruct{
				RawName: v.Name.Text(),
				Members: members,
				Docs:    p.stringExprs(v.Doc()),
			})
		default:
			return errors.New(fmt.Sprintf("unknown type %+v", v))
		}
	}

	for _, item := range p.spec.Types {
		switch v := (item).(type) {
		case spec.DefineStruct:
			for _, member := range v.Members {
				switch v := member.Type.(type) {
				case spec.DefineStruct:
					tp, err := p.findDefinedType(v.RawName)
					if err != nil {
						return err
					} else {
						member.Type = *tp
					}
				}
			}
		default:
			return errors.New(fmt.Sprintf("unknown type %+v", v))
		}
	}
	return nil
}

func (p parser) findDefinedType(name string) (*spec.Type, error) {
	for _, item := range p.spec.Types {
		if _, ok := item.(spec.DefineStruct); ok {
			if item.Name() == name {
				return &item, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("type %s not defined", name))
}

func (p parser) fieldToMember(field *ast.TypeField) spec.Member {
	return spec.Member{
		Name:     field.Name.Text(),
		Type:     p.astTypeToSpec(field.DataType),
		Tag:      field.Tag.Text(),
		Comment:  field.Comment().Text(),
		Docs:     p.stringExprs(field.Doc()),
		IsInline: field.IsAnonymous,
	}
}

func (p parser) astTypeToSpec(in ast.DataType) spec.Type {
	switch v := (in).(type) {
	case *ast.Literal:
		raw := v.Literal.Text()
		if api.IsBasicType(raw) {
			return spec.BasicType{RawName: raw}
		} else {
			return spec.DefineStruct{RawName: raw}
		}
	case *ast.Interface:
		return spec.InterfaceType{RawName: v.Literal.Text()}
	case *ast.Map:
		return spec.MapType{RawName: v.MapExpr.Text(), Key: v.Key.Text(), Value: p.astTypeToSpec(v.Value)}
	case *ast.Array:
		return spec.ArrayType{RawName: v.ArrayExpr.Text(), Value: p.astTypeToSpec(v.Literal)}
	case *ast.Pointer:
		raw := v.Name.Text()
		if api.IsBasicType(raw) {
			return spec.PointerType{RawName: v.PointerExpr.Text(), Type: spec.BasicType{RawName: raw}}
		} else {
			return spec.PointerType{RawName: v.PointerExpr.Text(), Type: spec.DefineStruct{RawName: raw}}
		}
	}
	panic(fmt.Sprintf("unspported type %+v", in))
}

func (p parser) stringExprs(docs []ast.Expr) []string {
	var result []string
	for _, item := range docs {
		result = append(result, item.Text())
	}
	return result
}

func (p parser) fillService() error {
	for _, item := range p.ast.Service {
		var group spec.Group
		if item.AtServer != nil {
			for _, kv := range item.AtServer.Kv {
				group.Annotation.Properties[kv.Key.Text()] = kv.Value.Text()
			}
		}

		for _, route := range item.ServiceApi.ServiceRoute {
			r := spec.Route{
				Annotation: spec.Annotation{},
				Method:     route.Route.Method.Text(),
				Path:       route.Route.Path.Text(),
			}
			if route.AtHandler != nil {
				r.Handler = route.AtHandler.Name.Text()
			}
			if route.Route.Req != nil {
				tp, err := p.findDefinedType(route.Route.Req.Name.Text())
				if err != nil {
					return err
				}

				r.RequestType = *tp
			}
			if route.Route.Reply != nil {
				tp, err := p.findDefinedType(route.Route.Reply.Name.Text())
				if err != nil {
					return err
				}

				r.ResponseType = *tp
			}
			group.Routes = append(group.Routes, r)
			p.spec.Service.Groups = append(p.spec.Service.Groups, group)

			name := item.ServiceApi.Name.Text()
			if len(p.spec.Service.Name) > 0 && p.spec.Service.Name != name {
				return errors.New(fmt.Sprintf("mulit service name defined %s and %s", name, p.spec.Service.Name))
			}
			p.spec.Service.Name = name
		}
	}
	return nil
}
