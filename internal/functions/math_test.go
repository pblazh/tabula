package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestProduct(t *testing.T) {
	testcases := []struct {
		name          string
		input         []ast.Expression
		expected      string
		expectedError string
	}{
		// Empty input
		{
			name:     "empty input returns one",
			input:    []ast.Expression{},
			expected: "<int 1>",
		},
		// Integer operations
		{
			name: "single integer",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
			},
			expected: "<int 5>",
		},
		{
			name: "multiple integers",
			input: []ast.Expression{
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
				ast.IntExpression{Value: 4},
			},
			expected: "<int 24>",
		},
		{
			name: "integers with zero",
			input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: 0},
				ast.IntExpression{Value: 5},
			},
			expected: "<int 0>",
		},
		{
			name: "integers with negative values",
			input: []ast.Expression{
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: -3},
				ast.IntExpression{Value: 2},
			},
			expected: "<int -60>",
		},
		// Float operations
		{
			name: "single float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
			},
			expected: "<float 5.50>",
		},
		{
			name: "multiple floats",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2.0},
				ast.FloatExpression{Value: 1.5},
				ast.FloatExpression{Value: 3.0},
			},
			expected: "<float 9.00>",
		},
		{
			name: "floats with zero",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.FloatExpression{Value: 0.0},
				ast.FloatExpression{Value: 2.5},
			},
			expected: "<float 0.00>",
		},
		// Mixed int and float operations (result type determined by first argument)
		{
			name: "int first, then float - result is int",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.FloatExpression{Value: 2.5},
			},
			expected: "<int 10>",
		},
		{
			name: "float first, then int - result is float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2.5},
				ast.IntExpression{Value: 4},
			},
			expected: "<float 10.00>",
		},
		// Error cases
		{
			name: "unsupported boolean first argument",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			expectedError: "unsupported function call (PRODUCT <bool true>)",
		},
		{
			name: "unsupported argument in integer product",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.StringExpression{Value: "hello"},
			},
			expectedError: "unsupported argument <str \"hello\"> for for (PRODUCT <int 5> <str \"hello\">)",
		},
		{
			name: "unsupported argument in float product",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.BooleanExpression{Value: true},
			},
			expectedError: "unsupported argument <bool true> for for (PRODUCT <float 5.50> <bool true>)",
		},
		{
			name: "unsupported string argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			expectedError: "unsupported function call (PRODUCT <str \"hello\">)",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Product(ast.CallExpression{
				Identifier: ast.IdentifierExpression{
					Token: lexer.Token{Literal: "PRODUCT"},
				}, Arguments: tc.input,
			}, tc.input...)

			if tc.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error %q but got result: %v", tc.expectedError, result)
					return
				}
				if err.Error() != tc.expectedError {
					t.Errorf("Expected error %q, got %q", tc.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

