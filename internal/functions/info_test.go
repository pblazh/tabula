package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestInfoFunctions(t *testing.T) {
	testcases := []struct {
		name  string
		input []ast.Expression
		cases []inputCase
	}{
		{
			name:  "empty input",
			input: []ast.Expression{},
			cases: []inputCase{
				{
					f:     "ISNUMBER",
					error: "ISNUMBER(number) expected 1 argument, but got 0 in (ISNUMBER), at <: input:0:0>",
				},
			},
		},
		{
			name: "single string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "single int",
			input: []ast.Expression{
				ast.StringExpression{Value: "7"},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool true>`,
				},
			},
		},
		{
			name: "single float",
			input: []ast.Expression{
				ast.StringExpression{Value: "7.4"},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool true>`,
				},
			},
		},
		{
			name: "single date",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 15:39"},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "single float node",
			input: []ast.Expression{
				ast.FloatExpression{Value: 2025},
			},
			cases: []inputCase{
				{
					f:     "ISNUMBER",
					error: "ISNUMBER(number) got a wrong argument <float 2025.00> in (ISNUMBER <float 2025.00>), at <: input:0:0>",
				},
			},
		},
		{
			name: "multiple values",
			input: []ast.Expression{
				ast.StringExpression{Value: "2025-08-17 15:39"},
				ast.IntExpression{Value: 39},
			},
			cases: []inputCase{
				{
					f:     "ISNUMBER",
					error: "ISNUMBER(number) expected 1 argument, but got 2 in (ISNUMBER <str \"2025-08-17 15:39\"> <int 39>), at <: input:0:0>",
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
