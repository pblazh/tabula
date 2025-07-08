package evaluator

import (
	"reflect"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestWriteValue(t *testing.T) {
	testcases := []struct {
		name     string
		key      string
		value    ast.Expression
		format   map[string]string
		expected map[string]string
	}{
		// Without format specification
		{
			name:     "write integer without format",
			key:      "result",
			value:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   map[string]string{},
			expected: map[string]string{"result": "42"},
		},
		{
			name:     "write float without format",
			key:      "result",
			value:    ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14"}},
			format:   map[string]string{},
			expected: map[string]string{"result": "3.14"},
		},
		{
			name:     "write string without format",
			key:      "result",
			value:    ast.StringExpression{Token: lexer.Token{Literal: "\"hello\""}},
			format:   map[string]string{},
			expected: map[string]string{"result": "\"hello\""},
		},
		{
			name:     "write boolean without format",
			key:      "result",
			value:    ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			format:   map[string]string{},
			expected: map[string]string{"result": "true"},
		},

		// With format specification
		{
			name:     "write integer with format",
			key:      "result",
			value:    ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format:   map[string]string{"result": "$%d"},
			expected: map[string]string{"result": "$42"},
		},
		{
			name:     "write float with format",
			key:      "result",
			value:    ast.FloatExpression{Value: 3.14159, Token: lexer.Token{Literal: "3.14159"}},
			format:   map[string]string{"result": "%.2f"},
			expected: map[string]string{"result": "3.14"},
		},
		{
			name:     "write string with format",
			key:      "result",
			value:    ast.StringExpression{Token: lexer.Token{Literal: "\"world\""}},
			format:   map[string]string{"result": "Hello, %s!"},
			expected: map[string]string{"result": "Hello, world!"},
		},
		{
			name:     "write boolean with format",
			key:      "result",
			value:    ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			format:   map[string]string{"result": "Value: %t"},
			expected: map[string]string{"result": "Value: true"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			context := make(map[string]string)
			err := WriteValue(tc.key, tc.value, context, tc.format)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(context, tc.expected) {
				t.Errorf("Expected context %v to equal %v", context, tc.expected)
			}
		})
	}
}

func TestWriteValueErrors(t *testing.T) {
	testcases := []struct {
		name   string
		value  ast.Expression
		format string
		error  string
	}{
		{
			name:   "invalid format with no placeholder",
			value:  ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format: "",
			error:  "invalid format \"\": no scanf placeholder found",
		},
		{
			name:   "invalid format with multiple placeholders",
			value:  ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			format: "%d %d",
			error:  "invalid format \"%d %d\": multiple scanf placeholders found (2), expected exactly one",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := WriteValue("key", tc.value, make(map[string]string), map[string]string{"key": tc.format})

			if err == nil {
				t.Errorf("Expected error but got none")
				return
			}
			if err.Error() != tc.error {
				t.Errorf("Expected error containing %q, got %q", tc.error, err.Error())
			}
		})
	}
}
