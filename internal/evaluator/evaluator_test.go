package evaluator

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/lexer"
	"github.com/pblazh/tabula/internal/parser"
)

func TestEvaluate(t *testing.T) {
	testcases := []struct {
		name    string
		program string
		input   [][]string
		output  [][]string
	}{
		{
			name:    "let statement",
			program: `let B2 = 42;`,
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			output: [][]string{
				{"1", "2"},
				{"3", "42"},
			},
		},
		{
			name: "fmt statement",
			program: `
				fmt B1 = "%dkg";
				let B1 = 42;
				fmt B2 = "%d pounds";
				let B2 = B1 * 2;
			`,
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			output: [][]string{
				{"1", "42kg"},
				{"3", "84 pounds"},
			},
		},
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
		{
			name: "multiple let statements with variables",
			program: `
				let x = 2;
				let A1 = 10;
				let A2 = 20;
				let A2 = A2 * 2;
				let B2 = (A1 + A2) * x;
			`,
			input: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			output: [][]string{
				{"10", "2"},
				{"40", "100"},
			},
		},
		{
			name: "REL function",
			program: `
			let C1:C2 = SUM(REF(REL(-1,0)), REF(REL(-2,0)));
			`,
			input: [][]string{
				{"1", "2", "0"},
				{"3", "4", "0"},
			},
			output: [][]string{
				{"1", "2", "3"},
				{"3", "4", "7"},
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
