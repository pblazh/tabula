package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

type (
	InternalFunc = func(call ast.CallExpression, args ...ast.Expression) (ast.Expression, error)
	dispatchMap  = map[string]InternalFunc
)

var DispatchMap dispatchMap = dispatchMap{
	// numberic functions
	"SUM": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "SUM(number...)"
		guard := MakeSameTypeGuard(format, ast.IsNumeric)
		return callNumbersFunction(format, sum, sum, guard, call, values...)
	},
	"ADD": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "ADD(number, number)"
		guard := MakeArityGuard(format, 2)
		return callNumbersFunction(format, sum, sum, guard, call, values...)
	},
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
	// string functions
	"CONCATENATE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "CONCATENATE(string...)"
		return Concat(format, call, values...)
	},
	"LEN": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "LEN(string)"
		return Len(format, call, values...)
	},
	"LOWER": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "LOWER(string)"
		return Lower(format, call, values...)
	},
	"UPPER": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "UPPER(string)"
		return Upper(format, call, values...)
	},
	"TRIM": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "TRIM(string)"
		return Trim(format, call, values...)
	},
	"EXACT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "EXACT(string, string)"
		return Exact(format, call, values...)
	},
	// bulean functions
	"FALSE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "FALSE(boolean)"
		return False(format, call, values...)
	},
}
