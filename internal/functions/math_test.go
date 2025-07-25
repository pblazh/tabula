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

func TestProduct(t *testing.T) {
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
					f:     "ABS",
					error: "(ABS) expected 1 arguments, but got 0",
				},
				{
					f:     "CEILING",
					error: "(CEILING) expected 2 arguments, but got 0",
				},
				{
					f:     "FLOOR",
					error: "(FLOOR) expected 2 arguments, but got 0",
				},
				{
					f:     "INT",
					error: "(INT) expected 1 arguments, but got 0",
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
					f:        "ABS",
					expected: "<int 5>",
				},
				{
					f:        "INT",
					expected: "<int -5>",
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
					f:     "ABS",
					error: "(ABS <int 2> <int 3> <int 4>) expected 1 arguments, but got 3",
				},
				{
					f:     "CEILING",
					error: "(CEILING <int 2> <int 3> <int 4>) expected 2 arguments, but got 3",
				},
				{
					f:     "FLOOR",
					error: "(FLOOR <int 2> <int 3> <int 4>) expected 2 arguments, but got 3",
				},
				{
					f:     "INT",
					error: "(INT <int 2> <int 3> <int 4>) expected 1 arguments, but got 3",
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
					f:        "ABS",
					expected: "<float 5.50>",
				},
				{
					f:        "INT",
					expected: "<int 5>",
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
					f:        "ABS",
					expected: "<float 3.70>",
				},
				{
					f:        "INT",
					expected: "<int -3>",
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
			},
		},
		// Mixed int and float operations (result type determined by first argument)
		{
			name: "int first, then float - result is int",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.FloatExpression{Value: 2.5},
			},
			cases: []inputCase{
				{
					f:        "PRODUCT",
					expected: "<int 10>",
				},
				{
					f:        "AVERAGE",
					expected: "<int 3>",
				},
			},
		},
		{
			name: "float first, then int - result is float",
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
			},
		},
		// Rounding
		{
			name: "float first, then 1",
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
			},
		},
		{
			name: "float first, then 10",
			input: []ast.Expression{
				ast.FloatExpression{Value: 126.55},
				ast.IntExpression{Value: 10},
			},
			cases: []inputCase{
				{
					f:        "CEILING",
					expected: "<float 130.00>",
				},
				{
					f:        "FLOOR",
					expected: "<float 120.00>",
				},
			},
		},
		// Error cases
		{
			name: "unsupported boolean first argument",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "unsupported function call (PRODUCT <bool true>)",
				},
				{
					f:     "AVERAGE",
					error: "unsupported function call (AVERAGE <bool true>)",
				},
				{
					f:     "ABS",
					error: "unsupported function call (ABS <bool true>)",
				},
				{
					f:     "CEILING",
					error: "(CEILING <bool true>) expected 2 arguments, but got 1",
				},
				{
					f:     "INT",
					error: "unsupported function call (INT <bool true>)",
				},
			},
		},
		{
			name: "unsupported argument in integer product",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "unsupported argument <str \"hello\"> for for (PRODUCT <int 5> <str \"hello\">)",
				},
				{
					f:     "AVERAGE",
					error: "unsupported argument <str \"hello\"> for for (AVERAGE <int 5> <str \"hello\">)",
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
					error: "unsupported argument <bool true> for for (PRODUCT <float 5.50> <bool true>)",
				},
				{
					f:     "AVERAGE",
					error: "unsupported argument <bool true> for for (AVERAGE <float 5.50> <bool true>)",
				},
			},
		},
		{
			name: "unsupported string argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:     "PRODUCT",
					error: "unsupported function call (PRODUCT <str \"hello\">)",
				},
				{
					f:     "AVERAGE",
					error: "unsupported function call (AVERAGE <str \"hello\">)",
				},
				{
					f:     "ABS",
					error: "unsupported function call (ABS <str \"hello\">)",
				},
				{
					f:     "CEILING",
					error: "(CEILING <str \"hello\">) expected 2 arguments, but got 1",
				},
				{
					f:     "INT",
					error: "unsupported function call (INT <str \"hello\">)",
				},
			},
		},
	}

	for _, tc := range testcases {
		for _, c := range tc.cases {
			t.Run(tc.name+":"+c.f, func(t *testing.T) {
				result, err := DispatchMap[c.f](ast.CallExpression{
					Identifier: ast.IdentifierExpression{
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
