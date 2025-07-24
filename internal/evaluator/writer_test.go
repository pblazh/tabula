package evaluator

import (
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestWriteValue(t *testing.T) {
	testcases := []struct {
		name     string
		format   string
		input    ast.Expression
		expected string
		error    string
	}{
		// Without format specification
		{
			name:     "write integer without format",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "",
			expected: "42",
		},
		{
			name:     "write float without format",
			input:    ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14"}},
			format:   "",
			expected: "3.14",
		},
		{
			name:     "write string without format",
			input:    ast.StringExpression{Value: "hello", Token: lexer.Token{Literal: "\"hello\""}},
			format:   "",
			expected: "hello",
		},
		{
			name:     "write boolean without format",
			input:    ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			format:   "",
			expected: "true",
		},

		// With format specification
		{
			name:     "write integer with format",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "$%d",
			expected: "$42",
		},
		{
			name:     "write float with format",
			input:    ast.FloatExpression{Value: 3.14159, Token: lexer.Token{Literal: "3.14159"}},
			format:   "%.2f",
			expected: "3.14",
		},
		{
			name:     "write string with format",
			input:    ast.StringExpression{Value: "world", Token: lexer.Token{Literal: "\"hello\""}},
			format:   "Hello, %s!",
			expected: "Hello, world!",
		},
		{
			name:     "write string without format",
			input:    ast.StringExpression{Value: "world", Token: lexer.Token{Literal: "\"hello\""}},
			format:   "",
			expected: "world",
		},
		{
			name:     "write boolean with format",
			input:    ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			format:   "Value: %t",
			expected: "Value: true",
		},

		// Error cases
		{
			name:     "invalid format with no placeholder",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "no placeholder here",
			expected: "",
			error:    "invalid format \"no placeholder here\": no scanf placeholder found",
		},
		{
			name:     "invalid format with multiple placeholders",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "%d %d",
			expected: "",
			error:    "invalid format \"%d %d\": multiple scanf placeholders found (2), expected exactly one",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := WriteValue(tc.input, tc.format)

			if tc.error != "" {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tc.error != "" && err.Error() != tc.error {
					t.Errorf("Expected error  %q, got %q", tc.error, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tc.expected {
				t.Errorf("Expected result %q, got %q", tc.expected, result)
			}
		})
	}
}
