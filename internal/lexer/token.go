package lexer

import (
	"fmt"
	"text/scanner"
)

type TokenPosition struct {
	Filename string
	Line     int
	Column   int
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	scanner.Position
}

func (t Token) String() string {
	return fmt.Sprintf("<Token %s:%s at %v>", t.Type, t.Literal, t.Position)
}

const (
	IDENT   = "IDENT"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"
	ASSIGN  = "ASSIGN"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	MULT    = "MULT"
	DIV     = "DIV"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	COLUMN  = "COLUMN"
	COMA    = "COMA"
)

var lexems map[rune]TokenType = map[rune]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'*': MULT,
	'/': DIV,
	'(': LPAREN,
	')': RPAREN,
	':': COLUMN,
	',': COMA,
}
