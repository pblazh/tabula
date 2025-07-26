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
		return Round(true, call, values...)
	},
	"FLOOR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Round(false, call, values...)
	},
	"POWER": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Power(call, values...)
	},
	"INT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Int(call, values...)
	},
}
