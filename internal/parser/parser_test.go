package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestPaser(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:  "simple assign",
			input: "let A1 = -10;",
		},
		{
			name:  "identifier",
			input: "something;",
		},
		{
			name:  "not identifier",
			input: "!something;",
		},
		{
			name:  "number",
			input: "10;",
		},
		{
			name:  "number",
			input: "-10;",
		},
		{
			name:   "true",
			input:  "true;",
			output: "<bool true>;",
		},
		{
			name:   "false",
			input:  "false;",
			output: "<bool false>;",
		},
		{
			name:   "infix",
			input:  "5 + 6 - 2;",
			output: "((5 + 6) - 2);",
		},
		{
			name:   "infix reverse precedence",
			input:  "5 + 6 * 2;",
			output: "(5 + (6 * 2));",
		},
		{
			name:   "infix precedence",
			input:  "5 / 6 + 2;",
			output: "((5 / 6) + 2);",
		},
		{
			name:   "multiple statements",
			input:  "let A1 = 5.6;\nlet A2 = x;\n",
			output: "let A1 = 5.6;let A2 = x;",
		},
		// {
		// 	name:   "multiple statements",
		// 	input:  "let A1 = 5.6;\nlet A2 = x;\nlet A3 = sum(A1:A2);",
		// 	output: "let A1 = 5.6;let A2 = x;let A3 = sum(A1:A2);",
		// },
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			program := parser.Parse()

			literal := ""
			for _, statement := range program {
				literal += statement.String()
			}

			output := tc.output
			if output == "" {
				output = tc.input
			}

			if literal != output {
				t.Errorf("Expected '%s' to equal '%s'", literal, tc.output)
			}
		})
	}
}
