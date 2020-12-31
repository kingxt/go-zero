package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

var fieldAccept = func(p *api.ApiParserParser, visitor *ast.ApiVisitor) interface{} {
	return p.Field().Accept(visitor)
}

func TestField(t *testing.T) {
	t.Run("anonymous", func(t *testing.T) {
		v, err := parser.Accept(fieldAccept, `User`)
		assert.Nil(t, err)
		f := v.(*ast.TypeField)
		assert.True(t, f.Equal(&ast.TypeField{
			IsAnonymous: true,
			DataType:    &ast.Literal{Literal: ast.NewTextExpr("User")},
		}))

		v, err = parser.Accept(fieldAccept, `*User`)
		assert.Nil(t, err)
		f = v.(*ast.TypeField)
		assert.True(t, f.Equal(&ast.TypeField{
			IsAnonymous: true,
			DataType: &ast.Pointer{
				PointerExpr: ast.NewTextExpr("*User"),
				Star:        ast.NewTextExpr("*"),
				Name:        ast.NewTextExpr("User"),
			},
		}))

		v, err = parser.Accept(fieldAccept, `
		// anonymous user
		*User // pointer type`)
		assert.Nil(t, err)
		f = v.(*ast.TypeField)
		assert.True(t, f.Equal(&ast.TypeField{
			IsAnonymous: true,
			DataType: &ast.Pointer{
				PointerExpr: ast.NewTextExpr("*User"),
				Star:        ast.NewTextExpr("*"),
				Name:        ast.NewTextExpr("User"),
			},
			DocExpr:     ast.NewTextExpr("// anonymous user"),
			CommentExpr: ast.NewTextExpr("// pointer type"),
		}))

		_, err = parser.Accept(fieldAccept, `interface`)
		assert.Error(t, err)

		_, err = parser.Accept(fieldAccept, `map`)
		assert.Error(t, err)

		_, err = parser.Accept(fieldAccept, `*var`)
		assert.Error(t, err)
	})

	t.Run("normal", func(t *testing.T) {
		v, err := parser.Accept(fieldAccept, `User int`)
		assert.Nil(t, err)
		f := v.(*ast.TypeField)
		assert.True(t, f.Equal(&ast.TypeField{
			Name:     ast.NewTextExpr("User"),
			DataType: &ast.Literal{Literal: ast.NewTextExpr("int")},
		}))
		v, err = parser.Accept(fieldAccept, `Foo Bar`)
		assert.Nil(t, err)
		f = v.(*ast.TypeField)
		assert.True(t, f.Equal(&ast.TypeField{
			Name:     ast.NewTextExpr("Foo"),
			DataType: &ast.Literal{Literal: ast.NewTextExpr("Bar")},
		}))
	})
}
