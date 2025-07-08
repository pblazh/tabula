package evaluator

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestReadValue(t *testing.T) {
	testcases := []struct {
		name           string
		key            string
		context        map[string]string
		format         map[string]string
		expectedResult ast.Expression
		expectError    bool
		errorContains  string
	}{
		// Error cases
		{
			name:           "key not found in context",
			key:            "missing",
			context:        map[string]string{},
			format:         map[string]string{},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "missing not found in context",
		},

		// Format specification cases
		{
			name:           "format specification with integer",
			key:            "value",
			context:        map[string]string{"value": "$42"},
			format:         map[string]string{"value": "$%d"},
			expectedResult: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			expectError:    false,
		},
		{
			name:           "format specification with float",
			key:            "value",
			context:        map[string]string{"value": "3.14kg"},
			format:         map[string]string{"value": "%fkg"},
			expectedResult: ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14"}},
			expectError:    false,
		},
		{
			name:           "format specification with string",
			key:            "value",
			context:        map[string]string{"value": "hello"},
			format:         map[string]string{"value": "%s"},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"hello\""}},
			expectError:    false,
		},
		{
			name:           "format specification with boolean true",
			key:            "value",
			context:        map[string]string{"value": "true"},
			format:         map[string]string{"value": "%t"},
			expectedResult: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			expectError:    false,
		},
		{
			name:           "format specification with boolean false",
			key:            "value",
			context:        map[string]string{"value": "false"},
			format:         map[string]string{"value": "%t"},
			expectedResult: ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}},
			expectError:    false,
		},
		{
			name:           "invalid format with integer",
			key:            "value",
			context:        map[string]string{"value": "not_a_number"},
			format:         map[string]string{"value": "%d"},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "failed to parse",
		},
		{
			name:           "invalid format with boolean",
			key:            "value",
			context:        map[string]string{"value": ""},
			format:         map[string]string{"value": "%t"},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "failed to parse",
		},
		{
			name:           "unsupported format should error",
			key:            "value",
			context:        map[string]string{"value": "hello"},
			format:         map[string]string{"value": "%v"},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "failed to parse",
		},
		{
			name:           "format with no placeholders should error",
			key:            "value",
			context:        map[string]string{"value": "hello"},
			format:         map[string]string{"value": "no placeholder here"},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "no scanf placeholder found",
		},
		{
			name:           "format with multiple placeholders should error",
			key:            "value",
			context:        map[string]string{"value": "hello world"},
			format:         map[string]string{"value": "%s %s"},
			expectedResult: nil,
			expectError:    true,
			errorContains:  "multiple scanf placeholders found",
		},
		{
			name:           "format validation passes for valid single placeholder",
			key:            "value",
			context:        map[string]string{"value": "$42"},
			format:         map[string]string{"value": "$%d"},
			expectedResult: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			expectError:    false,
		},

		// Default parsing - quoted strings
		{
			name:           "quoted string without format",
			key:            "value",
			context:        map[string]string{"value": "\"hello world\""},
			format:         map[string]string{},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"hello world\""}},
			expectError:    false,
		},
		{
			name:           "empty quoted string",
			key:            "value",
			context:        map[string]string{"value": "\"\""},
			format:         map[string]string{},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"\""}},
			expectError:    false,
		},

		// Default parsing - numbers
		{
			name:           "positive integer without format",
			key:            "value",
			context:        map[string]string{"value": "42"},
			format:         map[string]string{},
			expectedResult: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			expectError:    false,
		},
		{
			name:           "negative integer without format",
			key:            "value",
			context:        map[string]string{"value": "-42"},
			format:         map[string]string{},
			expectedResult: ast.IntExpression{Value: -42, Token: lexer.Token{Literal: "-42"}},
			expectError:    false,
		},
		{
			name:           "positive float without format",
			key:            "value",
			context:        map[string]string{"value": "3.14"},
			format:         map[string]string{},
			expectedResult: ast.FloatExpression{Value: 3.14, Token: lexer.Token{Literal: "3.14"}},
			expectError:    false,
		},
		{
			name:           "negative float without format",
			key:            "value",
			context:        map[string]string{"value": "-3.14"},
			format:         map[string]string{},
			expectedResult: ast.FloatExpression{Value: -3.14, Token: lexer.Token{Literal: "-3.14"}},
			expectError:    false,
		},
		{
			name:           "number with plus sign",
			key:            "value",
			context:        map[string]string{"value": "+42"},
			format:         map[string]string{},
			expectedResult: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "+42"}},
			expectError:    false,
		},

		// Default parsing - boolean values
		{
			name:           "boolean true without format",
			key:            "value",
			context:        map[string]string{"value": "true"},
			format:         map[string]string{},
			expectedResult: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			expectError:    false,
		},
		{
			name:           "boolean false without format",
			key:            "value",
			context:        map[string]string{"value": "false"},
			format:         map[string]string{},
			expectedResult: ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}},
			expectError:    false,
		},

		// Default parsing - fallback to string
		{
			name:           "unquoted string without format returns string",
			key:            "value",
			context:        map[string]string{"value": "hello"},
			format:         map[string]string{},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"hello\""}},
			expectError:    false,
		},
		{
			name:           "invalid number returns string",
			key:            "value",
			context:        map[string]string{"value": "abc123"},
			format:         map[string]string{},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"abc123\""}},
			expectError:    false,
		},
		{
			name:           "empty string returns string",
			key:            "value",
			context:        map[string]string{"value": ""},
			format:         map[string]string{},
			expectedResult: ast.StringExpression{Token: lexer.Token{Literal: "\"\""}},
			expectError:    false,
		},

		// Whitespace trimming tests
		{
			name:           "integer with whitespace",
			key:            "value",
			context:        map[string]string{"value": "  42  "},
			format:         map[string]string{},
			expectedResult: ast.IntExpression{Value: 42, Token: lexer.Token{Literal: "42"}},
			expectError:    false,
		},
		{
			name:           "boolean with whitespace",
			key:            "value",
			context:        map[string]string{"value": "  true  "},
			format:         map[string]string{},
			expectedResult: ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}},
			expectError:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ReadValue(tc.key, tc.context, tc.format)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tc.errorContains != "" && !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("Expected error containing %q, got %q", tc.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Compare AST expressions properly
			if !compareExpressions(result, tc.expectedResult) {
				t.Errorf("Expected result %v (%T), got %v (%T)", tc.expectedResult, tc.expectedResult, result, result)
			}
		})
	}
}

func compareExpressions(a, b ast.Expression) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	switch aExpr := a.(type) {
	case ast.IntExpression:
		if bExpr, ok := b.(ast.IntExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.FloatExpression:
		if bExpr, ok := b.(ast.FloatExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.StringExpression:
		if bExpr, ok := b.(ast.StringExpression); ok {
			return aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.BooleanExpression:
		if bExpr, ok := b.(ast.BooleanExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	}
	return false
}
