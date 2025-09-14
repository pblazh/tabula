package evaluator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/testutil"
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
			statement:     "fmt result = A1;",
			expectedError: `fmt accepts only strings, got 5 at test:1:5`,
		},
		{
			name:          "fmt statement with float should error",
			statement:     "fmt result = B1;",
			expectedError: `fmt accepts only strings, got 5.50 at test:1:5`,
		},
		{
			name:          "fmt statement with boolean should error",
			statement:     "fmt result = C1;",
			expectedError: `fmt accepts only strings, got true at test:1:5`,
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

			input := [][]string{{"5", "5.5", "true"}}
			format := make(map[string]string)
			context := make(map[string]string)
			for _, statement := range program {
				err = EvaluateStatement(statement, context, input, format)
				if err == nil {
					t.Errorf("Expected error, got none")
					return
				}

				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error containing %q, got %q", tc.expectedError, err.Error())
				}
			}
		})
	}
}

func TestEvaluateREL(t *testing.T) {
	testcases := []struct {
		name      string
		statement string
		input     [][]string
		output    [][]string
		error     string
	}{
		// Success cases
		{
			name:      "REL function basic usage",
			statement: "let A1 = REL(1, 0);",
			input: [][]string{
				{"10", "20", "30"},
				{"40", "50", "60"},
			},
			output: [][]string{
				{"20", "20", "30"},
				{"40", "50", "60"},
			},
		},
		{
			name:      "REL function with negative offsets",
			statement: "let C2 = REL(-1, -1);",
			input: [][]string{
				{"10", "20", "30"},
				{"40", "50", "60"},
			},
			output: [][]string{
				{"10", "20", "30"},
				{"40", "50", "20"}, // C2 gets value from B1
			},
		},
		{
			name:      "REL function in arithmetic expression",
			statement: "let A3 = REL(0, -1) + REL(1, 0);",
			input: [][]string{
				{"0", "0"},
				{"1", "0"},
				{"0", "2"},
			},
			output: [][]string{
				{"0", "0"},
				{"1", "0"},
				{"3", "2"},
			},
		},
		{
			name:      "REL function with same cell reference",
			statement: "let A1 = REL(0, 0);",
			input: [][]string{
				{"1", "0"},
				{"0", "0"},
			},
			output: [][]string{
				{"1", "0"},
				{"0", "0"},
			},
		},
		{
			name:      "REL function with evaluated cell reference",
			statement: "let A1 = REL(3-2, 9-8);",
			input: [][]string{
				{"1", "0"},
				{"0", "2"},
			},
			output: [][]string{
				{"2", "0"},
				{"0", "2"},
			},
		},
		{
			name:      "REL function with nested evaluated cell reference",
			statement: "let A1 = REL(ADD(1,1), ADD(1,1));",
			input: [][]string{
				{"1", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "3"},
			},
			output: [][]string{
				{"3", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "3"},
			},
		},
		{
			name:      "deeply nested REL functions",
			statement: "let B2 = IF(REL(-1, 0) > 10, SUM(REL(0, -1), REL(1, 1)), ABS(REL(-1, -1)));",
			input: [][]string{
				{"5", "8", "12"},
				{"20", "0", "30"},
				{"25", "35", "40"},
			},
			output: [][]string{
				{"5", "8", "12"},
				{"20", "48", "30"}, // B2: IF(A2 > 10, SUM(B1, C3), ABS(A1)) = IF(20 > 10, SUM(8, 40), ABS(5)) = SUM(8, 40) = 48
				{"25", "35", "40"},
			},
		},
		// Error cases
		{
			name:      "REL out of bounds row",
			statement: "let A1 = REL(0, 5);",
			input:     [][]string{{"10", "20"}},
			error:     "REL(0, 5) is outof bounds",
		},
		{
			name:      "REL out of bounds column",
			statement: "let A1 = REL(5, 0);",
			input:     [][]string{{"10", "20"}},
			error:     "REL(5, 0) is outof bounds",
		},
		{
			name:      "REL negative coordinates",
			statement: "let A1 = REL(-1, 0);",
			input:     [][]string{{"10", "20"}},
			error:     "REL(-1, 0) is outof bounds",
		},
		{
			name:      "REL with non-integer offset",
			statement: "let A1 = REL(\"hello\", 0);",
			input:     [][]string{{"10", "20"}},
			error:     "string is not supported by REL(\"hello\", 0)",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			program, err := testutil.ParseProgram(tc.statement)
			if err != nil {
				t.Fatalf("Failed to parse program: %v", err)
			}

			if len(program) == 0 {
				t.Fatalf("No statements parsed")
			}

			format := make(map[string]string)
			context := make(map[string]string)

			for _, statement := range program {
				err = EvaluateStatement(statement, context, tc.input, format)

				// Error case
				if tc.error != "" {
					if err == nil {
						t.Errorf("Expected error, got none")
						return
					}
					if !strings.Contains(err.Error(), tc.error) {
						t.Errorf("Expected error containing %q, got %q", tc.error, err.Error())
					}
					return
				}

				// Success case
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				// Check output for success cases
				if tc.error == "" && !reflect.DeepEqual(tc.input, tc.output) {
					t.Errorf("Expected output %v, got %v", tc.output, tc.input)
				}
			}
		})
	}
}
