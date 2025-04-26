package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestPaser(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected Program
	}{
		{
			name:     "empty",
			input:    "",
			expected: []Statement{},
		},
		{
			name:     "simple assign",
			input:    "A1 = 10",
			expected: []Statement{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(strings.NewReader(tc.input), tc.name)
			var tokens []lexer.Token
			for t := l.Next(); t.Type != lexer.EOF; t = l.Next() {
				fmt.Println(t)
				tokens = append(tokens, t)
			}

			program := Parse(tokens)
			if len(program) != 0 {
				t.Errorf("%v", program)
			}
		})
	}
}
