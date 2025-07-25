package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

type (
	InternalFunc = func(call ast.CallExpression, args ...ast.Expression) (ast.Expression, error)
	dispatchMap  = map[string]InternalFunc
)

var DispatchMap dispatchMap = dispatchMap{
	"SUM": Sum,
	"PRODUCT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callMathFunction(product, product, EmptyGuard, call, values...)
	},
	"AVERAGE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callMathFunction(average, average, EmptyGuard, call, values...)
	},
	"ABS": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callMathFunction(abs, abs, MakeArityGuard(1), call, values...)
	},
	"CEILING": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callMathFunction(ceil, ceil, MakeArityGuard(1), call, values...)
	},
}
