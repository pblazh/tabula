package functions

import (
	"testing"
	"time"

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
			name: "single string",
			input: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool true>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool false>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "empty string",
			input: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool true>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool false>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool true>`,
				},
			},
		},
		{
			name: "single int",
			input: []ast.Expression{
				ast.IntExpression{Value: 7},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool true>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool false>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool false>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "single float",
			input: []ast.Expression{
				ast.FloatExpression{Value: 7.4},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool true>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool false>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool false>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "single bloolean",
			input: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool false>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool true>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool false>`,
				},
			},
		},
		{
			name: "single date",
			input: []ast.Expression{
				ast.DateExpression{Value: time.Now()},
			},
			cases: []inputCase{
				{
					f:        "ISNUMBER",
					expected: `<bool false>`,
				},
				{
					f:        "ISTEXT",
					expected: `<bool false>`,
				},
				{
					f:        "ISLOGICAL",
					expected: `<bool false>`,
				},
				{
					f:        "ISBLANK",
					expected: `<bool false>`,
				},
			},
		},
		{
			name:  "empty input",
			input: []ast.Expression{},
			cases: []inputCase{
				{
					f:     "ISNUMBER",
					error: "ISNUMBER(any) expected 1 argument, but got 0 in (ISNUMBER), at <: input:0:0>",
				},
				{
					f:     "ISTEXT",
					error: "ISTEXT(any) expected 1 argument, but got 0 in (ISTEXT), at <: input:0:0>",
				},
				{
					f:     "ISLOGICAL",
					error: "ISLOGICAL(any) expected 1 argument, but got 0 in (ISLOGICAL), at <: input:0:0>",
				},
				{
					f:     "ISBLANK",
					error: "ISBLANK(any) expected 1 argument, but got 0 in (ISBLANK), at <: input:0:0>",
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
					error: "ISNUMBER(any) expected 1 argument, but got 2 in (ISNUMBER <str \"2025-08-17 15:39\"> <int 39>), at <: input:0:0>",
				},
				{
					f:     "ISTEXT",
					error: "ISTEXT(any) expected 1 argument, but got 2 in (ISTEXT <str \"2025-08-17 15:39\"> <int 39>), at <: input:0:0>",
				},
				{
					f:     "ISLOGICAL",
					error: "ISLOGICAL(any) expected 1 argument, but got 2 in (ISLOGICAL <str \"2025-08-17 15:39\"> <int 39>), at <: input:0:0>",
				},
				{
					f:     "ISBLANK",
					error: "ISBLANK(any) expected 1 argument, but got 2 in (ISBLANK <str \"2025-08-17 15:39\"> <int 39>), at <: input:0:0>",
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
