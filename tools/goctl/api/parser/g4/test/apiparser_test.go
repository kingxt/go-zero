package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
)

var (
	normalApi = `
	syntax="v1"
	
	info (
		foo: bar
	)

	type Foo {
		Bar int
	}

	@server(
		foo: bar
	)
	service foo-api{
		@doc("foo")
		@handler foo
		post /foo (Foo) returns (Foo)
	}
`
	missDeclarationApi = `
	@server(
		foo: bar
	)
	service foo-api{
		@doc("foo")
		@handler foo
		post /foo (Foo) returns (Foo)
	}
	`

	nestedApiImport = `
		import "foo.api"
	`
)

func TestApiParser(t *testing.T) {
	t.Run("missDeclarationApi", func(t *testing.T) {
		_, err := parser.ParseContent(missDeclarationApi)
		assert.Error(t, err)
		fmt.Printf("%+v\n", err)
	})

	t.Run("missDeclarationApi", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "foo.api")
		err := ioutil.WriteFile(file, []byte(nestedApiImport), os.ModePerm)
		if err != nil {
			return
		}
		_, err = parser.ParseContent(fmt.Sprintf(`import "%s"`, file))
		assert.Error(t, err)
	})

	t.Run("normal", func(t *testing.T) {
		v, err := parser.ParseContent(normalApi)
		assert.Nil(t, err)
		body := &ast.Body{
			Lp:   ast.NewTextExpr("("),
			Rp:   ast.NewTextExpr(")"),
			Name: ast.NewTextExpr("Foo"),
		}
		assert.True(t, v.Equal(&ast.Api{
			Syntax: &ast.SyntaxExpr{
				Syntax:  ast.NewTextExpr("syntax"),
				Assign:  ast.NewTextExpr("="),
				Version: ast.NewTextExpr(`"v1"`),
			},
			Info: &ast.InfoExpr{
				Info: ast.NewTextExpr("info"),
				Lp:   ast.NewTextExpr("("),
				Rp:   ast.NewTextExpr(")"),
				Kvs: []*ast.KvExpr{
					{
						Key:   ast.NewTextExpr("foo"),
						Value: ast.NewTextExpr("bar"),
					},
				},
			},
			Type: []ast.TypeExpr{
				&ast.TypeStruct{
					Name:   ast.NewTextExpr("Foo"),
					LBrace: ast.NewTextExpr("{"),
					RBrace: ast.NewTextExpr("}"),
					Fields: []*ast.TypeField{
						{
							Name:     ast.NewTextExpr("Bar"),
							DataType: &ast.Literal{Literal: ast.NewTextExpr("int")},
						},
					},
				},
			},
			Service: []*ast.Service{
				{
					AtServer: &ast.AtServer{
						AtServerToken: ast.NewTextExpr("@server"),
						Lp:            ast.NewTextExpr("("),
						Rp:            ast.NewTextExpr(")"),
						Kv: []*ast.KvExpr{
							{
								Key:   ast.NewTextExpr("foo"),
								Value: ast.NewTextExpr("bar"),
							},
						},
					},
					ServiceApi: &ast.ServiceApi{
						ServiceToken: ast.NewTextExpr("service"),
						Name:         ast.NewTextExpr("foo-api"),
						Lbrace:       ast.NewTextExpr("{"),
						Rbrace:       ast.NewTextExpr("}"),
						ServiceRoute: []*ast.ServiceRoute{
							{
								AtDoc: &ast.AtDoc{
									AtDocToken: ast.NewTextExpr("@doc"),
									Lp:         ast.NewTextExpr("("),
									Rp:         ast.NewTextExpr(")"),
									LineDoc:    ast.NewTextExpr(`"foo"`),
								},
								AtHandler: &ast.AtHandler{
									AtHandlerToken: ast.NewTextExpr("@handler"),
									Name:           ast.NewTextExpr("foo"),
								},
								Route: &ast.Route{
									Method:      ast.NewTextExpr("post"),
									Path:        ast.NewTextExpr("/foo"),
									Req:         body,
									ReturnToken: ast.NewTextExpr("returns"),
									Reply:       body,
								},
							},
						},
					},
				},
			},
		}))
	})
}
