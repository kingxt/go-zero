package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
)

var logEnable = true

func test(t *testing.T, do func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{}, expected interface{}, expectedErr bool, content string) {
	p := ast.NewParser()
	result, err := p.Accept(content, do)
	if expectedErr {
		assert.Error(t, err)
		if logEnable {
			fmt.Printf("%+v\r\n", err)
		}
		return
	}

	assert.Nil(t, err)
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(result)
	assert.JSONEq(t, string(expectedJson), string(actualJson))
}
