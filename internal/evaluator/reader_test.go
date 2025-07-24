package evaluator

import (
	"reflect"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestReadValue(t *testing.T) {
	testcases := []struct {
		name     string
		format   string
		input    string
		expected ast.Expression
		error    string
	}{
		// Format specification cases
		{
			name:     "format specification with integer",
			input:    "$42",
			format:   "$%d",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "$42"}},
		},
		{
			name:     "format specification with float",
			input:    "3.14kg",
			format:   "%fkg",
			expected: ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14kg"}},
		},
		{
			name:     "format specification with string",
			input:    "hello",
			format:   "%s",
			expected: ast.StringExpression{Value: "hello", Token: lexer.Token{Literal: "hello"}},
		},
		{
			name:     "format specification with boolean true",
			input:    "true",
			format:   "%t",
			expected: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
		},
		{
			name:     "format specification with boolean false",
			format:   "%t",
			input:    "false",
			expected: ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}},
		},
		{
			name:     "invalid format with integer",
			input:    "not a number",
			format:   "%d",
			expected: nil,
			error:    "failed to parse \"not a number\" with format \"%d\":expected integer",
		},
		{
			name:     "invalid format with boolean",
			input:    "",
			format:   "%t",
			expected: nil,
			error:    "failed to parse \"\" with format \"%t\":EOF",
		},
		{
			name:     "unsupported format should error",
			input:    "hello",
			format:   "%v",
			expected: nil,
			error:    "failed to parse \"hello\" with format \"%v\":",
		},
		{
			name:     "format with no placeholders should error",
			input:    "hello",
			format:   "no placeholder here",
			expected: nil,
			error:    "invalid format \"no placeholder here\": no scanf placeholder found",
		},
		{
			name:     "format with multiple placeholders should error",
			input:    "hello",
			format:   "%s %s",
			expected: nil,
			error:    "invalid format \"%s %s\": multiple scanf placeholders found (2), expected exactly one",
		},
		{
			name:     "format validation passes for valid single placeholder",
			input:    "$42",
			format:   "$%d",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "$42"}},
		},

		// Default parsing - quoted strings
		{
			name:     "quoted string without format",
			input:    "\"hello world\"",
			format:   "",
			expected: ast.StringExpression{Value: "\"hello world\"", Token: lexer.Token{Literal: "\"hello world\""}},
		},
		{
			name:     "empty quoted string",
			input:    "\"\"",
			format:   "",
			expected: ast.StringExpression{Value: "\"\"", Token: lexer.Token{Literal: "\"\""}},
		},

		// Default parsing - numbers
		{
			name:     "positive integer without format",
			input:    "42",
			format:   "",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
		},
		{
			name:     "negative integer without format",
			input:    "-42",
			format:   "",
			expected: ast.IntExpression{Value: -42, Token: lexer.Token{Literal: "-42"}},
		},
		{
			name:     "positive float without format",
			input:    "42.3",
			format:   "",
			expected: ast.FloatExpression{Value: 42.3, Token: lexer.Token{Literal: "42.3"}},
		},
		{
			name:     "negative float without format",
			input:    "-42.3",
			format:   "",
			expected: ast.FloatExpression{Value: -42.3, Token: lexer.Token{Literal: "-42.3"}},
		},
		{
			name:     "number with plus sign",
			input:    "+42",
			format:   "",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "+42"}},
		},

		// Default parsing - boolean values
		{
			name:     "boolean true without format",
			input:    "true",
			format:   "",
			expected: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
		},
		{
			name:     "boolean false without format",
			input:    "false",
			format:   "",
			expected: ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}},
		},

		// Default parsing - fallback to string
		{
			name:     "unquoted string without format returns string",
			input:    "hello",
			format:   "",
			expected: ast.StringExpression{Value: "hello", Token: lexer.Token{Literal: "hello"}},
		},
		{
			name:     "invalid number returns string",
			input:    "abc123",
			format:   "",
			expected: ast.StringExpression{Value: "abc123", Token: lexer.Token{Literal: "abc123"}},
		},
		{
			name:     "empty string returns string",
			input:    "",
			format:   "",
			expected: ast.StringExpression{Value: "", Token: lexer.Token{Literal: ""}},
		},

		// Whitespace trimming tests
		{
			name:     "integer with whitespace",
			input:    "\t42  ",
			format:   "",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
		},
		{
			name:     "boolean with whitespace",
			input:    "\ttrue  ",
			format:   "",
			expected: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadValue(tc.input, tc.format)

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

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v (%T), got %v (%T)", tc.expected, tc.expected, result, result)
			}
		})
	}
}
