package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
	"github.com/tal-tech/go-zero/tools/goctl/api/spec"
)

const importLit = `import "foo.api"`

const importGroup = `
import (
	"foo.api"
	"bar.api"
	"foo/bar.api"
)
`

func TestImport(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ImportSpec().Accept(visitor)
	}

	test(t, do, spec.ApiImport{
		List: []spec.Import{
			{
				Value: "foo.api",
			},
		},
	}, false, importLit)

	test(t, do, spec.ApiImport{
		List: []spec.Import{
			{
				Value: "foo.api",
			},
			{
				Value: "bar.api",
			},
			{
				Value: "foo/bar.api",
			},
		},
	}, false, importGroup)

	test(t, do, spec.ApiImport{}, false, `import ()`)
	test(t, do, nil, true, `import`)
	test(t, do, nil, true, `import user.api`)
	test(t, do, nil, true, `import "user.api "`)
	test(t, do, nil, true, `import "/"`)
	test(t, do, nil, true, `import " "`)
	test(t, do, nil, true, `import "user-.api"`)
	test(t, do, nil, true, ``)
	test(t, do, nil, true, `
	import (
		"user"
	)`,
	)

	test(t, do, nil, true, `
	import (
		"user.ap"
	)`,
	)

	test(t, do, nil, true, `
	import (
		"user\user.api"
	)`,
	)

	test(t, do, nil, true, `
	import (
		"user.api"
	`,
	)

	test(t, do, nil, true, `import (
		"user.api"
		"user.api"
	)`)
}

func TestImportToken(t *testing.T) {
	do := func(p *api.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ImportSpec().Accept(visitor)
	}
	test(t, do, false, true, `imports "user.api"`)
	test(t, do, false, true, `qw "user.api"`)
	test(t, do, false, true, `"user.api"`)
}
