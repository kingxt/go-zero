package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

func TestInfo(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.InfoBlock().Accept(visitor)
	}

	test(t, do, spec.Info{
		Proterties: map[string]string{
			"title": `foo`,
			"desc":  `bar`,
		},
	}, false, `info(
		title: "foo"
		desc: "bar"
	)`)

	test(t, do, spec.Info{
		Proterties: map[string]string{},
	}, false, `info()`)

	test(t, do, spec.Info{
		Proterties: map[string]string{
			"title": "",
			"desc":  "",
		}}, false, `info(
		title:
		desc:
	)`)

	test(t, do, spec.Info{
		Proterties: map[string]string{
			"title": "foo",
			"desc":  "foo\n\t\tbar",
		},
	}, false, `info(
		title: "foo"
		desc: "foo
		bar"		
	)`)

	test(t, do, spec.Info{
		Proterties: map[string]string{
			"title":   "this is title",
			"desc":    "new line\ndescription",
			"Chinese": "你好",
		},
	}, false, `info (
	title: this is title
	desc: "new line
description"
	Chinese: 你好
)`)

	test(t, do, nil, true, `info(
	title: ""
	title: ""
	)`)

	test(t, do, nil, true, `info`)
	test(t, do, nil, true, `info (`)
}

func TestInfoToken(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.InfoBlock().Accept(visitor)
	}
	test(t, do, nil, true, `inf ()`)
	test(t, do, nil, true, `import ()`)
}

func TestKvLit(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.InfoBlock().Accept(visitor)
	}
	test(t, do, spec.Info{
		Proterties: map[string]string{
			"foo":  "bar",
			"bar":  "foo",
			"fooo": "foo bar",
		},
	}, false, `info (
foo:bar
bar:foo
fooo:foo bar
)
`)
}
