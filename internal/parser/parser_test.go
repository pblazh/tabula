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
			name:   "infix",
			input:  "5 + 6 - 2;",
			output: "((5 + 6) - 2);",
		},
		{
			name:   "infix precedence",
			input:  "5 + 6 * 2;",
			output: "(5 + (6 * 2));",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			program := parser.Parse()

			if len(program) != 1 {
				t.Errorf("Expected one statement but got '%d'", len(program))
				return
			}

			literal := program[0].String()
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
