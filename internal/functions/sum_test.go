package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		name          string
		input         []ast.Expression
		expected      string
		expectedError string
	}{
		// Empty input
		{
			name:     "empty input returns zero",
			input:    []ast.Expression{},
			expected: "<int 0>",
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
				ast.IntExpression{Value: 10},
				ast.IntExpression{Value: -3},
				ast.IntExpression{Value: 2},
			},
			expected: "<int 9>",
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
				ast.FloatExpression{Value: 1.1},
				ast.FloatExpression{Value: 2.2},
				ast.FloatExpression{Value: 3.3},
			},
			expected: "<float 6.60>",
		},
		// Mixed int and float operations (result type determined by first argument)
		{
			name: "int first, then float - result is int",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.FloatExpression{Value: 3.7},
			},
			expected: "<int 8>",
		},
		{
			name: "float first, then int - result is float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.IntExpression{Value: 3},
			},
			expected: "<float 8.50>",
		},
		// String operations
		{
			name: "single string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			expected: "<str \"hello\">",
		},
		{
			name: "multiple strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "a"},
				ast.StringExpression{Value: "b"},
				ast.StringExpression{Value: "c"},
			},
			expected: "<str \"abc\">",
		},
		// Error cases
		{
			name: "unsupported boolean first argument",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			expectedError: "unsupported function call (SUM <bool true>)",
		},
		{
			name: "unsupported argument in integer sum",
			input: []ast.Expression{
				ast.IntExpression{Value: 5},
				ast.StringExpression{Value: "hello"},
			},
			expectedError: "unsupported argument <str \"hello\"> for (SUM <int 5> <str \"hello\">)",
		},
		{
			name: "unsupported argument in float sum",
			input: []ast.Expression{
				ast.FloatExpression{Value: 5.5},
				ast.BooleanExpression{Value: true},
			},
			expectedError: "unsupported argument <bool true> for (SUM <float 5.50> <bool true>)",
		},
		{
			name: "unsupported argument in string sum",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
			},
			expectedError: "unsupported argument <int 42> for (SUM <str \"hello\"> <int 42>)",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Sum(ast.CallExpression{
				Identifier: ast.IdentifierExpression{
					Token: lexer.Token{Literal: "SUM"},
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
