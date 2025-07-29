package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestBooleanFunctions(t *testing.T) {
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
					f:     "NOT",
					error: "NOT(boolean) expected 1 argument, but got 0 in (NOT), at <: input:0:0>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 0 in (AND), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 0 in (OR), at <: input:0:0>",
				},
				{
					f:        "TRUE",
					expected: "<bool true>",
				},
				{
					f:        "FALSE",
					expected: "<bool false>",
				},
			},
		},
		// Single boolean true
		{
			name: "single boolean true",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:        "NOT",
					expected: "<bool false>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <bool true>), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <bool true>), at <: input:0:0>",
				},
				{
					f:     "TRUE",
					error: "TRUE() expected 0 arguments, but got 1 in (TRUE <bool true>), at <: input:0:0>",
				},
				{
					f:     "FALSE",
					error: "FALSE() expected 0 arguments, but got 1 in (FALSE <bool true>), at <: input:0:0>",
				},
			},
		},
		// Single boolean false
		{
			name: "single boolean false",
			input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:        "NOT",
					expected: "<bool true>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <bool false>), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <bool false>), at <: input:0:0>",
				},
			},
		},
		// Two booleans: true AND false
		{
			name: "true and false",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:     "NOT",
					error: "NOT(boolean) expected 1 argument, but got 2 in (NOT <bool true> <bool false>), at <: input:0:0>",
				},
				{
					f:        "AND",
					expected: "<bool false>",
				},
				{
					f:        "OR",
					expected: "<bool true>",
				},
			},
		},
		// Two booleans: true AND true
		{
			name: "true and true",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:        "AND",
					expected: "<bool true>",
				},
				{
					f:        "OR",
					expected: "<bool true>",
				},
			},
		},
		// Two booleans: false AND false
		{
			name: "false and false",
			input: []ast.Expression{
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:        "AND",
					expected: "<bool false>",
				},
				{
					f:        "OR",
					expected: "<bool false>",
				},
			},
		},
		// Three booleans (should error for AND/OR)
		{
			name: "three booleans",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.BooleanExpression{Value: false},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 3 in (AND <bool true> <bool false> <bool true>), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 3 in (OR <bool true> <bool false> <bool true>), at <: input:0:0>",
				},
			},
		},
		// String argument (should error)
		{
			name: "string argument",
			input: []ast.Expression{
				ast.StringExpression{Value: "true"},
			},
			cases: []inputCase{
				{
					f:     "NOT",
					error: "NOT(boolean) got a wrong argument <str \"true\"> in (NOT <str \"true\">), at <: input:0:0>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <str \"true\">), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <str \"true\">), at <: input:0:0>",
				},
			},
		},
		// Mixed types for AND/OR
		{
			name: "boolean and string",
			input: []ast.Expression{
				ast.BooleanExpression{Value: true},
				ast.StringExpression{Value: "false"},
			},
			cases: []inputCase{
				{
					f:     "AND",
					error: "AND(boolean, boolean) got a wrong argument <str \"false\"> in (AND <bool true> <str \"false\">), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) got a wrong argument <str \"false\"> in (OR <bool true> <str \"false\">), at <: input:0:0>",
				},
			},
		},
		// Integer argument (should error)
		{
			name: "integer argument",
			input: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			cases: []inputCase{
				{
					f:     "NOT",
					error: "NOT(boolean) got a wrong argument <int 1> in (NOT <int 1>), at <: input:0:0>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <int 1>), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <int 1>), at <: input:0:0>",
				},
			},
		},
		// Float argument (should error)
		{
			name: "float argument",
			input: []ast.Expression{
				ast.FloatExpression{Value: 1.0},
			},
			cases: []inputCase{
				{
					f:     "NOT",
					error: "NOT(boolean) got a wrong argument <float 1.00> in (NOT <float 1.00>), at <: input:0:0>",
				},
				{
					f:     "AND",
					error: "AND(boolean, boolean) expected 2 arguments, but got 1 in (AND <float 1.00>), at <: input:0:0>",
				},
				{
					f:     "OR",
					error: "OR(boolean, boolean) expected 2 arguments, but got 1 in (OR <float 1.00>), at <: input:0:0>",
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
