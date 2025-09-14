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
					expects: "1",
				},
				{
					f:       "AVERAGE",
					expects: "0",
				},
				{
					f:       "SUM",
					expects: "0",
				},
				{
					f:       "COUNT",
					expects: "0",
				},
				{
					f:       "COUNTA",
					expects: "0",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number expects 1 argument, got 0 in ABS(), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number expects 2 arguments, got 0 in CEILING(), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number expects 2 arguments, got 0 in FLOOR(), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number expects 2 arguments, got 0 in ROUND(), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expects 1 argument, got 0 in INT(), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 0 in POWER(), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 0 in MOD(), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number expects 1 argument, got 0 in SQRT(), at <: input:0:0>",
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
					expects: "-5",
				},
				{
					f:       "AVERAGE",
					expects: "-5",
				},
				{
					f:       "MAX",
					expects: "-5",
				},
				{
					f:       "MAXA",
					expects: "-5",
				},
				{
					f:       "MIN",
					expects: "-5",
				},
				{
					f:       "MINA",
					expects: "-5",
				},
				{
					f:       "ABS",
					expects: "5",
				},
				{
					f:       "INT",
					expects: "-5",
				},
				{
					f:       "SUM",
					expects: "-5",
				},
				{
					f:       "COUNT",
					expects: "1",
				},
				{
					f:       "COUNTA",
					expects: "1",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument -5 in SQRT(-5), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in POWER(-5), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in MOD(-5), at <: input:0:0>",
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
					expects: "24",
				},
				{
					f:       "AVERAGE",
					expects: "3",
				},
				{
					f:       "MAX",
					expects: "4",
				},
				{
					f:       "MAXA",
					expects: "4",
				},
				{
					f:       "MIN",
					expects: "2",
				},
				{
					f:       "MINA",
					expects: "2",
				},
				{
					f:       "SUM",
					expects: "9",
				},
				{
					f:       "COUNT",
					expects: "3",
				},
				{
					f:       "COUNTA",
					expects: "3",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number expects 1 argument, got 3 in ABS(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number expects 2 arguments, got 3 in CEILING(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number expects 2 arguments, got 3 in FLOOR(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number expects 2 arguments, got 3 in ROUND(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expects 1 argument, got 3 in INT(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 3 in POWER(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 3 in MOD(2, 3, 4), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number expects 1 argument, got 3 in SQRT(2, 3, 4), at <: input:0:0>",
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
					expects: "0",
				},
				{
					f:       "AVERAGE",
					expects: "5",
				},
				{
					f:       "MAX",
					expects: "10",
				},
				{
					f:       "MAXA",
					expects: "10",
				},
				{
					f:       "MIN",
					expects: "0",
				},
				{
					f:       "MINA",
					expects: "0",
				},
				{
					f:       "SUM",
					expects: "15",
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
					expects: "-60",
				},
				{
					f:       "AVERAGE",
					expects: "3",
				},
				{
					f:       "MAX",
					expects: "10",
				},
				{
					f:       "MAXA",
					expects: "10",
				},
				{
					f:       "MIN",
					expects: "-3",
				},
				{
					f:       "MINA",
					expects: "-3",
				},
				{
					f:       "SUM",
					expects: "9",
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
					expects: "5.50",
				},
				{
					f:       "AVERAGE",
					expects: "5.50",
				},
				{
					f:       "MAX",
					expects: "5.50",
				},
				{
					f:       "MAXA",
					expects: "5.50",
				},
				{
					f:       "MIN",
					expects: "5.50",
				},
				{
					f:       "MINA",
					expects: "5.50",
				},
				{
					f:       "ABS",
					expects: "5.50",
				},
				{
					f:       "INT",
					expects: "5",
				},
				{
					f:       "SUM",
					expects: "5.50",
				},
				{
					f:       "CEILING",
					expects: "6",
				},
				{
					f:       "FLOOR",
					expects: "5",
				},
				{
					f:       "ROUND",
					expects: "6",
				},
				{
					f:       "SQRT",
					expects: "2.35",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in POWER(5.50), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in MOD(5.50), at <: input:0:0>",
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
					expects: "-3.70",
				},
				{
					f:       "AVERAGE",
					expects: "-3.70",
				},
				{
					f:       "MAX",
					expects: "-3.70",
				},
				{
					f:       "MAXA",
					expects: "-3.70",
				},
				{
					f:       "MIN",
					expects: "-3.70",
				},
				{
					f:       "MINA",
					expects: "-3.70",
				},
				{
					f:       "ABS",
					expects: "3.70",
				},
				{
					f:       "INT",
					expects: "-3",
				},
				{
					f:       "SUM",
					expects: "-3.70",
				},
				{
					f:       "CEILING",
					expects: "-3",
				},
				{
					f:       "FLOOR",
					expects: "-4",
				},
				{
					f:       "ROUND",
					expects: "-4",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument -3.70 in SQRT(-3.70), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in POWER(-3.70), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in MOD(-3.70), at <: input:0:0>",
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
					expects: "9.00",
				},
				{
					f:       "AVERAGE",
					expects: "2.17",
				},
				{
					f:       "MAX",
					expects: "3.00",
				},
				{
					f:       "MAXA",
					expects: "3.00",
				},
				{
					f:       "MIN",
					expects: "1.50",
				},
				{
					f:       "MINA",
					expects: "1.50",
				},
				{
					f:       "SUM",
					expects: "6.50",
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
					expects: "0.00",
				},
				{
					f:       "AVERAGE",
					expects: "2.67",
				},
				{
					f:       "MAX",
					expects: "5.50",
				},
				{
					f:       "MAXA",
					expects: "5.50",
				},
				{
					f:       "MIN",
					expects: "0.00",
				},
				{
					f:       "MINA",
					expects: "0.00",
				},
				{
					f:       "SUM",
					expects: "8.00",
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
					expects: "12.50",
				},
				{
					f:       "AVERAGE",
					expects: "3.75",
				},
				{
					f:       "MAX",
					expects: "5.00",
				},
				{
					f:       "MAXA",
					expects: "5.00",
				},
				{
					f:       "MIN",
					expects: "2.50",
				},
				{
					f:       "MINA",
					expects: "2.50",
				},
				{
					f:       "SUM",
					expects: "7.50",
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
					expects: "10.00",
				},
				{
					f:       "AVERAGE",
					expects: "3.25",
				},
				{
					f:       "MAX",
					expects: "4.00",
				},
				{
					f:       "MAXA",
					expects: "4.00",
				},
				{
					f:       "MIN",
					expects: "2.50",
				},
				{
					f:       "MINA",
					expects: "2.50",
				},
				{
					f:       "POWER",
					expects: "39.06",
				},
				{
					f:       "SUM",
					expects: "6.50",
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
					expects: "22",
				},
				{
					f:       "FLOOR",
					expects: "20",
				},
				{
					f:       "POWER",
					expects: "441.00",
				},
				{
					f:       "MOD",
					expects: "1",
				},
				{
					f:       "SUM",
					expects: "23",
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
					expects: "3",
				},
				{
					f:       "FLOOR",
					expects: "2",
				},
				{
					f:       "POWER",
					expects: "2.50",
				},
				{
					f:       "MOD",
					expects: "0.50",
				},
				{
					f:       "SUM",
					expects: "3.50",
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
					expects: "129",
				},
				{
					f:       "FLOOR",
					expects: "126",
				},
				{
					f:       "POWER",
					expects: "2026685.91",
				},
				{
					f:       "MOD",
					expects: "0.55",
				},
				{
					f:       "SUM",
					expects: "129.55",
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
					error: "PRODUCT(values:number...):number received invalid argument true in PRODUCT(true), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument true in AVERAGE(true), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument true in MAX(true), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument true in MAXA(true), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument true in MIN(true), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument true in MINA(true), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number received invalid argument true in ABS(true), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number received invalid argument true in CEILING(true), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number received invalid argument true in FLOOR(true), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number received invalid argument true in ROUND(true), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) received invalid argument true in INT(true), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument true in SUM(true), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in POWER(true), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in MOD(true), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument true in SQRT(true), at <: input:0:0>",
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
					error: "PRODUCT(values:number...):number received invalid argument \"hello\" in PRODUCT(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument \"hello\" in AVERAGE(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument \"hello\" in MAX(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument \"hello\" in MAXA(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument \"hello\" in MIN(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument \"hello\" in MINA(5, \"hello\"), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument \"hello\" in SUM(5, \"hello\"), at <: input:0:0>",
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
					error: "PRODUCT(values:number...):number received invalid argument true in PRODUCT(5.50, true), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument true in AVERAGE(5.50, true), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument true in MAX(5.50, true), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument true in MAXA(5.50, true), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument true in MIN(5.50, true), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument true in MINA(5.50, true), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument true in SUM(5.50, true), at <: input:0:0>",
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
					error: "PRODUCT(values:number...):number received invalid argument \"hello\" in PRODUCT(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(values:number...):number received invalid argument \"hello\" in AVERAGE(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(values:number...):number received invalid argument \"hello\" in MAX(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(values:number|string...):number received invalid argument \"hello\" in MAXA(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(value:number):number received invalid argument \"hello\" in ABS(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(value:number, significance:[number]):number received invalid argument \"hello\" in CEILING(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(value:number, significance:[number]):number received invalid argument \"hello\" in FLOOR(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "ROUND",
					error: "ROUND(value:number, places:[number]):number received invalid argument \"hello\" in ROUND(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) received invalid argument \"hello\" in INT(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(values:number...):number received invalid argument \"hello\" in MIN(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument \"hello\" in MINA(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(base:number, exponent:number):number expects 2 arguments, got 1 in POWER(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(values:number...):number received invalid argument \"hello\" in SUM(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "MOD",
					error: "MOD(dividend:number, divisor:number):number expects 2 arguments, got 1 in MOD(\"hello\"), at <: input:0:0>",
				},
				{
					f:     "SQRT",
					error: "SQRT(value:number):number received invalid argument \"hello\" in SQRT(\"hello\"), at <: input:0:0>",
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
					expects: "5.00",
				},
				{
					f:       "MINA",
					expects: "2.00",
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
					expects: "20",
				},
				{
					f:       "MINA",
					expects: "10",
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
					expects: "42.00",
				},
				{
					f:       "MINA",
					expects: "-15.00",
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
					error: "MAXA(values:number|string...):number received invalid argument \"not_a_number\" in MAXA(\"not_a_number\", 5), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(values:number|string...):number received invalid argument \"not_a_number\" in MINA(\"not_a_number\", 5), at <: input:0:0>",
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
					expects: "2",
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
					expects: "-2",
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
					expects: "2",
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
					expects: "-2",
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
					expects: "0",
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
					expects: "0.00",
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
					expects: "1.70",
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
					expects: "0.70",
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
					expects: "3",
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
					expects: "5",
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
					expects: "4.00",
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
					expects: "3",
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
					expects: "2.50",
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
					expects: "0",
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
					expects: "1",
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
					expects: "12",
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
					expects: "0.50",
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
					expects: "2",
				},
				{
					f:       "COUNTA",
					expects: "4",
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
					expects: "0",
				},
				{
					f:       "COUNTA",
					expects: "3",
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
					expects: "0",
				},
				{
					f:       "COUNTA",
					expects: "2",
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
					expects: "1",
				},
				{
					f:       "COUNTA",
					expects: "3",
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
					expects: "0",
				},
				{
					f:       "COUNTA",
					expects: "0",
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
					expects: "3",
				},
				{
					f:       "COUNTA",
					expects: "3",
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
					expects: "4",
				},
				{
					f:       "COUNTA",
					expects: "5",
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
