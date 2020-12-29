package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

var importAccept = func(p *api.ApiParserParser, visitor *ast.ApiVisitor) interface{} {
	return p.ImportSpec().Accept(visitor)
}

func TestImport(t *testing.T) {
	t.Run("matched", func(t *testing.T) {
		v, err := parser.Accept(importAccept, `import "foo.api"`)
		assert.Nil(t, err)

		list := v.([]*ast.ImportExpr)
		for _, each := range list {
			assert.True(t, each.Equal(&ast.ImportExpr{
				Import: ast.NewTextExpr("import"),
				Value:  ast.NewTextExpr(`"foo.api"`),
			}))
		}
	})

	t.Run("matched doc", func(t *testing.T) {
		v, err := parser.Accept(importAccept, `
		/**doc*/
		import "foo.api" /**line doc*/`)
		assert.Nil(t, err)

		list := v.([]*ast.ImportExpr)
		for _, each := range list {
			assert.True(t, each.Equal(&ast.ImportExpr{
				Import:      ast.NewTextExpr("import"),
				Value:       ast.NewTextExpr(`"foo.api"`),
				DocExpr:     ast.NewTextExpr("/**doc*/"),
				CommentExpr: ast.NewTextExpr("/**line doc*/"),
			}))
		}
	})

	t.Run("matched comment", func(t *testing.T) {
		v, err := parser.Accept(importAccept, `
		// comment block
		import "foo.api" // line comment`)
		assert.Nil(t, err)

		list := v.([]*ast.ImportExpr)
		for _, each := range list {
			assert.True(t, each.Equal(&ast.ImportExpr{
				Import:      ast.NewTextExpr("import"),
				Value:       ast.NewTextExpr(`"foo.api"`),
				DocExpr:     ast.NewTextExpr("// comment block"),
				CommentExpr: ast.NewTextExpr("// line comment"),
			}))
		}
	})

	t.Run("mismatched import", func(t *testing.T) {
		_, err := parser.Accept(importAccept, `
		 "foo.api"`)
		assert.Error(t, err)

		_, err = parser.Accept(importAccept, `
		 impor "foo.api"`)
		assert.Error(t, err)
	})

	t.Run("mismatched value", func(t *testing.T) {
		_, err := parser.Accept(importAccept, `
		 import "foo"`)
		assert.Error(t, err)

		_, err = parser.Accept(importAccept, `
		 import ""`)
		assert.Error(t, err)

		_, err = parser.Accept(importAccept, `
		 import `)
		assert.Error(t, err)
	})
}
