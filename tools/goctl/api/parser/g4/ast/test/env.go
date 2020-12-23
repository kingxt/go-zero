package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
)

var logEnable = true

func test(t *testing.T, do func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{}, expected interface{}, expectedErr bool, content string) {
	p := ast.NewParser()
	result, err := p.Accept(content, do)
	if logEnable {
		if err != nil {
			logx.Error(err)
		}
	}
	if expectedErr {
		assert.Error(t, err)
		return
	}

	assert.Nil(t, err)

	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(result)
	assert.JSONEq(t, string(expectedJson), string(actualJson))
}
