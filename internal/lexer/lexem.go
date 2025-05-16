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
	IDENT            = "IDENT"
	ILLEGAL          = "ILLEGAL"
	EOF              = "EOF"
	INT              = "INT"
	FLOAT            = "FLOAT"
	TRUE             = "TRUE"
	STRING           = "STRING"
	FALSE            = "FALSE"
	ASSIGN           = "ASSIGN"
	PLUS             = "PLUS"
	MINUS            = "MINUS"
	MULT             = "MULT"
	DIV              = "DIV"
	REM              = "REM"
	LPAREN           = "LPAREN"
	RPAREN           = "RPAREN"
	COLUMN           = "COLUMN"
	COMA             = "COMA"
	LET              = "LET"
	SEMI             = "SEMI"
	EQUAL            = "EQUAL"
	NOT_EQUAL        = "NOT_EQUAL"
	LESS             = "LESS"
	GREATER          = "GREATER"
	GREATER_OR_EQUAL = "GREATER_OR_EQUAL"
	LESS_OR_EQUAL    = "LESS_OR_EQUAL"
	NOT              = "NOT"
	AND              = "AND"
	OR               = "OR"
)

var lexems = map[string]LexemType{
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"*":  MULT,
	"/":  DIV,
	"%":  REM,
	"(":  LPAREN,
	")":  RPAREN,
	":":  COLUMN,
	",":  COMA,
	";":  SEMI,
	"<":  LESS,
	">":  GREATER,
	">=": GREATER_OR_EQUAL,
	"<=": LESS_OR_EQUAL,
	"!":  NOT,
	"==": EQUAL,
	"!=": NOT_EQUAL,
}
