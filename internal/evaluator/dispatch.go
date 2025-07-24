package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/functions"
)

type (
	InternalFunc = func(call ast.CallExpression, args ...ast.Expression) (ast.Expression, error)
	DispatchMap  = map[string]InternalFunc
)

var dispatchMap DispatchMap = DispatchMap{
	"SUM": functions.Sum,
}
