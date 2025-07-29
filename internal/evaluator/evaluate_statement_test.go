package evaluator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

func TestEvaluateStatement(t *testing.T) {
	testcases := []struct {
		name            string
		statement       string
		input           [][]string
		context         map[string]string
		expectedContext map[string]string
		expectedFormat  map[string]string
		expectedInput   [][]string
	}{
		{
			name:            "simple let statement with arithmetic",
			statement:       "let result = 5 + 3;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "8"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "multi let statement with arithmetic",
			statement:       "let result = 5 + 3; let another = 9;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "8", "another": "9"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "let statement override arithmetic",
			statement:       "let result = 5 + 3; let result = 9;",
			context:         map[string]string{},
			expectedContext: map[string]string{"result": "9"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "let statement reading from context",
			statement:       "let another = 5 + result;",
			context:         map[string]string{"result": "9"},
			expectedContext: map[string]string{"result": "9", "another": "14"},
			expectedFormat:  map[string]string{},
		},
		{
			name:            "simple fmt statement with string",
			statement:       "fmt result = \"hello\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "hello"},
		},
		{
			name:            "multi fmt statement with strings",
			statement:       "fmt result = \"hello\"; fmt another = \"world\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "hello", "another": "world"},
		},
		{
			name:            "fmt statement override strings",
			statement:       "fmt result = \"hello\"; fmt result = \"world\";",
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{"result": "world"},
		},
		// Data update
		{
			name:            "update string value",
			statement:       "let A1 = \"hello\";",
			input:           [][]string{{"world"}},
			context:         map[string]string{},
			expectedContext: map[string]string{},
			expectedFormat:  map[string]string{},
			expectedInput:   [][]string{{"hello"}},
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
				err = EvaluateStatement(statement, tc.context, tc.input, format)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
			}

			if !reflect.DeepEqual(tc.context, tc.expectedContext) {
				t.Errorf("Expected context %q, got %q", tc.expectedContext, tc.context)
			}

			if !reflect.DeepEqual(format, tc.expectedFormat) {
				t.Errorf("Expected format %q, got %q", tc.expectedFormat, format)
			}

			if !reflect.DeepEqual(tc.input, tc.expectedInput) {
				t.Errorf("Expected format %q, got %q", tc.expectedInput, tc.input)
			}
		})
	}
}

func TestEvaluateStatementErrors(t *testing.T) {
	testcases := []struct {
		name          string
		statement     string
		expectedError string
	}{
		{
			name:          "fmt statement with integer should error",
			statement:     "fmt result = 5;",
			expectedError: "fmt <IDENT:result test:1:5> accepts only strings, but got <int 5>",
		},
		{
			name:          "fmt statement with float should error",
			statement:     "fmt result = 5.5;",
			expectedError: "fmt <IDENT:result test:1:5> accepts only strings, but got <float 5.50>",
		},
		{
			name:          "fmt statement with boolean should error",
			statement:     "fmt result = true;",
			expectedError: "fmt <IDENT:result test:1:5> accepts only strings, but got <bool true>",
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

			var input [][]string
			format := make(map[string]string)
			context := make(map[string]string)
			for _, statement := range program {
				err = EvaluateStatement(statement, context, input, format)
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
