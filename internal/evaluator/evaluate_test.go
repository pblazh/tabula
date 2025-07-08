package evaluator

import (
	"testing"

	"github.com/pblazh/csvss/internal/testutil"
)

func TestEvaluateStatement(t *testing.T) {
	testcases := []struct {
		name      string
		statement string
		context   map[string]string
		expected  map[string]string
	}{
		{
			name:      "simple let statement with arithmetic",
			statement: "let result = 5 + 3;",
			context:   map[string]string{},
			expected:  map[string]string{"result": "<int 8>"},
		},
		{
			name:      "multi let statement with arithmetic",
			statement: "let result = 5 + 3; let another = 9;",
			context:   map[string]string{},
			expected:  map[string]string{"result": "<int 8>", "another": "<int 9>"},
		},
		{
			name:      "let statement override arithmetic",
			statement: "let result = 5 + 3; let result = 9;",
			context:   map[string]string{},
			expected:  map[string]string{"result": "<int 9>"},
		},
		{
			name:      "let statement reading from context",
			statement: "let another = 5 + result;",
			context:   map[string]string{"result": "9"},
			expected:  map[string]string{"result": "9", "another": "<int 14>"},
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

			for _, statement := range program {
				err = EvaluateStatement(statement, tc.context)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
			}

			// Check that context matches expected
			if len(tc.context) != len(tc.expected) {
				t.Errorf("Expected context length %d, got %d", len(tc.expected), len(tc.context))
				return
			}

			for key, expectedValue := range tc.expected {
				if actualValue, exists := tc.context[key]; !exists {
					t.Errorf("Expected key %q in context", key)
				} else if actualValue != expectedValue {
					t.Errorf("Expected %q = %q, got %q", key, expectedValue, actualValue)
				}
			}
		})
	}
}
