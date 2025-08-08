package evaluator

import (
	"reflect"
	"testing"
	"time"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func getTime(value string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", value)
	return t
}

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
			name:     "format specification with precision",
			input:    "3.14кг",
			format:   "%.2fкг",
			expected: ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14кг"}},
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
			error:    "parsing time \"hello\" as \"%v\": cannot parse \"hello\" as \"%v\"",
		},
		{
			name:     "format with no placeholders should error",
			input:    "hello",
			format:   "no placeholder here",
			expected: nil,
			error:    "parsing time \"hello\" as \"no placeholder here\": cannot parse \"hello\" as \"no placeholder here\"",
		},
		{
			name:     "format with multiple placeholders should error",
			input:    "hello",
			format:   "%s %s",
			expected: nil,
			error:    "failed to parse \"hello\" with format \"%s %s\":too few operands for format '%s'",
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
			expected: ast.StringExpression{Value: "\"hello world\"", Token: lexer.Token{Literal: "\"hello world\""}},
		},
		{
			name:     "empty quoted string",
			input:    "\"\"",
			expected: ast.StringExpression{Value: "\"\"", Token: lexer.Token{Literal: "\"\""}},
		},

		// Default parsing - numbers
		{
			name:     "positive integer without format",
			input:    "42",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
		},
		{
			name:     "negative integer without format",
			input:    "-42",
			expected: ast.IntExpression{Value: -42, Token: lexer.Token{Literal: "-42"}},
		},
		{
			name:     "positive float without format",
			input:    "42.3",
			expected: ast.FloatExpression{Value: 42.3, Token: lexer.Token{Literal: "42.3"}},
		},
		{
			name:     "negative float without format",
			input:    "-42.3",
			expected: ast.FloatExpression{Value: -42.3, Token: lexer.Token{Literal: "-42.3"}},
		},
		{
			name:     "number with plus sign",
			input:    "+42",
			expected: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "+42"}},
		},

		// Default parsing - boolean values
		{
			name:     "boolean true without format",
			input:    "true",
			expected: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
		},
		{
			name:     "boolean false without format",
			input:    "false",
			expected: ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}},
		},
		// Default parsing - date values
		{
			name:     "iso date without format",
			input:    "2023-10-01",
			expected: ast.DateExpression{Value: getTime("2023-10-01 00:00:00"), Token: lexer.Token{Literal: "2023-10-01"}},
		},
		{
			name:     "iso datetime without format",
			input:    "2023-10-01 13:41",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:00"), Token: lexer.Token{Literal: "2023-10-01 13:41"}},
		},
		{
			name:     "iso datetime sec without format",
			input:    "2023-10-01 13:41:51",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:51"), Token: lexer.Token{Literal: "2023-10-01 13:41:51"}},
		},

		{
			name:     "eu date without format",
			input:    "01.10.2023",
			expected: ast.DateExpression{Value: getTime("2023-10-01 00:00:00"), Token: lexer.Token{Literal: "01.10.2023"}},
		},
		{
			name:     "eu datetime without format",
			input:    "01.10.2023 13:41",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:00"), Token: lexer.Token{Literal: "01.10.2023 13:41"}},
		},
		{
			name:     "eu datetime sec without format",
			input:    "01.10.2023 13:41:51",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:51"), Token: lexer.Token{Literal: "01.10.2023 13:41:51"}},
		},

		{
			name:     "us date without format",
			input:    "10/01/2023",
			expected: ast.DateExpression{Value: getTime("2023-10-01 00:00:00"), Token: lexer.Token{Literal: "10/01/2023"}},
		},
		{
			name:     "us datetime without format",
			input:    "10/01/2023 13:41",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:00"), Token: lexer.Token{Literal: "10/01/2023 13:41"}},
		},
		{
			name:     "us datetime sec without format",
			input:    "10/01/2023 13:41:51",
			expected: ast.DateExpression{Value: getTime("2023-10-01 13:41:51"), Token: lexer.Token{Literal: "10/01/2023 13:41:51"}},
		},
		{
			name:     "timeonly",
			input:    "13:41:51",
			expected: ast.DateExpression{Value: getTime("0000-01-01 13:41:51"), Token: lexer.Token{Literal: "13:41:51"}},
		},
		{
			name:     "kitchen",
			input:    "03:41PM",
			expected: ast.DateExpression{Value: getTime("0000-01-01 15:41:00"), Token: lexer.Token{Literal: "03:41PM"}},
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
