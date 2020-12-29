package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"
)

type (
	ApiVisitor struct {
		api.BaseApiParserVisitor
		debug    bool
		prefix   string
		infoFlag bool
	}

	VisitorOption func(v *ApiVisitor)

	Spec interface {
		Doc() Expr
		Comment() Expr
		Format() error
		Equal(v interface{}) bool
	}

	Expr interface {
		Prefix() string
		Line() int
		Column() int
		Text() string
		Start() int
		Stop() int
	}
)

func NewApiVisitor(options ...VisitorOption) *ApiVisitor {
	v := &ApiVisitor{}
	for _, opt := range options {
		opt(v)
	}
	return v
}

func (v *ApiVisitor) fmtErrorf(expr Expr, msg string) {
	panic(fmt.Sprintf("%s line %d:%d  %s", v.prefix, expr.Line(), expr.Column(), msg))
}

func WithVisitorPrefix(prefix string) VisitorOption {
	return func(v *ApiVisitor) {
		v.prefix = prefix
	}
}

var VisitorDebug = WithVisitorDebug()

func WithVisitorDebug() VisitorOption {
	return func(v *ApiVisitor) {
		v.debug = true
	}
}

type defaultExpr struct {
	prefix, v    string
	line, column int
	start, stop  int
}

func NewExpr(prefix, v string, line, column, start, stop int) *defaultExpr {
	return &defaultExpr{
		prefix: prefix,
		line:   line,
		column: column,
		v:      v,
		start:  start,
		stop:   stop,
	}
}

func NewTextExpr(v string) *defaultExpr {
	return &defaultExpr{
		v: v,
	}
}

func (v *ApiVisitor) newExprWithTerminalNode(node antlr.TerminalNode) *defaultExpr {
	instance := &defaultExpr{}
	if node != nil {
		token := node.GetSymbol()
		return v.newExprWithToken(token)
	}

	return instance
}

func (v *ApiVisitor) newExprWithToken(token antlr.Token) *defaultExpr {
	instance := &defaultExpr{}
	if token != nil {
		instance.prefix = v.prefix
		instance.v = token.GetText()
		instance.line = token.GetLine()
		instance.column = token.GetColumn()
		instance.start = token.GetStart()
		instance.stop = token.GetStop()
	}

	return instance
}

func (e *defaultExpr) Prefix() string {
	return e.prefix
}

func (e *defaultExpr) Line() int {
	return e.line
}

func (e *defaultExpr) Column() int {
	return e.column
}

func (e *defaultExpr) Text() string {
	return e.v
}

func (e *defaultExpr) Start() int {
	return e.start
}

func (e *defaultExpr) Stop() int {
	return e.stop
}

func ExprEqual(expr1, expr2 Expr) bool {
	if expr1 == nil {
		if expr2 != nil {
			return false
		}
		return true
	}

	if expr2 == nil {
		return false
	}

	return expr1.Text() == expr2.Text()
}

func EqualDoc(spec1, spec2 Spec) bool {
	if !ExprEqual(spec1.Doc(), spec2.Doc()) {
		return false
	}

	if !ExprEqual(spec1.Comment(), spec2.Comment()) {
		return false
	}

	return true
}
