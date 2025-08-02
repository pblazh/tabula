package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestProgramSorting(t *testing.T) {
	testcases := []struct {
		name        string
		input       string
		output      string
		shouldError bool
	}{
		{
			name: "simple dependency chain",
			input: `let A1 = B1 + C1;
let B1 = 10;
let C1 = 20;`,
			output: "let B1 = <int 10>;let C1 = <int 20>;let A1 = (+ B1 C1);",
		},
		{
			name: "complex dependency chain",
			input: `let D1 = B1 + C1;
let B1 = A1;
let A1 = 10;
let C1 = 20;`,
			output: "let A1 = <int 10>;let B1 = A1;let C1 = <int 20>;let D1 = (+ B1 C1);",
		},
		{
			name: "no dependencies",
			input: `let A1 = 10;
let B1 = 20;
let C1 = 30;`,
			output: "let A1 = <int 10>;let B1 = <int 20>;let C1 = <int 30>;",
		},
		{
			name: "circular dependency",
			input: `let A1 = B1;
let B1 = A1;`,
			shouldError: true,
		},
		{
			name: "call expression dependencies",
			input: `let D1 = SUM(A1, B1, C1);
let B1 = 20;
let A1 = 10;
let C1 = 30;`,
			output: "let A1 = <int 10>;let B1 = <int 20>;let C1 = <int 30>;let D1 = (SUM A1 B1 C1);",
		},
		{
			name: "mixed statements with expressions",
			input: `let C1 = A1 + B1;
A1 + B1;
let A1 = 10;
let B1 = 20;
C1 * 2;`,
			output: "let A1 = <int 10>;let B1 = <int 20>;let C1 = (+ A1 B1);(+ A1 B1);(* C1 <int 2>);",
		},
		{
			name: "range expression dependencies",
			input: `let D1 = SUM(A1:C1);
let B1 = 10;
let C1 = 20;
let E1 = 30;`,
			output: "let B1 = <int 10>;let C1 = <int 20>;let D1 = (SUM A1 B1 C1);let E1 = <int 30>;",
		},
		{
			name: "prefix expression dependencies",
			input: `let B1 = -A1;
let A1 = 10;`,
			output: "let A1 = <int 10>;let B1 = (- A1);",
		},
		{
			name: "nested expression dependencies",
			input: `let D1 = (A1 + B1) * C1;
let A1 = 10;
let B1 = 20;
let C1 = 30;`,
			output: "let A1 = <int 10>;let B1 = <int 20>;let C1 = <int 30>;let D1 = (* (+ A1 B1) C1);",
		},
		{
			name: "three-level circular dependency",
			input: `let A1 = B1;
let B1 = C1;
let C1 = A1;`,
			shouldError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			program, _, err := parser.Parse()
			if err != nil {
				t.Errorf("Unexpected parsing error: %v", err)
				return
			}

			sorted, err := ast.SortProgram(program)

			if tc.shouldError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected sorting error: %v", err)
				return
			}

			literal := ""
			for _, statement := range sorted {
				literal += statement.String()
			}

			if literal != tc.output {
				t.Errorf("Expected '%s' to equal '%s'", literal, tc.output)
			}
		})
	}
}
