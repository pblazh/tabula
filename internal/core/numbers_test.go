package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

type inputCase struct {
	f       string
	expects string
	error   string
}

func TestMathFunctions(t *testing.T) {
	testcases := []struct {
		name  string
		input []ast.Expression
		cases []inputCase
	}{
		// Empty input
		{
			name:  "empty input",
			input: []ast.Expression{},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<int 1>",
				},
				{
					f:       "AVERAGE",
					expects: "<int 0>",
				},
				{
					f:       "SUM",
					expects: "<int 0>",
				},
				{
					f:       "COUNT",
					expects: "<int 0>",
				},
				{
					f:       "COUNTA",
					expects: "<int 0>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number expects 1 argument, got 0 in (ABS), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number expects 2 arguments, got 0 in (CEILING), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number expects 2 arguments, got 0 in (FLOOR), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number expects 2 arguments, got 0 in (ROUND), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expects 1 argument, got 0 in (INT), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 0 in (POWER), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 0 in (MOD), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number expects 1 argument, got 0 in (SQRT), at <: input:0:0>",
				},
			},
		},
		// Integer operations
		{
			name: "single integer",
			input: []ast.Expression{
				ast.IntExpression{Value: -5},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<int -5>",
				},
				{
					f:       "AVERAGE",
					expects: "<int -5>",
				},
				{
					f:       "MAX",
					expects: "<int -5>",
				},
				{
					f:       "MAXA",
					expects: "<int -5>",
				},
				{
					f:       "MIN",
					expects: "<int -5>",
				},
				{
					f:       "MINA",
					expects: "<int -5>",
				},
				{
					f:       "ABS",
					expects: "<int 5>",
				},
				{
					f:       "INT",
					expects: "<int -5>",
				},
				{
					f:       "SUM",
					expects: "<int -5>",
				},
				{
					f:       "COUNT",
					expects: "<int 1>",
				},
				{
					f:       "COUNTA",
					expects: "<int 1>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument <int -5> in (SQRT <int -5>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in (POWER <int -5>), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in (MOD <int -5>), at <: input:0:0>",
				},
			},
		},
		{
			name: "multiple integers",
			input: []ast.Expression{
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 4},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<int 24>",
				},
				{
					f:       "AVERAGE",
					expects: "<int 3>",
				},
				{
					f:       "MAX",
					expects: "<int 4>",
				},
				{
					f:       "MAXA",
					expects: "<int 4>",
				},
				{
					f:       "MIN",
					expects: "<int 2>",
				},
				{
					f:       "MINA",
					expects: "<int 2>",
				},
				{
					f:       "SUM",
					expects: "<int 9>",
				},
				{
					f:       "COUNT",
					expects: "<int 3>",
				},
				{
					f:       "COUNTA",
					expects: "<int 3>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number expects 1 argument, got 3 in (ABS <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number expects 2 arguments, got 3 in (CEILING <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number expects 2 arguments, got 3 in (FLOOR <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number expects 2 arguments, got 3 in (ROUND <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expects 1 argument, got 3 in (INT <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 3 in (POWER <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 3 in (MOD <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number expects 1 argument, got 3 in (SQRT <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
			},
		},
		{
			name: "integers with zero",
			input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 0},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<int 0>",
				},
				{
					f:       "AVERAGE",
					expects: "<int 5>",
				},
				{
					f:       "MAX",
					expects: "<int 10>",
				},
				{
					f:       "MAXA",
					expects: "<int 10>",
				},
				{
					f:       "MIN",
					expects: "<int 0>",
				},
				{
					f:       "MINA",
					expects: "<int 0>",
				},
				{
					f:       "SUM",
					expects: "<int 15>",
				},
			},
		},
		{
			name: "integers with negative values",
			input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: -3},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<int -60>",
				},
				{
					f:       "AVERAGE",
					expects: "<int 3>",
				},
				{
					f:       "MAX",
					expects: "<int 10>",
				},
				{
					f:       "MAXA",
					expects: "<int 10>",
				},
				{
					f:       "MIN",
					expects: "<int -3>",
				},
				{
					f:       "MINA",
					expects: "<int -3>",
				},
				{
					f:       "SUM",
					expects: "<int 9>",
				},
			},
		},
		// Float operations
		{
			name: "single float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float 5.50>",
				},
				{
					f:       "AVERAGE",
					expects: "<float 5.50>",
				},
				{
					f:       "MAX",
					expects: "<float 5.50>",
				},
				{
					f:       "MAXA",
					expects: "<float 5.50>",
				},
				{
					f:       "MIN",
					expects: "<float 5.50>",
				},
				{
					f:       "MINA",
					expects: "<float 5.50>",
				},
				{
					f:       "ABS",
					expects: "<float 5.50>",
				},
				{
					f:       "INT",
					expects: "<int 5>",
				},
				{
					f:       "SUM",
					expects: "<float 5.50>",
				},
				{
					f:       "CEILING",
					expects: "<int 6>",
				},
				{
					f:       "FLOOR",
					expects: "<int 5>",
				},
				{
					f:       "ROUND",
					expects: "<int 6>",
				},
				{
					f:       "SQRT",
					expects: "<float 2.35>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in (POWER <float 5.50>), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in (MOD <float 5.50>), at <: input:0:0>",
				},
			},
		},
		{
			name: "negative float",
			input: []ast.Expression{
				ast.FloatExpression{Value: -3.7},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float -3.70>",
				},
				{
					f:       "AVERAGE",
					expects: "<float -3.70>",
				},
				{
					f:       "MAX",
					expects: "<float -3.70>",
				},
				{
					f:       "MAXA",
					expects: "<float -3.70>",
				},
				{
					f:       "MIN",
					expects: "<float -3.70>",
				},
				{
					f:       "MINA",
					expects: "<float -3.70>",
				},
				{
					f:       "ABS",
					expects: "<float 3.70>",
				},
				{
					f:       "INT",
					expects: "<int -3>",
				},
				{
					f:       "SUM",
					expects: "<float -3.70>",
				},
				{
					f:       "CEILING",
					expects: "<int -3>",
				},
				{
					f:       "FLOOR",
					expects: "<int -4>",
				},
				{
					f:       "ROUND",
					expects: "<int -4>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument <float -3.70> in (SQRT <float -3.70>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in (POWER <float -3.70>), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in (MOD <float -3.70>), at <: input:0:0>",
				},
			},
		},
		{
			name: "multiple floats",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2.0},
				ast.FloatExpression{Value: 1.5},
				ast.FloatExpression{Value: 3.0},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float 9.00>",
				},
				{
					f:       "AVERAGE",
					expects: "<float 2.17>",
				},
				{
					f:       "MAX",
					expects: "<float 3.00>",
				},
				{
					f:       "MAXA",
					expects: "<float 3.00>",
				},
				{
					f:       "MIN",
					expects: "<float 1.50>",
				},
				{
					f:       "MINA",
					expects: "<float 1.50>",
				},
				{
					f:       "SUM",
					expects: "<float 6.50>",
				},
			},
		},
		{
			name: "floats with zero",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.FloatExpression{Value: 0.0},
				ast.FloatExpression{Value: 2.5},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float 0.00>",
				},
				{
					f:       "AVERAGE",
					expects: "<float 2.67>",
				},
				{
					f:       "MAX",
					expects: "<float 5.50>",
				},
				{
					f:       "MAXA",
					expects: "<float 5.50>",
				},
				{
					f:       "MIN",
					expects: "<float 0.00>",
				},
				{
					f:       "MINA",
					expects: "<float 0.00>",
				},
				{
					f:       "SUM",
					expects: "<float 8.00>",
				},
			},
		},
		// Mixed int and float operations (result type is float if any argument is float)
		{
			name: "int and float",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.FloatExpression{Value: 2.5},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float 12.50>",
				},
				{
					f:       "AVERAGE",
					expects: "<float 3.75>",
				},
				{
					f:       "MAX",
					expects: "<float 5.00>",
				},
				{
					f:       "MAXA",
					expects: "<float 5.00>",
				},
				{
					f:       "MIN",
					expects: "<float 2.50>",
				},
				{
					f:       "MINA",
					expects: "<float 2.50>",
				},
				{
					f:       "SUM",
					expects: "<float 7.50>",
				},
			},
		},
		{
			name: "float and int",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2.5},
				ast.IntExpression{Value: 4},
			},
			cases: []inputCase{
				{
					f:       "PRODUCT",
					expects: "<float 10.00>",
				},
				{
					f:       "AVERAGE",
					expects: "<float 3.25>",
				},
				{
					f:       "MAX",
					expects: "<float 4.00>",
				},
				{
					f:       "MAXA",
					expects: "<float 4.00>",
				},
				{
					f:       "MIN",
					expects: "<float 2.50>",
				},
				{
					f:       "MINA",
					expects: "<float 2.50>",
				},
				{
					f:       "POWER",
					expects: "<float 39.06>",
				},
				{
					f:       "SUM",
					expects: "<float 6.50>",
				},
			},
		},
		{
			name: "two ints",
			input: []ast.Expression{
				ast.IntExpression{Value: 21},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:       "CEILING",
					expects: "<int 22>",
				},
				{
					f:       "FLOOR",
					expects: "<int 20>",
				},
				{
					f:       "POWER",
					expects: "<float 441.00>",
				},
				{
					f:       "MOD",
					expects: "<int 1>",
				},
				{
					f:       "SUM",
					expects: "<int 23>",
				},
			},
		},
		{
			name: "float, then 1",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2.5},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:       "CEILING",
					expects: "<int 3>",
				},
				{
					f:       "FLOOR",
					expects: "<int 2>",
				},
				{
					f:       "POWER",
					expects: "<float 2.50>",
				},
				{
					f:       "MOD",
					expects: "<float 0.50>",
				},
				{
					f:       "SUM",
					expects: "<float 3.50>",
				},
			},
		},
		{
			name: "float and smaller int",
			input: []ast.Expression{
				ast.FloatExpression{Value: 126.55},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:       "CEILING",
					expects: "<int 129>",
				},
				{
					f:       "FLOOR",
					expects: "<int 126>",
				},
				{
					f:       "POWER",
					expects: "<float 2026685.91>",
				},
				{
					f:       "MOD",
					expects: "<float 0.55>",
				},
				{
					f:       "SUM",
					expects: "<float 129.55>",
				},
			},
		},
		{
			name: "boolean",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "PRODUCT(values:number...):number received invalid argument <bool true> in (PRODUCT <bool true>), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument <bool true> in (AVERAGE <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument <bool true> in (MAX <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument <bool true> in (MAXA <bool true>), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument <bool true> in (MIN <bool true>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument <bool true> in (MINA <bool true>), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number received invalid argument <bool true> in (ABS <bool true>), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number received invalid argument <bool true> in (CEILING <bool true>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number received invalid argument <bool true> in (FLOOR <bool true>), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number received invalid argument <bool true> in (ROUND <bool true>), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) received invalid argument <bool true> in (INT <bool true>), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument <bool true> in (SUM <bool true>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in (POWER <bool true>), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in (MOD <bool true>), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument <bool true> in (SQRT <bool true>), at <: input:0:0>",
				},
			},
		},
		{
			name: "int and string",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "PRODUCT(values:number...):number received invalid argument <str \"hello\"> in (PRODUCT <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument <str \"hello\"> in (AVERAGE <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument <str \"hello\"> in (MAX <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument <str \"hello\"> in (MAXA <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument <str \"hello\"> in (MIN <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument <str \"hello\"> in (MINA <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument <str \"hello\"> in (SUM <int 5> <str \"hello\">), at <: input:0:0>",
				},
			},
		},
		{
			name: "unsupported argument in float product",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "PRODUCT(values:number...):number received invalid argument <bool true> in (PRODUCT <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument <bool true> in (AVERAGE <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument <bool true> in (MAX <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument <bool true> in (MAXA <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument <bool true> in (MIN <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument <bool true> in (MINA <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument <bool true> in (SUM <float 5.50> <bool true>), at <: input:0:0>",
				},
			},
		},
		{
			name: "string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "PRODUCT(values:number...):number received invalid argument <str \"hello\"> in (PRODUCT <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument <str \"hello\"> in (AVERAGE <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument <str \"hello\"> in (MAX <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument <str \"hello\"> in (MAXA <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number received invalid argument <str \"hello\"> in (ABS <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number received invalid argument <str \"hello\"> in (CEILING <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number received invalid argument <str \"hello\"> in (FLOOR <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number received invalid argument <str \"hello\"> in (ROUND <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) received invalid argument <str \"hello\"> in (INT <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument <str \"hello\"> in (MIN <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument <str \"hello\"> in (MINA <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in (POWER <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument <str \"hello\"> in (SUM <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in (MOD <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument <str \"hello\"> in (SQRT <str \"hello\">), at <: input:0:0>",
				},
			},
		},
		// MAXA/MINA-specific tests for string parsing
		{
			name: "with string numbers",
			input: []ast.Expression{
				ast.StringExpression{Value: "5"},
				ast.StringExpression{Value: "3.7"},
				ast.IntExpression{Value: 2},
			},
			cases: []inputCase{
				{
					f:       "MAXA",
					expects: "<float 5.00>",
				},
				{
					f:       "MINA",
					expects: "<float 2.00>",
				},
			},
		},
		{
			name: "with integer strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "10"},
				ast.StringExpression{Value: "20"},
				ast.IntExpression{Value: 15},
			},
			cases: []inputCase{
				{
					f:       "MAXA",
					expects: "<int 20>",
				},
				{
					f:       "MINA",
					expects: "<int 10>",
				},
			},
		},
		{
			name: "with mixed valid strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "  42  "},
				ast.StringExpression{Value: "-15"},
				ast.FloatExpression{Value: 3.14},
			},
			cases: []inputCase{
				{
					f:       "MAXA",
					expects: "<float 42.00>",
				},
				{
					f:       "MINA",
					expects: "<float -15.00>",
				},
			},
		},
		{
			name: "with invalid string",
			input: []ast.Expression{
				ast.StringExpression{Value: "not_a_number"},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument <str \"not_a_number\"> in (MAXA <str \"not_a_number\"> <int 5>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument <str \"not_a_number\"> in (MINA <str \"not_a_number\"> <int 5>), at <: input:0:0>",
				},
			},
		},
		// MOD function specific tests
		{
			name: "MOD: positive integers",
			input: []ast.Expression{
				ast.IntExpression{Value: 17},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int 2>",
				},
			},
		},
		{
			name: "MOD: negative dividend",
			input: []ast.Expression{
				ast.IntExpression{Value: -17},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int -2>",
				},
			},
		},
		{
			name: "MOD: negative divisor",
			input: []ast.Expression{
				ast.IntExpression{Value: 17},
				ast.IntExpression{Value: -5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int 2>",
				},
			},
		},
		{
			name: "MOD: both negative",
			input: []ast.Expression{
				ast.IntExpression{Value: -17},
				ast.IntExpression{Value: -5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int -2>",
				},
			},
		},
		{
			name: "MOD: exact division",
			input: []ast.Expression{
				ast.IntExpression{Value: 15},
				ast.IntExpression{Value: 5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int 0>",
				},
			},
		},
		{
			name: "MOD: with floats",
			input: []ast.Expression{
				ast.FloatExpression{Value: 7.5},
				ast.FloatExpression{Value: 2.5},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<float 0.00>",
				},
			},
		},
		{
			name: "MOD: mixed int and float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 10.7},
				ast.IntExpression{Value: 3},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<float 1.70>",
				},
			},
		},
		{
			name: "MOD: division by one",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.7},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<float 0.70>",
				},
			},
		},
		{
			name: "MOD: small dividend, large divisor",
			input: []ast.Expression{
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:       "MOD",
					expects: "<int 3>",
				},
			},
		},
		// SQRT function specific tests
		{
			name: "SQRT: perfect squares",
			input: []ast.Expression{
				ast.IntExpression{Value: 25},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<int 5>",
				},
			},
		},
		{
			name: "SQRT: perfect square floats",
			input: []ast.Expression{
				ast.FloatExpression{Value: 16.0},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<float 4.00>",
				},
			},
		},
		{
			name: "SQRT: non-perfect squares",
			input: []ast.Expression{
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<int 3>",
				},
			},
		},
		{
			name: "SQRT: decimal values",
			input: []ast.Expression{
				ast.FloatExpression{Value: 6.25},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<float 2.50>",
				},
			},
		},
		{
			name: "SQRT: zero",
			input: []ast.Expression{
				ast.IntExpression{Value: 0},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<int 0>",
				},
			},
		},
		{
			name: "SQRT: one",
			input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<int 1>",
				},
			},
		},
		{
			name: "SQRT: large perfect square",
			input: []ast.Expression{
				ast.IntExpression{Value: 144},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<int 12>",
				},
			},
		},
		{
			name: "SQRT: small decimal",
			input: []ast.Expression{
				ast.FloatExpression{Value: 0.25},
			},
			cases: []inputCase{
				{
					f:       "SQRT",
					expects: "<float 0.50>",
				},
			},
		},
		// COUNT and COUNTA specific tests
		{
			name: "COUNT: mixed numeric and non-numeric types",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.StringExpression{Value: "hello"},
				ast.FloatExpression{Value: 3.14},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 2>",
				},
				{
					f:       "COUNTA",
					expects: "<int 4>",
				},
			},
		},
		{
			name: "COUNT: only strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: "123"},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 0>",
				},
				{
					f:       "COUNTA",
					expects: "<int 3>",
				},
			},
		},
		{
			name: "COUNT: only booleans",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 0>",
				},
				{
					f:       "COUNTA",
					expects: "<int 2>",
				},
			},
		},
		{
			name: "COUNTA: with empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "world"},
				ast.StringExpression{Value: ""},
				ast.IntExpression{Value: 42},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 1>",
				},
				{
					f:       "COUNTA",
					expects: "<int 3>",
				},
			},
		},
		{
			name: "COUNT and COUNTA: all empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 0>",
				},
				{
					f:       "COUNTA",
					expects: "<int 0>",
				},
			},
		},
		{
			name: "COUNT: zeros should be counted",
			input: []ast.Expression{
				ast.IntExpression{Value: 0},
				ast.FloatExpression{Value: 0.0},
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 3>",
				},
				{
					f:       "COUNTA",
					expects: "<int 3>",
				},
			},
		},
		{
			name: "COUNT: complex mixed scenario",
			input: []ast.Expression{
				ast.IntExpression{Value: -5},
				ast.FloatExpression{Value: 2.5},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "text"},
				ast.DateExpression{},
				ast.IntExpression{Value: 0},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:       "COUNT",
					expects: "<int 4>",
				},
				{
					f:       "COUNTA",
					expects: "<int 5>",
				},
			},
		},
	}

	for _, tc := range testcases {
		for _, c := range tc.cases {
			t.Run(tc.name+":"+c.f, func(t *testing.T) {
				result, err := DispatchMap[c.f](ast.CallExpression{
					Identifier: ast.IdentifierExpression{
						Value: c.f,
						Token: lexer.Token{Literal: c.f},
					}, Arguments: tc.input,
				}, tc.input...)

				if c.error != "" {
					if err == nil {
						t.Errorf("Expected error %q got result: %v", c.error, result)
						return
					}
					if err.Error() != c.error {
						t.Errorf("Expected error %q, got %q", c.error, err.Error())
					}
					return
				}

				if err != nil {
					t.Errorf("Unexpects error: %v", err)
					return
				}

				if result.String() != c.expects {
					t.Errorf("Expected %s, got %s", c.expects, result.String())
				}
			})
		}
	}
}
