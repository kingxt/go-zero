package parser

import (
	"fmt"
	"path/filepath"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
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
	p.convert2Spec()
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
	p.convert2Spec()
	return spec, nil
}

// todo
func (p parser) convert2Spec() {
	p.fillInfo()
	p.fillSyntax()
	p.fillImport()
	p.fillTypes()
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

func (p parser) fillTypes() {
	for _, item := range p.ast.Type {
		switch v := (item).(type) {
		case *ast.TypeStruct:
			var members []spec.Member
			for _, item := range v.Fields {
				members = append(members, p.fieldToMember(item))
			}
			p.spec.Types = append(p.spec.Types, spec.TypeStruct{
				RawName: v.Name.Text(),
				Members: members,
				Docs:    p.stringExprs(v.Doc()),
			})
		default:
			panic(fmt.Sprintf("unknown type %+v", v))
		}
	}
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
		return spec.BasicType{RawName: raw}
	}
}

func (p parser) stringExprs(docs []ast.Expr) []string {
	var result []string
	for _, item := range docs {
		result = append(result, item.Text())
	}
	return result
}
