package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestParserRanges(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "range expression",
			input:  "a1:B1;",
			output: "(: A1 B1);",
		},
		{
			name:   "range expression A1:C1",
			input:  "A1:C1;",
			output: "(: A1 B1 C1);",
		},
		{
			name:   "range expression vertical A1:A3",
			input:  "A1:A3;",
			output: "(: A1 A2 A3);",
		},
		{
			name:   "range expression reverse C1:A1",
			input:  "C1:A1;",
			output: "(: C1 B1 A1);",
		},
		{
			name:   "range expression single cell A1:A1",
			input:  "A1:A1;",
			output: "(: A1);",
		},
		{
			name:   "range expression in call",
			input:  "SUM(A1:C1);",
			output: "(SUM A1 B1 C1);",
		},
		{
			name:   "range expression in let statement",
			input:  "let total = SUM(B1:D1);",
			output: "let total = (SUM B1 C1 D1);",
		},
		{
			name:   "range expression large horizontal",
			input:  "A1:E1;",
			output: "(: A1 B1 C1 D1 E1);",
		},
		{
			name:   "range expression large vertical",
			input:  "B1:B5;",
			output: "(: B1 B2 B3 B4 B5);",
		},
		{
			name:   "range expression with multi-letter columns",
			input:  "AA1:AC1;",
			output: "(: AA1 AB1 AC1);",
		},
		{
			name:   "range expression in both rows and columns",
			input:  "A1:C3;",
			output: "(: A1 B1 C1 A2 B2 C2 A3 B3 C3);",
		},
		{
			name:   "multiple ranges in call",
			input:  "SUM(A1:B1, C1:D1);",
			output: "(SUM A1 B1 C1 D1);",
		},
		{
			name:   "range in complex expression",
			input:  "A1 + SUM(B1:D1) * 2;",
			output: "(+ A1 (* (SUM B1 C1 D1) <int 2>));",
		},
		{
			name:   "range with parentheses",
			input:  "(A1:C1);",
			output: "(: A1 B1 C1);",
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
