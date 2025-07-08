package evaluator

import (
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

func TestInfixExpressionEvaluate(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		// Addition
		{
			name:     "int + int",
			input:    "5 + 3",
			expected: "<int 8>",
		},
		{
			name:     "int + float",
			input:    "5 + 3.5",
			expected: "<float 8.50>",
		},
		{
			name:     "float + int",
			input:    "5.5 + 3",
			expected: "<float 8.50>",
		},
		{
			name:     "float + float",
			input:    "5.5 + 3.2",
			expected: "<float 8.70>",
		},
		// Subtraction
		{
			name:     "int - int",
			input:    "10 - 3",
			expected: "<int 7>",
		},
		{
			name:     "int - float",
			input:    "10 - 3.5",
			expected: "<float 6.50>",
		},
		{
			name:     "float - int",
			input:    "10.5 - 3",
			expected: "<float 7.50>",
		},
		{
			name:     "float - float",
			input:    "10.5 - 3.2",
			expected: "<float 7.30>",
		},
		// Multiplication
		{
			name:     "int * int",
			input:    "5 * 3",
			expected: "<int 15>",
		},
		{
			name:     "int * float",
			input:    "5 * 2.5",
			expected: "<float 12.50>",
		},
		{
			name:     "float * int",
			input:    "4.5 * 2",
			expected: "<float 9.00>",
		},
		{
			name:     "float * float",
			input:    "2.5 * 3.0",
			expected: "<float 7.50>",
		},
		// Division
		{
			name:     "int / int",
			input:    "10 / 2",
			expected: "<int 5>",
		},
		{
			name:     "int / float",
			input:    "10 / 2.5",
			expected: "<float 4.00>",
		},
		{
			name:     "float / int",
			input:    "10.5 / 2",
			expected: "<float 5.25>",
		},
		{
			name:     "float / float",
			input:    "10.5 / 2.5",
			expected: "<float 4.20>",
		},
		// Equality
		{
			name:     "int == int (true)",
			input:    "5 == 5",
			expected: "<bool true>",
		},
		{
			name:     "int == int (false)",
			input:    "5 == 3",
			expected: "<bool false>",
		},
		{
			name:     "int == float (true)",
			input:    "5 == 5.0",
			expected: "<bool true>",
		},
		{
			name:     "int == float (false)",
			input:    "5 == 5.01",
			expected: "<bool false>",
		},
		{
			name:     "float == int (true)",
			input:    "5.0 == 5",
			expected: "<bool true>",
		},
		{
			name:     "float == int (false)",
			input:    "5.01 == 5",
			expected: "<bool false>",
		},
		{
			name:     "float == float (true)",
			input:    "3.14 == 3.14",
			expected: "<bool true>",
		},
		{
			name:     "float == float (false)",
			input:    "3.14 == 2.71",
			expected: "<bool false>",
		},
		{
			name:     "bool == bool (true)",
			input:    "true == true",
			expected: "<bool true>",
		},
		{
			name:     "bool == bool (false)",
			input:    "true == false",
			expected: "<bool false>",
		},
		// Inequality
		{
			name:     "int != int (true)",
			input:    "5 != 3",
			expected: "<bool true>",
		},
		{
			name:     "int != int (false)",
			input:    "5 != 5",
			expected: "<bool false>",
		},
		{
			name:     "int != float (true)",
			input:    "5 != 3.5",
			expected: "<bool true>",
		},
		{
			name:     "int != float (false)",
			input:    "5 != 5.0",
			expected: "<bool false>",
		},
		{
			name:     "float != int (true)",
			input:    "3.5 != 5",
			expected: "<bool true>",
		},
		{
			name:     "float != int (false)",
			input:    "5.0 != 5",
			expected: "<bool false>",
		},
		{
			name:     "float != float (true)",
			input:    "3.14 != 2.71",
			expected: "<bool true>",
		},
		{
			name:     "float != float (false)",
			input:    "3.14 != 3.14",
			expected: "<bool false>",
		},
		{
			name:     "bool != bool (true)",
			input:    "true != false",
			expected: "<bool true>",
		},
		{
			name:     "bool != bool (false)",
			input:    "true != true",
			expected: "<bool false>",
		},
		// Less Than Comparison
		{
			name:     "int < int (true)",
			input:    "3 < 5",
			expected: "<bool true>",
		},
		{
			name:     "int < int (false)",
			input:    "5 < 3",
			expected: "<bool false>",
		},
		{
			name:     "int < int (equal)",
			input:    "5 < 5",
			expected: "<bool false>",
		},
		{
			name:     "int < float (true)",
			input:    "3 < 5.5",
			expected: "<bool true>",
		},
		{
			name:     "int < float (false)",
			input:    "5 < 3.5",
			expected: "<bool false>",
		},
		{
			name:     "float < int (true)",
			input:    "3.5 < 5",
			expected: "<bool true>",
		},
		{
			name:     "float < int (false)",
			input:    "5.5 < 3",
			expected: "<bool false>",
		},
		{
			name:     "float < float (true)",
			input:    "3.2 < 5.7",
			expected: "<bool true>",
		},
		{
			name:     "float < float (false)",
			input:    "5.7 < 3.2",
			expected: "<bool false>",
		},
		// Greater Than Comparison
		{
			name:     "int > int (true)",
			input:    "5 > 3",
			expected: "<bool true>",
		},
		{
			name:     "int > int (false)",
			input:    "3 > 5",
			expected: "<bool false>",
		},
		{
			name:     "int > int (equal)",
			input:    "5 > 5",
			expected: "<bool false>",
		},
		{
			name:     "int > float (true)",
			input:    "5 > 3.5",
			expected: "<bool true>",
		},
		{
			name:     "int > float (false)",
			input:    "3 > 5.5",
			expected: "<bool false>",
		},
		{
			name:     "float > int (true)",
			input:    "5.5 > 3",
			expected: "<bool true>",
		},
		{
			name:     "float > int (false)",
			input:    "3.5 > 5",
			expected: "<bool false>",
		},
		{
			name:     "float > float (true)",
			input:    "5.7 > 3.2",
			expected: "<bool true>",
		},
		{
			name:     "float > float (false)",
			input:    "3.2 > 5.7",
			expected: "<bool false>",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			result, err := EvaluateExpression(expr, make(map[string]string), make(map[string]string))
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}

func TestPrefixExpressionEvaluate(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{
			name:     "negation of int",
			input:    "-5",
			expected: "<int -5>",
			hasError: false,
		},
		{
			name:     "negation of float",
			input:    "-3.5",
			expected: "<float -3.50>",
			hasError: false,
		},
		{
			name:     "negation of negative int",
			input:    "-(-5)",
			expected: "<int 5>",
			hasError: false,
		},
		{
			name:     "logical not of true",
			input:    "!true",
			expected: "<bool false>",
			hasError: false,
		},
		{
			name:     "logical not of false",
			input:    "!false",
			expected: "<bool true>",
			hasError: false,
		},
		{
			name:     "addition and multiplication precedence",
			input:    "2 + 3 * 4",
			expected: "<int 14>",
		},
		{
			name:     "multiplication and addition precedence",
			input:    "3 * 4 + 2",
			expected: "<int 14>",
		},
		{
			name:     "parentheses override precedence",
			input:    "(2 + 3) * 4",
			expected: "<int 20>",
		},
		{
			name:     "complex arithmetic",
			input:    "10 - 2 * 3 + 1",
			expected: "<int 5>",
		},
		{
			name:     "division and multiplication",
			input:    "12 / 3 * 2",
			expected: "<int 8>",
		},
		{
			name:     "comparison with arithmetic",
			input:    "5 + 3 > 7",
			expected: "<bool true>",
		},
		{
			name:     "equality with arithmetic",
			input:    "2 * 3 == 6",
			expected: "<bool true>",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			result, err := EvaluateExpression(expr, make(map[string]string), make(map[string]string))
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.String() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result.String())
			}
		})
	}
}
