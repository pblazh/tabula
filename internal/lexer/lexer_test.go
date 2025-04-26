package lexer

import (
	"strings"
	"testing"
	"text/scanner"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:  "supported tokens",
			input: "=+-*/()",
			expected: []Token{
				{
					Type: ASSIGN,
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type: PLUS,
					Position: scanner.Position{
						Column: 2,
					},
				},
				{
					Type: MINUS,
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type: MULT,
					Position: scanner.Position{
						Column: 4,
					},
				},
				{
					Type: DIV,
					Position: scanner.Position{
						Column: 5,
					},
				},
				{
					Type: LPAREN,
					Position: scanner.Position{
						Column: 6,
					},
				},
				{
					Type: RPAREN,
					Position: scanner.Position{
						Column: 7,
					},
				},
				{
					Type: EOF,
					Position: scanner.Position{
						Column: 8,
					},
				},
			},
		},
		{
			name:  "expression",
			input: "A1=A2+34 * SUM(B1:B2, 9)",
			expected: []Token{
				{
					Type:    IDENT,
					Literal: "A1",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type: ASSIGN,
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    IDENT,
					Literal: "A2",
					Position: scanner.Position{
						Column: 4,
					},
				},
				{
					Type: PLUS,
					Position: scanner.Position{
						Column: 6,
					},
				},
				{
					Type:    INT,
					Literal: "34",
					Position: scanner.Position{
						Column: 7,
					},
				},
				{
					Type: MULT,
					Position: scanner.Position{
						Column: 10,
					},
				},
				{
					Type:    IDENT,
					Literal: "SUM",
					Position: scanner.Position{
						Column: 12,
					},
				},
				{
					Type: LPAREN,
					Position: scanner.Position{
						Column: 15,
					},
				},
				{
					Type:    IDENT,
					Literal: "B1",
					Position: scanner.Position{
						Column: 16,
					},
				},
				{
					Type: COLUMN,
					Position: scanner.Position{
						Column: 18,
					},
				},
				{
					Type:    IDENT,
					Literal: "B2",
					Position: scanner.Position{
						Column: 19,
					},
				},
				{
					Type: COMA,
					Position: scanner.Position{
						Column: 21,
					},
				},
				{
					Type:    INT,
					Literal: "9",
					Position: scanner.Position{
						Column: 23,
					},
				},
				{
					Type: RPAREN,
					Position: scanner.Position{
						Column: 24,
					},
				},
				{
					Type: EOF,
					Position: scanner.Position{
						Column: 25,
					},
				},
			},
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		lexer := New(r, test.name)

		for _, token := range test.expected {
			nextToken := lexer.Next()
			EqualTokens(t, nextToken, token)
		}
	}
}

func EqualTokens(t testing.TB, a, b Token) {
	t.Helper()

	if a.Type != b.Type ||
		a.Literal != b.Literal ||
		a.Column != b.Column {
		t.Errorf("Expected %v to equal %v", a, b)
	}
}
