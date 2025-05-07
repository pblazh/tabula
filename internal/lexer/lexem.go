package lexer

import (
	"fmt"
	"text/scanner"
)

type LexemType string

type Lexem struct {
	Type    LexemType
	Literal string
	scanner.Position
}

func (t Lexem) String() string {
	return fmt.Sprintf("<%s:%s at %v>", t.Type, t.Literal, t.Position)
}

const (
	IDENT     = "IDENT"
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	NUMBER    = "NUMBER"
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	MULT      = "MULT"
	DIV       = "DIV"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	COLUMN    = "COLUMN"
	COMA      = "COMA"
	LET       = "LET"
	SEMI      = "SEMI"
	NOT       = "NOT"
	EQUAL     = "EQUAL"
	NOT_EQUAL = "NOT_EQUAL"
	LESS      = "LESS"
	GREATER   = "GREATER"
)

var lexems map[rune]LexemType = map[rune]LexemType{
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'*': MULT,
	'/': DIV,
	'(': LPAREN,
	')': RPAREN,
	':': COLUMN,
	',': COMA,
	';': SEMI,
	'!': NOT,
}
