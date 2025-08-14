package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

type (
	InternalFunc = func(call ast.CallExpression, args ...ast.Expression) (ast.Expression, error)
	dispatchMap  = map[string]InternalFunc
)

var DispatchMap dispatchMap = dispatchMap{
	"CALL": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "CALL(command, string...)"
		return Call(format, call, values...)
	},
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
	"MAX": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callNumbersFunction("MAX(number...)", max, max, EmptyGuard, call, values...)
	},
	"MAXA": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		converted, failed := parseStringExpressions(call, values)
		if failed != nil {
			return nil, ErrUnsupportedArgument("MAXA(number...)", call, failed)
		}
		return callNumbersFunction("MAXA(number|string...)", max, max, EmptyGuard, call, converted...)
	},
	"MIN": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return callNumbersFunction("MIN(number...)", min, min, EmptyGuard, call, values...)
	},
	"MINA": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		converted, failed := parseStringExpressions(call, values)
		if failed != nil {
			return nil, ErrUnsupportedArgument("MINA(number...)", call, failed)
		}
		return callNumbersFunction("MINA(number|string...)", min, min, EmptyGuard, call, converted...)
	},
	"ABS": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "ABS(number)"
		return callNumbersFunction(format, abs, abs, MakeArityGuard(format, 1), call, values...)
	},
	"CEILING": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "CEILING(number, [number])"
		return RoundUp(format, call, values...)
	},
	"FLOOR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "FLOOR(number, [number])"
		return RoundDown(format, call, values...)
	},
	"ROUND": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "ROUND(number, [number])"
		return Round(format, call, values...)
	},
	"POWER": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Power(call, values...)
	},
	"INT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return Int(call, values...)
	},
	"MOD": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "MOD(number, number)"
		guard := MakeExactTypesGuard(format, ast.IsNumeric, ast.IsNumeric)
		return callNumbersFunction(format, mod, mod, guard, call, values...)
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
	"IF": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "IF(boolean, any, any)"
		return If(format, call, values...)
	},
	"NOT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "NOT(boolean)"
		return Not(format, call, values...)
	},
	"AND": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "AND(boolean, boolean)"
		return And(format, call, values...)
	},
	"OR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "OR(boolean, boolean)"
		return Or(format, call, values...)
	},
	"TRUE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "TRUE()"
		return True(format, call, values...)
	},
	"FALSE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "FALSE()"
		return False(format, call, values...)
	},

	// Dates
	"TODATE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "TODATE(string, string)"
		return ToDate(format, call, values...)
	},
	"FROMDATE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "FROMDATE(string, date)"
		return FromDate(format, call, values...)
	},
	"DAY": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "DAY(date)"
		return Day(format, call, values...)
	},
	"HOUR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "HOUR(date)"
		return Hour(format, call, values...)
	},
	"MINUTE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "MINUTE(date)"
		return Minute(format, call, values...)
	},
	"MONTH": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "MONTH(date)"
		return Month(format, call, values...)
	},
	"SECOND": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "SECOND(date)"
		return Second(format, call, values...)
	},
	"YEAR": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "YEAR(date)"
		return Year(format, call, values...)
	},
	"WEEKDAY": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "WEEKDAY(date)"
		return Weekday(format, call, values...)
	},
	"NOW": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "NOW()"
		return Now(format, call, values...)
	},
	"DATE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "DATE(year, month, day)"
		return Date(format, call, values...)
	},
	"DATEDIF": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "DATEDIF(from, to, unit)"
		return DateDiff(format, call, values...)
	},
	"DAYS": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "DAYS(from, to)"
		return Days(format, call, values...)
	},
	"DATEVALUE": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		format := "DATEVALUE(string)"
		return DateValue(format, call, values...)
	},
	"COUNT": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return count(call, values...), nil
	},
	"COUNTA": func(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
		return counta(call, values...), nil
	},
}
