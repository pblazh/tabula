package evaluator

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

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
