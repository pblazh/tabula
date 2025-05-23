package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestPaserErrors(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "missed assign value",
			input:  "let A1 = ;",
			output: "unexpected ; at missed assign value:1:10",
		},
		{
			name:   "missed assign identifier",
			input:  "let = 8;",
			output: "expected an identifier, but got = at missed assign identifier:1:5",
		},
		{
			name:   "not terminated quote",
			input:  "let x = \"9;",
			output: "literal not terminated at not terminated quote:1:9",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			res, err := parser.Parse()

			if err == nil {
				t.Errorf("Expected '%s' to return error '%s', but got %s", tc.input, tc.output, res)
			}
			if err != nil && err.Error() != tc.output {
				t.Errorf("Expected '%s' to equal '%s'", err.Error(), tc.output)
			}
		})
	}
}
