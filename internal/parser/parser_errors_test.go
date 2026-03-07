package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/lexer"
)

func TestParserErrors(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "missed assign value",
			input:  `let A1 = ;`,
			output: "unexpected ; at missed assign value:1:10",
		},
		{
			name:   "missed assign identifier",
			input:  `let = 8;`,
			output: "expected identifier, got = at missed assign identifier:1:5",
		},
		{
			name:   "not terminated quote",
			input:  `let x = "9;`,
			output: "can not read expression, literal not terminated at not terminated quote:1:9",
		},
		{
			name:   "invalid range with variables",
			input:  `A:B;`,
			output: "failed to expand, invalid range A:B at invalid range with variables:1:2",
		},
		{
			name:   "invalid fmt statements int",
			input:  `fmt A1 = 1;`,
			output: "expected string, got 1 at invalid fmt statements int:1:10",
		},
		{
			name:   "invalid fmt statements bool",
			input:  `fmt A1 = true;`,
			output: "expected string, got true at invalid fmt statements bool:1:10",
		},
		{
			name:   "unsupported operator",
			input:  `2 : 3`,
			output: `expected identifier, got 2 at unsupported operator:1:6`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			res, _, err := parser.Parse()

			if err == nil {
				t.Errorf("Expected '%s' to return error '%s', got %s", tc.input, tc.output, res)
			}
			if err != nil && err.Error() != tc.output {
				t.Errorf("Expected '%s' to equal '%s'", err.Error(), tc.output)
			}
		})
	}
}
