package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

type inputCase struct {
	f        string
	expected string
	error    string
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
					f:        "PRODUCT",
					expected: "<int 1>",
				},
				{
					f:        "AVERAGE",
					expected: "<int 0>",
				},
				{
					f:        "MAX",
					expected: "<int -9223372036854775808>",
				},
				{
					f:        "MAXA",
					expected: "<int -9223372036854775808>",
				},
				{
					f:        "MIN",
					expected: "<int 9223372036854775807>",
				},
				{
					f:        "MINA",
					expected: "<int 9223372036854775807>",
				},
				{
					f:        "SUM",
					expected: "<int 0>",
				},
				{
					f:     "ABS",
					error: "ABS(number) expected 1 argument, but got 0 in (ABS), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 0 in (CEILING), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 0 in (FLOOR), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expected 1 argument, but got 0 in (INT), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 0 in (POWER), at <: input:0:0>",
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
					f:        "PRODUCT",
					expected: "<int -5>",
				},
				{
					f:        "AVERAGE",
					expected: "<int -5>",
				},
				{
					f:        "MAX",
					expected: "<int -5>",
				},
				{
					f:        "MAXA",
					expected: "<int -5>",
				},
				{
					f:        "MIN",
					expected: "<int -5>",
				},
				{
					f:        "MINA",
					expected: "<int -5>",
				},
				{
					f:        "ABS",
					expected: "<int 5>",
				},
				{
					f:        "INT",
					expected: "<int -5>",
				},
				{
					f:        "SUM",
					expected: "<int -5>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 1 in (POWER <int -5>), at <: input:0:0>",
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
					f:        "PRODUCT",
					expected: "<int 24>",
				},
				{
					f:        "AVERAGE",
					expected: "<int 3>",
				},
				{
					f:        "MAX",
					expected: "<int 4>",
				},
				{
					f:        "MAXA",
					expected: "<int 4>",
				},
				{
					f:        "MIN",
					expected: "<int 2>",
				},
				{
					f:        "MINA",
					expected: "<int 2>",
				},
				{
					f:        "SUM",
					expected: "<int 9>",
				},
				{
					f:     "ABS",
					error: "ABS(number) expected 1 argument, but got 3 in (ABS <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 3 in (CEILING <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 3 in (FLOOR <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) expected 1 argument, but got 3 in (INT <int 2> <int 3> <int 4>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 3 in (POWER <int 2> <int 3> <int 4>), at <: input:0:0>",
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
					f:        "PRODUCT",
					expected: "<int 0>",
				},
				{
					f:        "AVERAGE",
					expected: "<int 5>",
				},
				{
					f:        "MAX",
					expected: "<int 10>",
				},
				{
					f:        "MAXA",
					expected: "<int 10>",
				},
				{
					f:        "MIN",
					expected: "<int 0>",
				},
				{
					f:        "MINA",
					expected: "<int 0>",
				},
				{
					f:        "SUM",
					expected: "<int 15>",
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
					f:        "PRODUCT",
					expected: "<int -60>",
				},
				{
					f:        "AVERAGE",
					expected: "<int 3>",
				},
				{
					f:        "MAX",
					expected: "<int 10>",
				},
				{
					f:        "MAXA",
					expected: "<int 10>",
				},
				{
					f:        "MIN",
					expected: "<int -3>",
				},
				{
					f:        "MINA",
					expected: "<int -3>",
				},
				{
					f:        "SUM",
					expected: "<int 9>",
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
					f:        "PRODUCT",
					expected: "<float 5.50>",
				},
				{
					f:        "AVERAGE",
					expected: "<float 5.50>",
				},
				{
					f:        "MAX",
					expected: "<float 5.50>",
				},
				{
					f:        "MAXA",
					expected: "<float 5.50>",
				},
				{
					f:        "MIN",
					expected: "<float 5.50>",
				},
				{
					f:        "MINA",
					expected: "<float 5.50>",
				},
				{
					f:        "ABS",
					expected: "<float 5.50>",
				},
				{
					f:        "INT",
					expected: "<int 5>",
				},
				{
					f:        "SUM",
					expected: "<float 5.50>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 1 in (CEILING <float 5.50>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 1 in (FLOOR <float 5.50>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 1 in (POWER <float 5.50>), at <: input:0:0>",
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
					f:        "PRODUCT",
					expected: "<float -3.70>",
				},
				{
					f:        "AVERAGE",
					expected: "<float -3.70>",
				},
				{
					f:        "MAX",
					expected: "<float -3.70>",
				},
				{
					f:        "MAXA",
					expected: "<float -3.70>",
				},
				{
					f:        "MIN",
					expected: "<float -3.70>",
				},
				{
					f:        "MINA",
					expected: "<float -3.70>",
				},
				{
					f:        "ABS",
					expected: "<float 3.70>",
				},
				{
					f:        "INT",
					expected: "<int -3>",
				},
				{
					f:        "SUM",
					expected: "<float -3.70>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 1 in (CEILING <float -3.70>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 1 in (FLOOR <float -3.70>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 1 in (POWER <float -3.70>), at <: input:0:0>",
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
					f:        "PRODUCT",
					expected: "<float 9.00>",
				},
				{
					f:        "AVERAGE",
					expected: "<float 2.17>",
				},
				{
					f:        "MAX",
					expected: "<float 3.00>",
				},
				{
					f:        "MAXA",
					expected: "<float 3.00>",
				},
				{
					f:        "MIN",
					expected: "<float 1.50>",
				},
				{
					f:        "MINA",
					expected: "<float 1.50>",
				},
				{
					f:        "SUM",
					expected: "<float 6.50>",
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
					f:        "PRODUCT",
					expected: "<float 0.00>",
				},
				{
					f:        "AVERAGE",
					expected: "<float 2.67>",
				},
				{
					f:        "MAX",
					expected: "<float 5.50>",
				},
				{
					f:        "MAXA",
					expected: "<float 5.50>",
				},
				{
					f:        "MIN",
					expected: "<float 0.00>",
				},
				{
					f:        "MINA",
					expected: "<float 0.00>",
				},
				{
					f:        "SUM",
					expected: "<float 8.00>",
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
					f:        "PRODUCT",
					expected: "<float 12.50>",
				},
				{
					f:        "AVERAGE",
					expected: "<float 3.75>",
				},
				{
					f:        "MAX",
					expected: "<float 5.00>",
				},
				{
					f:        "MAXA",
					expected: "<float 5.00>",
				},
				{
					f:        "MIN",
					expected: "<float 2.50>",
				},
				{
					f:        "MINA",
					expected: "<float 2.50>",
				},
				{
					f:        "SUM",
					expected: "<float 7.50>",
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
					f:        "PRODUCT",
					expected: "<float 10.00>",
				},
				{
					f:        "AVERAGE",
					expected: "<float 3.25>",
				},
				{
					f:        "MAX",
					expected: "<float 4.00>",
				},
				{
					f:        "MAXA",
					expected: "<float 4.00>",
				},
				{
					f:        "MIN",
					expected: "<float 2.50>",
				},
				{
					f:        "MINA",
					expected: "<float 2.50>",
				},
				{
					f:        "POWER",
					expected: "<float 39.06>",
				},
				{
					f:        "SUM",
					expected: "<float 6.50>",
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
					f:        "CEILING",
					expected: "<float 22.00>",
				},
				{
					f:        "FLOOR",
					expected: "<float 20.00>",
				},
				{
					f:        "POWER",
					expected: "<float 441.00>",
				},
				{
					f:        "SUM",
					expected: "<int 23>",
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
					f:        "CEILING",
					expected: "<float 3.00>",
				},
				{
					f:        "FLOOR",
					expected: "<float 2.00>",
				},
				{
					f:        "POWER",
					expected: "<float 2.50>",
				},
				{
					f:        "SUM",
					expected: "<float 3.50>",
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
					f:        "CEILING",
					expected: "<float 129.00>",
				},
				{
					f:        "FLOOR",
					expected: "<float 126.00>",
				},
				{
					f:        "POWER",
					expected: "<float 2026685.91>",
				},
				{
					f:        "SUM",
					expected: "<float 129.55>",
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
					error: "PRODUCT(number...) got a wrong argument <bool true> in (PRODUCT <bool true>), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(number...) got a wrong argument <bool true> in (AVERAGE <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(number...) got a wrong argument <bool true> in (MAX <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(number...) got a wrong argument <bool true> in (MAXA <bool true>), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(number...) got a wrong argument <bool true> in (MIN <bool true>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(number...) got a wrong argument <bool true> in (MINA <bool true>), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(number) got a wrong argument <bool true> in (ABS <bool true>), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 1 in (CEILING <bool true>), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 1 in (FLOOR <bool true>), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) got a wrong argument <bool true> in (INT <bool true>), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(number...) got a wrong argument <bool true> in (SUM <bool true>), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 1 in (POWER <bool true>), at <: input:0:0>",
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
					error: "PRODUCT(number...) got a wrong argument <str \"hello\"> in (PRODUCT <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(number...) got a wrong argument <str \"hello\"> in (AVERAGE <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(number...) got a wrong argument <str \"hello\"> in (MAX <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(number...) got a wrong argument <str \"hello\"> in (MAXA <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(number...) got a wrong argument <str \"hello\"> in (MIN <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(number...) got a wrong argument <str \"hello\"> in (MINA <int 5> <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(number...) got a wrong argument <str \"hello\"> in (SUM <int 5> <str \"hello\">), at <: input:0:0>",
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
					error: "PRODUCT(number...) got a wrong argument <bool true> in (PRODUCT <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(number...) got a wrong argument <bool true> in (AVERAGE <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(number...) got a wrong argument <bool true> in (MAX <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(number...) got a wrong argument <bool true> in (MAXA <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(number...) got a wrong argument <bool true> in (MIN <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(number...) got a wrong argument <bool true> in (MINA <float 5.50> <bool true>), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(number...) got a wrong argument <bool true> in (SUM <float 5.50> <bool true>), at <: input:0:0>",
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
					error: "PRODUCT(number...) got a wrong argument <str \"hello\"> in (PRODUCT <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "AVERAGE",
					error: "AVERAGE(number...) got a wrong argument <str \"hello\"> in (AVERAGE <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAX",
					error: "MAX(number...) got a wrong argument <str \"hello\"> in (MAX <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MAXA",
					error: "MAXA(number...) got a wrong argument <str \"hello\"> in (MAXA <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "ABS",
					error: "ABS(number) got a wrong argument <str \"hello\"> in (ABS <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "CEILING",
					error: "CEILING(number, number) expected 2 arguments, but got 1 in (CEILING <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "FLOOR",
					error: "FLOOR(number, number) expected 2 arguments, but got 1 in (FLOOR <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "INT",
					error: "INT(number) got a wrong argument <str \"hello\"> in (INT <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MIN",
					error: "MIN(number...) got a wrong argument <str \"hello\"> in (MIN <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(number...) got a wrong argument <str \"hello\"> in (MINA <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "POWER",
					error: "POWER(number, number) expected 2 arguments, but got 1 in (POWER <str \"hello\">), at <: input:0:0>",
				},
				{
					f:     "SUM",
					error: "SUM(number...) got a wrong argument <str \"hello\"> in (SUM <str \"hello\">), at <: input:0:0>",
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
					f:        "MAXA",
					expected: "<float 5.00>",
				},
				{
					f:        "MINA",
					expected: "<float 2.00>",
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
					f:        "MAXA",
					expected: "<int 20>",
				},
				{
					f:        "MINA",
					expected: "<int 10>",
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
					f:        "MAXA",
					expected: "<float 42.00>",
				},
				{
					f:        "MINA",
					expected: "<float -15.00>",
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
					error: "MAXA(number...) got a wrong argument <str \"not_a_number\"> in (MAXA <str \"not_a_number\"> <int 5>), at <: input:0:0>",
				},
				{
					f:     "MINA",
					error: "MINA(number...) got a wrong argument <str \"not_a_number\"> in (MINA <str \"not_a_number\"> <int 5>), at <: input:0:0>",
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
						t.Errorf("Expected error %q but got result: %v", c.error, result)
						return
					}
					if err.Error() != c.error {
						t.Errorf("Expected error %q, got %q", c.error, err.Error())
					}
					return
				}

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				if result.String() != c.expected {
					t.Errorf("Expected %s, got %s", c.expected, result.String())
				}
			})
		}
	}
}
