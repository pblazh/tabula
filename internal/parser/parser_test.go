package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func TestPaser(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected ast.Program
	}{
		{
			name:  "simple assign",
			input: "let A1 = 10;",
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

			literal := program[0].Literal()
			if literal != tc.input {
				t.Errorf("Expected '%s' to equal '%s'", literal, tc.input)
			}
		})
	}
}
