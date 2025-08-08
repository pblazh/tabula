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
		{
			name:     "write date without format",
			input:    ast.DateExpression{Value: getTime("2025-08-12 17:51:21"), Token: lexer.Token{Literal: "2025-08-12 17:51:21"}},
			format:   "",
			expected: "2025-08-12 17:51:21",
		},

		// Date with format specification
		{
			name:     "write date with format",
			input:    ast.DateExpression{Value: getTime("2025-08-12 17:51:21"), Token: lexer.Token{Literal: "2025-08-12 17:51:21"}},
			format:   "2006/01/02",
			expected: "2025/08/12",
		},
		{
			name:     "write date with time format",
			input:    ast.DateExpression{Value: getTime("2025-08-12 17:51:21"), Token: lexer.Token{Literal: "2025-08-12 17:51:21"}},
			format:   "15:04:05",
			expected: "17:51:21",
		},
		{
			name:     "write date with custom format",
			input:    ast.DateExpression{Value: getTime("2025-08-12 17:51:21"), Token: lexer.Token{Literal: "2025-08-12 17:51:21"}},
			format:   "Jan 2, 2006 at 3:04 PM",
			expected: "Aug 12, 2025 at 5:51 PM",
		},
		{
			name:     "write date with RFC3339 format",
			input:    ast.DateExpression{Value: getTime("2025-08-12 17:51:21"), Token: lexer.Token{Literal: "2025-08-12 17:51:21"}},
			format:   "2006-01-02T15:04:05Z07:00",
			expected: "2025-08-12T17:51:21Z",
		},

		// Error cases
		{
			name:     "invalid format with no placeholder",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "no placeholder here",
			expected: "no placeholder here%!(EXTRA int=42)",
		},
		{
			name:     "invalid format with multiple placeholders",
			input:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   "%d %d",
			expected: "42 %!d(MISSING)",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := WriteValue(tc.input, tc.format)

			if tc.error != "" {
				if err == nil {
					t.Errorf("Expected error but got none. Result: %v", result)
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
