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
		expected []Lexem
	}{
		{
			name:  "supported tokens",
			input: "=+-*/()%",
			expected: []Lexem{
				{
					Type:    ASSIGN,
					Literal: "=",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    PLUS,
					Literal: "+",
					Position: scanner.Position{
						Column: 2,
					},
				},
				{
					Type:    MINUS,
					Literal: "-",
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    MULT,
					Literal: "*",
					Position: scanner.Position{
						Column: 4,
					},
				},
				{
					Type:    DIV,
					Literal: "/",
					Position: scanner.Position{
						Column: 5,
					},
				},
				{
					Type:    LPAREN,
					Literal: "(",
					Position: scanner.Position{
						Column: 6,
					},
				},
				{
					Type:    RPAREN,
					Literal: ")",
					Position: scanner.Position{
						Column: 7,
					},
				},
				{
					Type:    REM,
					Literal: "%",
					Position: scanner.Position{
						Column: 8,
					},
				},
				{
					Type: EOF,
					Position: scanner.Position{
						Column: 9,
					},
				},
			},
		},

		{
			name:  "expression",
			input: "let A1=A2+34 * SUM(B1:B2, 9.1) % 9;",
			expected: []Lexem{
				{
					Type:    LET,
					Literal: "let",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    IDENT,
					Literal: "A1",
					Position: scanner.Position{
						Column: 5,
					},
				},
				{
					Type:    ASSIGN,
					Literal: "=",
					Position: scanner.Position{
						Column: 7,
					},
				},
				{
					Type:    IDENT,
					Literal: "A2",
					Position: scanner.Position{
						Column: 8,
					},
				},
				{
					Type:    PLUS,
					Literal: "+",
					Position: scanner.Position{
						Column: 10,
					},
				},
				{
					Type:    INT,
					Literal: "34",
					Position: scanner.Position{
						Column: 11,
					},
				},
				{
					Type:    MULT,
					Literal: "*",
					Position: scanner.Position{
						Column: 14,
					},
				},
				{
					Type:    IDENT,
					Literal: "SUM",
					Position: scanner.Position{
						Column: 16,
					},
				},
				{
					Type:    LPAREN,
					Literal: "(",
					Position: scanner.Position{
						Column: 19,
					},
				},
				{
					Type:    IDENT,
					Literal: "B1",
					Position: scanner.Position{
						Column: 20,
					},
				},
				{
					Type:    COLUMN,
					Literal: ":",
					Position: scanner.Position{
						Column: 22,
					},
				},
				{
					Type:    IDENT,
					Literal: "B2",
					Position: scanner.Position{
						Column: 23,
					},
				},
				{
					Type:    COMA,
					Literal: ",",
					Position: scanner.Position{
						Column: 25,
					},
				},
				{
					Type:    FLOAT,
					Literal: "9.1",
					Position: scanner.Position{
						Column: 27,
					},
				},
				{
					Type:    RPAREN,
					Literal: ")",
					Position: scanner.Position{
						Column: 30,
					},
				},
				{
					Type:    REM,
					Literal: "%",
					Position: scanner.Position{
						Column: 32,
					},
				},
				{
					Type:    INT,
					Literal: "9",
					Position: scanner.Position{
						Column: 34,
					},
				},
				{
					Type:    SEMI,
					Literal: ";",
					Position: scanner.Position{
						Column: 35,
					},
				},
				{
					Type: EOF,
					Position: scanner.Position{
						Column: 36,
					},
				},
			},
		},

		{
			name:  "equal expression",
			input: "a == b",
			expected: []Lexem{
				{
					Type:    IDENT,
					Literal: "a",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    EQUAL,
					Literal: "==",
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    IDENT,
					Literal: "b",
					Position: scanner.Position{
						Column: 6,
					},
				},
			},
		},

		{
			name:  "not equal expression",
			input: "a != b",
			expected: []Lexem{
				{
					Type:    IDENT,
					Literal: "a",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    NOT_EQUAL,
					Literal: "!=",
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    IDENT,
					Literal: "b",
					Position: scanner.Position{
						Column: 6,
					},
				},
			},
		},

		{
			name:  "greater equal expression",
			input: "a >= b",
			expected: []Lexem{
				{
					Type:    IDENT,
					Literal: "a",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    GREATER_OR_EQUAL,
					Literal: ">=",
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    IDENT,
					Literal: "b",
					Position: scanner.Position{
						Column: 6,
					},
				},
			},
		},

		{
			name:  "less equal expression",
			input: "a <= b",
			expected: []Lexem{
				{
					Type:    IDENT,
					Literal: "a",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    LESS_OR_EQUAL,
					Literal: "<=",
					Position: scanner.Position{
						Column: 3,
					},
				},
				{
					Type:    IDENT,
					Literal: "b",
					Position: scanner.Position{
						Column: 6,
					},
				},
			},
		},

		{
			name:  "bool expressions",
			input: "true false",
			expected: []Lexem{
				{
					Type:    TRUE,
					Literal: "true",
					Position: scanner.Position{
						Column: 1,
					},
				},
				{
					Type:    FALSE,
					Literal: "false",
					Position: scanner.Position{
						Column: 6,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := strings.NewReader(test.input)
			lexer := New(r, test.name)

			for _, lexem := range test.expected {
				next := lexer.Next()
				EqualLexems(t, next, lexem)
			}
		})
	}
}

func EqualLexems(t testing.TB, a, b Lexem) {
	t.Helper()

	if a.Type != b.Type ||
		a.Literal != b.Literal ||
		a.Column != b.Column {
		t.Errorf("Expected %v to equal %v", a, b)
	}
}
