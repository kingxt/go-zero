package ast

import "github.com/tal-tech/go-zero/tools/goctl/api/parser/g4/g4gen/api"

type KvExpr struct {
	Key         Expr
	Value       Expr
	DocExpr     []Expr
	CommentExpr Expr
}

func (v *ApiVisitor) VisitKvLit(ctx *api.KvLitContext) interface{} {
	key := v.newExprWithToken(ctx.GetKey())
	value := v.newExprWithToken(ctx.GetValue())
	return &KvExpr{
		Key:         key,
		Value:       value,
		DocExpr:     v.getDoc(ctx),
		CommentExpr: v.getComment(ctx),
	}
}

func (k *KvExpr) Format() error {
	// todo
	return nil
}

func (k *KvExpr) Equal(v interface{}) bool {
	if v == nil {
		return false
	}

	kv, ok := v.(*KvExpr)
	if !ok {
		return false
	}

	if !EqualDoc(k, kv) {
		return false
	}

	return k.Key.Equal(kv.Key) && k.Value.Equal(kv.Value)
}

func (k *KvExpr) Doc() []Expr {
	return k.DocExpr
}

func (k *KvExpr) Comment() Expr {
	return k.CommentExpr
}
