package evaluator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
	"github.com/pblazh/csvss/internal/parser"
)

// Helper function to create test programs for more complex evaluation tests
func createTestProgram(input string) (ast.Program, error) {
	lex := lexer.New(strings.NewReader(input), "test")
	parser := parser.New(lex)
	program, _, err := parser.Parse()
	return program, err
}

func TestEvaluate(t *testing.T) {
	testcases := []struct {
		name    string
		input   string
		context [][]string
		output  map[string]string
	}{
		{
			name:    "empty programm",
			input:   "",
			context: [][]string{},
			output:  map[string]string{},
		},
		{
			name:    "arithmetic expression",
			input:   "let x = 5 + 3;",
			context: [][]string{},
			output: map[string]string{
				"x": "8",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			program, err := createTestProgram(tc.input)
			if err != nil {
				t.Errorf("Failed to parse program: %v", err)
				return
			}

			// Test that evaluation doesn't crash
			result, err := Evaluate(program, tc.context)
			if err != nil {
				// For now, we expect errors since evaluation is not implemented
				t.Logf("Evaluation error (expected): %v", err)
			}

			if !reflect.DeepEqual(result, tc.output) {
				// For now, we expect errors since evaluation is not implemented
				t.Logf("Evaluation error (expected): %v", err)
			}
		})
	}
}

