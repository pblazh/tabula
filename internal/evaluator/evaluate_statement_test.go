package evaluator

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

func TestEvaluateStatement(t *testing.T) {
	testcases := []struct {
		name            string
		statement       string
		context         map[string]string
		expectedContext map[string]string
		expectedFormat  map[string]string
	}{
		{
			name:            "simple let statement with arithmetic",
			statement:       "let result = 5 + 3;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "<int 8>"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "multi let statement with arithmetic",
			statement:       "let result = 5 + 3; let another = 9;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "<int 8>", "another": "<int 9>"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "let statement override arithmetic",
			statement:       "let result = 5 + 3; let result = 9;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "<int 9>"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "let statement reading from context",
			statement:       "let another = 5 + result;",
			context:         map[string]string{"result": "9"},
			expectedContext: map[string]string{"result": "9", "another": "<int 14>"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "simple fmt statement with string",
			statement:       "fmt result = \"hello\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "<str \"hello\">"},
		},
		{
			name:            "multi fmt statement with strings",
			statement:       "fmt result = \"hello\"; fmt another = \"world\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "<str \"hello\">", "another": "<str \"world\">"},
		},
		{
			name:            "fmt statement override strings",
			statement:       "fmt result = \"hello\"; fmt result = \"world\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "<str \"world\">"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			program, err := testutil.ParseProgram(tc.statement)
			if err != nil {
				t.Fatalf("Failed to parse statement: %v", err)
			}

			if len(program) == 0 {
				t.Fatalf("No statements parsed")
			}
			format := make(map[string]string)
			for _, statement := range program {
				err = EvaluateStatement(statement, tc.context, format)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
			}

			// Check that context matches expected
			if len(tc.context) != len(tc.expectedContext) {
				t.Errorf("Expected context length %d, got %d", len(tc.expectedContext), len(tc.context))
				return
			}

			for key, expectedValue := range tc.expectedContext {
				if actualValue, exists := tc.context[key]; !exists {
					t.Errorf("Expected key %q in context", key)
				} else if actualValue != expectedValue {
					t.Errorf("Expected context %q = %q, got %q", key, expectedValue, actualValue)
				}
			}

			// Check that format matches expected
			if len(format) != len(tc.expectedFormat) {
				t.Errorf("Expected format length %d, got %d", len(tc.expectedFormat), len(format))
				return
			}

			for key, expectedValue := range tc.expectedFormat {
				if actualValue, exists := format[key]; !exists {
					t.Errorf("Expected key %q in format", key)
				} else if actualValue != expectedValue {
					t.Errorf("Expected format %q = %q, got %q", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestEvaluateStatementErrors(t *testing.T) {
	testcases := []struct {
		name          string
		statement     string
		context       map[string]string
		expectedError string
	}{
		{
			name:          "fmt statement with integer should error",
			statement:     "fmt result = 5;",
			context:       map[string]string{},
			expectedError: "fmt <IDENT:result at test:1:5> accepts only strings, but got <int 5>",
		},
		{
			name:          "fmt statement with float should error",
			statement:     "fmt result = 5.5;",
			context:       map[string]string{},
			expectedError: "fmt <IDENT:result at test:1:5> accepts only strings, but got <float 5.50>",
		},
		{
			name:          "fmt statement with boolean should error",
			statement:     "fmt result = true;",
			context:       map[string]string{},
			expectedError: "fmt <IDENT:result at test:1:5> accepts only strings, but got <bool true>",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			program, err := testutil.ParseProgram(tc.statement)
			if err != nil {
				t.Fatalf("Failed to parse statement: %v", err)
			}

			if len(program) == 0 {
				t.Fatalf("No statements parsed")
			}

			format := make(map[string]string)
			for _, statement := range program {
				err = EvaluateStatement(statement, tc.context, format)
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}

				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error containing %q, got %q", tc.expectedError, err.Error())
				}
			}
		})
	}
}
