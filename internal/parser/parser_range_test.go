package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/lexer"
)

func TestParserRanges(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "range expression",
			input:  `a1:B1;`,
			output: "A1:B1;",
		},
		{
			name:   "range with parentheses",
			input:  `(A1:C1);`,
			output: "A1:C1;",
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

			var literal strings.Builder
			for _, statement := range program {
				literal.WriteString(statement.String())
			}

			output := tc.output
			if output == "" {
				output = tc.input
			}

			if literal.String() != output {
				t.Errorf("Expected '%s' to equal '%s'", literal.String(), output)
			}
		})
	}
}
