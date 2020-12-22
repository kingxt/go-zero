package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
)

var logEnable = false

func test(t *testing.T, do func(p *ast.Parser, visitor *ast.ApiVisitor) interface{}, expected interface{}, expectedErr bool, content string) {
	defer func() {
		p := recover()
		if expectedErr {
			if logEnable {
				fmt.Printf("%+v\n", p)
			}
			assert.NotNil(t, p)
			return
		}
		assert.Nil(t, p)
	}()

	p := ast.NewParser(content)
	visitor := ast.NewApiVisitor()
	result := do(p, visitor)
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(result)
	assert.JSONEq(t, string(expectedJson), string(actualJson))
}
