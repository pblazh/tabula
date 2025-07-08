package evaluator

import (
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

func TestOperationErrors(t *testing.T) {
	testcases := []struct {
		name          string
		input         string
		expectedError string
	}{
		// Division by zero errors
		{
			name:          "int division by zero",
			input:         "10 / 0",
			expectedError: "division by zero at <DIV:/ at test:1:4>",
		},
		{
			name:          "float division by zero",
			input:         "10.5 / 0.0",
			expectedError: "division by zero at <DIV:/ at test:1:6>",
		},
		{
			name:          "int division by float zero",
			input:         "10 / 0.0",
			expectedError: "division by zero at <DIV:/ at test:1:4>",
		},
		{
			name:          "float division by int zero",
			input:         "10.5 / 0",
			expectedError: "division by zero at <DIV:/ at test:1:6>",
		},
		// Type mismatch errors for arithmetic operations
		{
			name:          "bool + int",
			input:         "true + 5",
			expectedError: "operator <PLUS:+ at test:1:6> is not supported for type: boolean and integer",
		},
		{
			name:          "int + bool",
			input:         "5 + true",
			expectedError: "operator <PLUS:+ at test:1:3> is not supported for type: integer and boolean",
		},
		{
			name:          "bool - int",
			input:         "true - 5",
			expectedError: "operator <MINUS:- at test:1:6> is not supported for type: boolean and integer",
		},
		{
			name:          "bool * int",
			input:         "true * 5",
			expectedError: "operator <MULT:* at test:1:6> is not supported for type: boolean and integer",
		},
		{
			name:          "bool / int",
			input:         "true / 5",
			expectedError: "cannot divide boolean by integer",
		},
		// Type mismatch errors for comparison operations
		{
			name:          "int < bool",
			input:         "5 < true",
			expectedError: "numeric comparison <LESS:< at test:1:3> is not supported for type: integer and boolean",
		},
		{
			name:          "bool > int",
			input:         "true > 5",
			expectedError: "numeric comparison <GREATER:> at test:1:6> is not supported for type: integer and boolean",
		},
		// Prefix operation errors
		{
			name:          "negation of bool",
			input:         "-true",
			expectedError: "<MINUS:- at test:1:1> is not supported for type: boolean",
		},
		{
			name:          "logical not of int",
			input:         "!5",
			expectedError: "<NOT:! at test:1:1> is not supported for type: integer",
		},
		{
			name:          "logical not of float",
			input:         "!3.14",
			expectedError: "<NOT:! at test:1:1> is not supported for type: float",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expr := testutil.ParseExpression(t, tc.input)

			result, err := EvaluateExpression(expr, make(map[string]string))
			if err == nil {
				t.Errorf("Expected error but got result: %s", result.String())
				return
			}

			if err.Error() != tc.expectedError {
				t.Errorf("Expected error %q, got %q", tc.expectedError, err.Error())
			}
		})
	}
}
