package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const unDeclaretionStruct = `
type Foo {
	Bar Bar
}
`

func TestParseContent(t *testing.T) {
	parser := NewParser()
	_, err := parser.ParseContent(unDeclaretionStruct)
	assert.Error(t, err)
}
