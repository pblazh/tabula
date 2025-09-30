package core

import (
	"github.com/pblazh/tabula/internal/ast"
)

type (
	InternalFunc = func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, args ...ast.Expression) (ast.Expression, error)
	dispatchMap = map[string]InternalFunc
)

var DispatchMap dispatchMap = dispatchMap{
	"EXEC": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "EXEC(command:string, args:string...):string"
		return Exec(format, call, values...)
	},

	// numberic functions
	"SUM": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "SUM(values:number...):number"
		guard := MakeSameTypeGuard(format, ast.IsNumeric)
		return callNumbersFunction(format, sum, sum, guard, call, values...)
	},

	"ADD": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ADD(a:number, b:number):number"
		guard := MakeArityGuard(format, 2)
		return callNumbersFunction(format, sum, sum, guard, call, values...)
	},

	"PRODUCT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return callNumbersFunction("PRODUCT(values:number...):number", product, product, EmptyGuard, call, values...)
	},

	"AVERAGE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return callNumbersFunction("AVERAGE(values:number...):number", average, average, EmptyGuard, call, values...)
	},

	"MAX": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return callNumbersFunction("MAX(values:number...):number", max, max, EmptyGuard, call, values...)
	},

	"MAXA": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		converted, failed := parseStringExpressions(call, values)
		if failed != nil {
			return nil, ErrUnsupportedArgument("MAXA(values:number|string...):number",
				call, failed)
		}
		return callNumbersFunction("MAXA(values:number|string...):number", max, max, EmptyGuard, call, converted...)
	},

	"MIN": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return callNumbersFunction("MIN(values:number...):number", min, min, EmptyGuard, call, values...)
	},

	"MINA": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		converted, failed := parseStringExpressions(call, values)
		if failed != nil {
			return nil, ErrUnsupportedArgument("MINA(values:number|string...):number",
				call, failed)
		}
		return callNumbersFunction("MINA(values:number|string...):number", min, min, EmptyGuard, call, converted...)
	},

	"ABS": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ABS(value:number):number"
		return callNumbersFunction(format, abs, abs, MakeArityGuard(format, 1), call, values...)
	},

	"CEILING": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "CEILING(value:number, significance:[number]):number"
		return RoundUp(format, call, values...)
	},

	"FLOOR": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "FLOOR(value:number, significance:[number]):number"
		return RoundDown(format, call, values...)
	},

	"ROUND": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ROUND(value:number, places:[number]):number"
		return Round(format, call, values...)
	},

	"POWER": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "POWER(base:number, exponent:number):number"
		return Power(format, call, values...)
	},

	"INT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return Int(call, values...)
	},

	"MOD": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "MOD(dividend:number, divisor:number):number"
		guard := MakeExactTypesGuard(format, ast.IsNumeric, ast.IsNumeric)
		return callNumbersFunction(format, mod, mod, guard, call, values...)
	},

	"SQRT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "SQRT(value:number):number"
		return Sqrt(format, call, values...)
	},

	// string functions
	"CONCATENATE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "CONCATENATE(values:string...):string"
		return Concat(format, call, values...)
	},

	"LEN": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "LEN(value:string):number"
		return Len(format, call, values...)
	},

	"LOWER": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "LOWER(value:string):string"
		return Lower(format, call, values...)
	},

	"UPPER": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "UPPER(value:string):string"
		return Upper(format, call, values...)
	},

	"TRIM": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "TRIM(value:string):string"
		return Trim(format, call, values...)
	},

	"EXACT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "EXACT(a:string, b:string):boolean"
		return Exact(format, call, values...)
	},

	"FIND": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "FIND(what:string, where:string, [start:int]):number"
		return Find(format, call, values...)
	},

	"LEFT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "LEFT(value:string, [amount:int]):string"
		return Left(format, call, values...)
	},

	"RIGHT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "RIGHT(value:string, [amount:int]):string"
		return Right(format, call, values...)
	},

	"MID": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "MID(value:string, start:int, amount:int):string"
		return Mid(format, call, values...)
	},

	"SUBSTITUTE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "SUBSTITUTE(text:string, old:string, new:string, [instances:int]):string"
		return Substitute(format, call, values...)
	},

	"VALUE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "VALUE(value:string):number"
		return Value(format, call, values...)
	},

	// bulean functions
	"IF": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "IF(predicate:boolean, positive:any, negative:any):any"
		return If(format, call, values...)
	},

	"NOT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "NOT(value:boolean):boolean"
		return Not(format, call, values...)
	},

	"AND": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "AND(a:boolean, b:boolean):boolean"
		return And(format, call, values...)
	},

	"OR": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "OR(a:boolean, b:boolean):boolean"
		return Or(format, call, values...)
	},

	"TRUE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "TRUE():boolean"
		return True(format, call, values...)
	},

	"FALSE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "FALSE():boolean"
		return False(format, call, values...)
	},

	// Dates
	"TODATE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "TODATE(layout:string, value:string):date"
		return ToDate(format, call, values...)
	},

	"FROMDATE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "FROMDATE(layout:string, source:date):string"
		return FromDate(format, call, values...)
	},

	"DAY": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "DAY(value:date):number"
		return Day(format, call, values...)
	},

	"HOUR": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "HOUR(value:date):number"
		return Hour(format, call, values...)
	},

	"MINUTE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "MINUTE(value:date):number"
		return Minute(format, call, values...)
	},

	"MONTH": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "MONTH(value:date):number"
		return Month(format, call, values...)
	},

	"SECOND": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "SECOND(value:date):number"
		return Second(format, call, values...)
	},

	"YEAR": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "YEAR(value:date):number"
		return Year(format, call, values...)
	},

	"WEEKDAY": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "WEEKDAY(value:date):number"
		return Weekday(format, call, values...)
	},

	"NOW": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "NOW():date"
		return Now(format, call, values...)
	},

	"DATE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "DATE(year:number, month:number, day:number):date"
		return Date(format, call, values...)
	},

	"DATEDIF": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "DATEDIF(start:date, end:date, unit:string):number"
		return DateDiff(format, call, values...)
	},

	"DAYS": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "DAYS(start:date, end:date):number"
		return Days(format, call, values...)
	},

	"DATEVALUE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "DATEVALUE(value:string):date"
		return DateValue(format, call, values...)
	},

	"COUNT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return count(call, values...), nil
	},

	"COUNTA": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		return counta(call, values...), nil
	},

	// Information functions
	"ISNUMBER": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ISNUMBER(value:any):boolean"
		return IsNumber(format, call, values...)
	},

	"ISTEXT": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ISTEXT(value:any):boolean"
		return IsText(format, call, values...)
	},

	"ISLOGICAL": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ISLOGICAL(value:any):boolean"
		return IsBoolean(format, call, values...)
	},

	"ISBLANK": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ISBLANK(value:any):boolean"
		return IsBlank(format, call, values...)
	},

	// Lookup functions
	"ADDRESS": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ADDRESS(row:int, column:int):string"
		return Address(format, call, values...)
	},

	"ROW": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "ROW(cell:string):int"
		return Row(format, call, values...)
	},

	"COLUMN": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "COLUMN(cell:string):int"
		return Column(format, call, values...)
	},

	"REF": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "REF(cell:string):any"
		return Ref(context, input, formats, format, call, values...)
	},

	"RANGE": func(context map[string]string, input [][]string, formats map[string]string,
		call ast.CallExpression, values ...ast.Expression,
	) (ast.Expression, error) {
		format := "RANGE(a:string, b:string):range"
		return Range(context, input, formats, format, call, values...)
	},
}
