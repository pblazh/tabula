package functions

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestStringFunctions(t *testing.T) {
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
					f:        "CONCATENATE",
					expected: `<str "">`,
				},
			},
		},
		// Single string
		{
			name: "single string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "hello">`,
				},
			},
		},
		// Multiple strings
		{
			name: "multiple strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.StringExpression{Value: " "},
				ast.StringExpression{Value: "world"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "hello world">`,
				},
			},
		},
		// Empty strings
		{
			name: "with empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: "start"},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: "end"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "startend">`,
				},
			},
		},
		// All empty strings
		{
			name: "all empty strings",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: `<str "">`,
				},
			},
		},
		// Special characters
		{
			name: "special characters",
			input: []ast.Expression{
				ast.StringExpression{Value: "Line1\n"},
				ast.StringExpression{Value: "Line2\t"},
				ast.StringExpression{Value: "Line3"},
			},
			cases: []inputCase{
				{
					f:        "CONCATENATE",
					expected: "<str \"Line1\nLine2\tLine3\">",
				},
			},
		},
		// Numbers and other types (should error)
		{
			name: "string and integer",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 42},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <int 42> in (CONCATENATE <str "hello"> <int 42>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "string and float",
			input: []ast.Expression{
				ast.StringExpression{Value: "value: "},
				ast.FloatExpression{Value: 3.14},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <float 3.14> in (CONCATENATE <str "value: "> <float 3.14>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "string and boolean",
			input: []ast.Expression{
				ast.StringExpression{Value: "result: "},
				ast.BooleanExpression{Value: true},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <bool true> in (CONCATENATE <str "result: "> <bool true>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "integer only",
			input: []ast.Expression{
				ast.IntExpression{Value: 123},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <int 123> in (CONCATENATE <int 123>), at <: input:0:0>`,
				},
			},
		},
		{
			name: "boolean only",
			input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:     "CONCATENATE",
					error: `CONCATENATE(string...) got a wrong argument <bool false> in (CONCATENATE <bool false>), at <: input:0:0>`,
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

