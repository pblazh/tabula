package core

import (
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func TestCall(t *testing.T) {
	t.Skip("Skipping CALL function tests - external command execution is environment-dependent")
	tests := []struct {
		name     string
		args     []ast.Expression
		expected string
		hasError bool
	}{
		{
			name: "echo with single arg",
			args: []ast.Expression{
				ast.StringExpression{Value: "echo"},
				ast.StringExpression{Value: "Hello World"},
			},
			expected: "Hello World",
			hasError: false,
		},
		{
			name: "echo with multiple args",
			args: []ast.Expression{
				ast.StringExpression{Value: "echo"},
				ast.StringExpression{Value: "Hello"},
				ast.StringExpression{Value: "World"},
			},
			expected: "Hello World",
			hasError: false,
		},
		{
			name: "nonexistent command",
			args: []ast.Expression{
				ast.StringExpression{Value: "nonexistentcommand"},
			},
			expected: "",
			hasError: true,
		},
		{
			name: "empty command name",
			args: []ast.Expression{
				ast.StringExpression{Value: ""},
			},
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a dummy CallExpression
			token := lexer.Token{Type: lexer.IDENT, Literal: "CALL"}
			call := ast.CallExpression{Token: token}

			result, err := Call("CALL(command, string...)", call, tt.args...)

			if tt.hasError && err == nil {
				t.Errorf("Expected error but got none")
				return
			}

			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.hasError {
				// For error cases, just check that we got an error
				return
			}

			// Check result type
			if result == nil {
				t.Errorf("Expected result but got nil")
				return
			}

			strResult, ok := result.(ast.StringExpression)
			if !ok {
				t.Errorf("Expected StringExpression but got %T", result)
				return
			}

			if strResult.Value != tt.expected {
				t.Errorf("Expected %q but got %q", tt.expected, strResult.Value)
			}
		})
	}
}

