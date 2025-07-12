package evaluator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
	"github.com/pblazh/csvss/internal/parser"
)

func TestEvaluate(t *testing.T) {
	testcases := []struct {
		name    string
		program string
		input   [][]string
		output  [][]string
	}{
		// {
		// 	name:    "simple let statement",
		// 	program: `let B2 = 42;`,
		// 	input: [][]string{
		// 		{"1", "2"},
		// 		{"3", "4"},
		// 	},
		// 	output: [][]string{
		// 		{"1", "2"},
		// 		{"3", "42"},
		// 	},
		// },
		{
			name: "multiple let statements",
			program: `
				let A1 = 10;
				let A2 = 20;
				let B2 = A1 + A2;
			`,
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			output: [][]string{
				{"10", "2"},
				{"20", "30"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Parse the input
			lex := lexer.New(strings.NewReader(tc.program), tc.name)
			p := parser.New(lex)
			program, _, err := p.Parse()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			// Evaluate the program
			result, err := Evaluate(program, tc.input)
			if err != nil {
				t.Errorf("Unexpected evaluation error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Expected %v to equal %v", result, tc.output)
				return
			}
		})
	}
}

