package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/lexer"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "simple assign",
			input:  "let A1 = -10;",
			output: "let A1 = -10;",
		},
		{
			name:   "range assign",
			input:  "let A1:A3 = 1;",
			output: "let A1 = 1;let A2 = 1;let A3 = 1;",
		},
		{
			name:   "multy values assign",
			input:  "let A1,A3 = 1;",
			output: "let A1 = 1;let A3 = 1;",
		},
		{
			name:   "multy range assign",
			input:  "let A1:A3, B1:b3 = 1;",
			output: "let A1 = 1;let A2 = 1;let A3 = 1;let B1 = 1;let B2 = 1;let B3 = 1;",
		},
		{
			name:   "mixed range values assign",
			input:  "let a, A1:A3, B1:b3, C2, b = 1;",
			output: "let a = 1;let A1 = 1;let A2 = 1;let A3 = 1;let B1 = 1;let B2 = 1;let B3 = 1;let C2 = 1;let b = 1;",
		},
		{
			name:   "range fmt",
			input:  "fmt A1:A3 = \"%s\";",
			output: "fmt A1 = \"%s\";fmt A2 = \"%s\";fmt A3 = \"%s\";",
		},
		{
			name:   "multy values fmt",
			input:  "fmt A1,A3 = \"%s\";",
			output: "fmt A1 = \"%s\";fmt A3 = \"%s\";",
		},
		{
			name:   "multy range fmt",
			input:  "fmt A1:A3, B1:b3 = \"%s\";",
			output: "fmt A1 = \"%s\";fmt A2 = \"%s\";fmt A3 = \"%s\";fmt B1 = \"%s\";fmt B2 = \"%s\";fmt B3 = \"%s\";",
		},
		{
			name:   "mixed range values fmt",
			input:  "fmt a, A1:A3, B1:b3, C2, b = \"%s\";",
			output: "fmt a = \"%s\";fmt A1 = \"%s\";fmt A2 = \"%s\";fmt A3 = \"%s\";fmt B1 = \"%s\";fmt B2 = \"%s\";fmt B3 = \"%s\";fmt C2 = \"%s\";fmt b = \"%s\";",
		},
		{
			name:   "identifier",
			input:  "something;",
			output: "something;",
		},
		{
			name:   "not identifier",
			input:  "!something;",
			output: "!something;",
		},
		{
			name:   "number",
			input:  "10;",
			output: "10;",
		},
		{
			name:   "number",
			input:  "-10;",
			output: "-10;",
		},
		{
			name:   "true",
			input:  "true;",
			output: "true;",
		},
		{
			name:   "false",
			input:  "false;",
			output: "false;",
		},
		{
			name:   "compare with boolen",
			input:  "3 > 5 == false;",
			output: "3 > 5 == false;",
		},
		{
			name:   "compare boolen",
			input:  "true == 3 < 5;",
			output: "true == 3 < 5;",
		},
		{
			name:   "infix",
			input:  "5 + 6 - 2;",
			output: "5 + 6 - 2;",
		},
		{
			name:   "infix reverse precedence",
			input:  "5 + 6 * 2;",
			output: "5 + 6 * 2;",
		},
		{
			name:   "infix precedence",
			input:  "5 / 6 + 2;",
			output: "5 / 6 + 2;",
		},
		{
			name:   "multiple statements",
			input:  "let A1 = 5.6;\nlet A2 = x;\n",
			output: "let A1 = 5.60;let A2 = x;",
		},
		{
			name:   "parenteces precedence",
			input:  "(5 + 6) * 2;",
			output: "5 + 6 * 2;",
		},
		{
			name:   "call expression with no arguments",
			input:  "SUM();",
			output: "SUM();",
		},
		{
			name:   "call expression with one argument",
			input:  "SUM(5);",
			output: "SUM(5);",
		},
		{
			name:   "call expression with multiple numbers",
			input:  "SUM(5, 6);",
			output: "SUM(5, 6);",
		},
		{
			name:   "call expression with multiple strings",
			input:  "SUM(\"5\", \"6\");",
			output: "SUM(\"5\", \"6\");",
		},
		{
			name:   "cell expression",
			input:  "a1;",
			output: "A1;",
		},
		{
			name:   "range expression",
			input:  "a1:B1;",
			output: "[A1, B1];",
		},
		{
			name:   "string expression",
			input:  "let a = \"hello\";",
			output: "let a = \"hello\";",
		},
		{
			name:   "multiple statements",
			input:  "let A1 = 5.6;\nlet A2 = x;\nlet A3 = sum(A1:A2);",
			output: "let A1 = 5.60;let A2 = x;let A3 = sum(A1, A2);",
		},
		{
			name:   "one statements",
			input:  "A1",
			output: "A1;",
		},
		{
			name:   "no statements",
			input:  "",
			output: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			program, _, err := parser.Parse()
			if err != nil {
				t.Errorf("Unexpected error '%v'", err)
			}

			literal := ""
			for _, statement := range program {
				literal += statement.String()
			}

			output := tc.output
			if output == "" {
				output = tc.input
			}

			if literal != output {
				t.Errorf("Expected '%s' to equal '%s'", literal, output)
			}
		})
	}
}
