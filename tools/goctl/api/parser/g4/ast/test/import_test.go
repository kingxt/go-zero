package test

import (
	"testing"

	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/ast"
	parser "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen"
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
	do := func(p *parser.ApiParser, visitor *ast.ApiVisitor) interface{} {
		return p.ImportSpec().Accept(visitor)
	}

	test(t, do, spec.ApiImport{
		List: []string{"foo.api"},
	}, false, importLit)

	test(t, do, spec.ApiImport{
		List: []string{"foo.api", "bar.api", "foo/bar.api"},
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
