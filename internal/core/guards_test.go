package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

var call ast.CallExpression = ast.CallExpression{
	Identifier: ast.IdentifierExpression{
		Value: "TEST",
		Token: lexer.Token{Literal: "TEST"},
	},
}

func TestEmptyGuard(t *testing.T) {
	tests := []struct {
		name   string
		values []ast.Expression
	}{
		{
			name:   "no arguments",
			values: []ast.Expression{},
		},
		{
			name: "one argument",
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
		},
		{
			name: "multiple arguments",
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.FloatExpression{Value: 2.5},
				ast.StringExpression{Value: "test"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call.Arguments = tt.values
			err := EmptyGuard(call, tt.values...)
			if err != nil {
				t.Errorf("EmptyGuard should never return error, got: %v", err)
			}
		})
	}
}

func TestMakeArityGuard(t *testing.T) {
	tests := []struct {
		name         string
		arity        int
		values       []ast.Expression
		expectsError string
	}{
		{
			name:         "zero arity with no arguments",
			arity:        0,
			values:       []ast.Expression{},
			expectsError: "",
		},
		{
			name:  "zero arity with arguments",
			arity: 0,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() expects 0 arguments, got 1 in TEST(1), at <: input:0:0>",
		},
		{
			name:  "one arity with correct argument",
			arity: 1,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "",
		},
		{
			name:         "one arity with no arguments",
			arity:        1,
			values:       []ast.Expression{},
			expectsError: "TEST() expects 1 argument, got 0 in TEST(), at <: input:0:0>",
		},
		{
			name:  "one arity with too many arguments",
			arity: 1,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
			},
			expectsError: "TEST() expects 1 argument, got 2 in TEST(1, 2), at <: input:0:0>",
		},
		{
			name:  "two arity with correct arguments",
			arity: 2,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.FloatExpression{Value: 2.5},
			},
			expectsError: "",
		},
		{
			name:  "two arity with one argument",
			arity: 2,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() expects 2 arguments, got 1 in TEST(1), at <: input:0:0>",
		},
		{
			name:  "negative arity should always pass",
			arity: -1,
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
				ast.IntExpression{Value: 3},
			},
			expectsError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guard := MakeArityGuard("TEST()", tt.arity)
			call.Arguments = tt.values
			err := guard(call, tt.values...)

			if tt.expectsError == "" {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error %q, got nil", tt.expectsError)
				} else if err.Error() != tt.expectsError {
					t.Errorf("Expected error %q, got %q", tt.expectsError, err.Error())
				}
			}
		})
	}
}

func TestMakeExactTypesGuard(t *testing.T) {
	tests := []struct {
		name         string
		typeGuards   []typeGuard
		values       []ast.Expression
		expectsError string
	}{
		{
			name:         "no type guards with no arguments",
			typeGuards:   []typeGuard{},
			values:       []ast.Expression{},
			expectsError: "",
		},
		{
			name:       "no type guards with arguments",
			typeGuards: []typeGuard{},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() expects 0 arguments, got 1 in TEST(1), at <: input:0:0>",
		},
		{
			name:       "one numeric guard with int",
			typeGuards: []typeGuard{ast.IsNumeric},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "",
		},
		{
			name:       "one numeric guard with float",
			typeGuards: []typeGuard{ast.IsNumeric},
			values: []ast.Expression{
				ast.FloatExpression{Value: 1.5},
			},
			expectsError: "",
		},
		{
			name:       "one numeric guard with string",
			typeGuards: []typeGuard{ast.IsNumeric},
			values: []ast.Expression{
				ast.StringExpression{Value: "hello"},
			},
			expectsError: "TEST() received invalid argument \"hello\" in TEST(\"hello\"), at <: input:0:0>",
		},
		{
			name:       "one numeric guard with boolean",
			typeGuards: []typeGuard{ast.IsNumeric},
			values: []ast.Expression{
				ast.BooleanExpression{Value: true},
			},
			expectsError: "TEST() received invalid argument true in TEST(true), at <: input:0:0>",
		},
		{
			name:       "two numeric guards with correct types",
			typeGuards: []typeGuard{ast.IsNumeric, ast.IsNumeric},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.FloatExpression{Value: 2.5},
			},
			expectsError: "",
		},
		{
			name:       "two numeric guards with one wrong type",
			typeGuards: []typeGuard{ast.IsNumeric, ast.IsNumeric},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.StringExpression{Value: "hello"},
			},
			expectsError: "TEST() received invalid argument \"hello\" in TEST(1, \"hello\"), at <: input:0:0>",
		},
		{
			name:       "mixed type guards with correct types",
			typeGuards: []typeGuard{ast.IsNumeric, ast.IsString},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.StringExpression{Value: "hello"},
			},
			expectsError: "",
		},
		{
			name:       "mixed type guards with wrong types",
			typeGuards: []typeGuard{ast.IsNumeric, ast.IsString},
			values: []ast.Expression{
				ast.StringExpression{Value: "hello"},
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() received invalid argument \"hello\" in TEST(\"hello\", 1), at <: input:0:0>",
		},
		{
			name:       "wrong arity - too few arguments",
			typeGuards: []typeGuard{ast.IsNumeric, ast.IsNumeric},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() expects 2 arguments, got 1 in TEST(1), at <: input:0:0>",
		},
		{
			name:       "wrong arity - too many arguments",
			typeGuards: []typeGuard{ast.IsNumeric},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
				ast.IntExpression{Value: 2},
			},
			expectsError: "TEST() expects 1 argument, got 2 in TEST(1, 2), at <: input:0:0>",
		},
		{
			name:       "boolean guard with correct type",
			typeGuards: []typeGuard{ast.IsBoolean},
			values: []ast.Expression{
				ast.BooleanExpression{Value: false},
			},
			expectsError: "",
		},
		{
			name:       "boolean guard with wrong type",
			typeGuards: []typeGuard{ast.IsBoolean},
			values: []ast.Expression{
				ast.IntExpression{Value: 1},
			},
			expectsError: "TEST() received invalid argument 1 in TEST(1), at <: input:0:0>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guard := MakeExactTypesGuard("TEST()", tt.typeGuards...)
			call.Arguments = tt.values
			err := guard(call, tt.values...)

			if tt.expectsError == "" {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error %q, got nil", tt.expectsError)
				} else if err.Error() != tt.expectsError {
					t.Errorf("Expected error %q, got %q", tt.expectsError, err.Error())
				}
			}
		})
	}
}
