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
		return callNumbersFunction("PRODUCT(number...)", product, product, EmptyGuard, call, values...)
	},
	"AVERAGE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callNumbersFunction("AVERAGE(number...)", average, average, EmptyGuard, call, values...)
	},
	"ABS": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "ABS(number)"
		return callNumbersFunction(format, abs, abs, MakeArityGuard(format, 1), call, values...)
	},
	"CEILING": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "CEILING(number, number)"
		return Round(true, format, call, values...)
	},
	"FLOOR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "FLOOR(number, number)"
		return Round(false, format, call, values...)
	},
	"POWER": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Power(call, values...)
	},
	"INT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Int(call, values...)
	},
}
