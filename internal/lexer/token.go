// Package lexer provides lexical analysis for the CSV spreadsheet language.
// It tokenizes input text into tokens that can be processed by the parser.
package lexer

import (
	"fmt"
	"path"
	"text/scanner"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	scanner.Position
}

func (t Token) String() string {
	return FormatPosition(t.Position)
}

func FormatPosition(pos scanner.Position) string {
	var fileName string
	if pos.Filename == "" {
		fileName = "input"
	} else {
		_, fileName = path.Split(pos.Filename)
	}
	return fmt.Sprintf("%s:%d:%d", fileName, pos.Line, pos.Column)
}

const (
	ERROR            = "ERROR"
	ILLEGAL          = "ILLEGAL"
	IDENT            = "IDENT"
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
	COMMA            = "COMMA"
	LET              = "LET"
	FMT              = "FMT"
	INCLUDE          = "INCLUDE"
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

var tokens = map[string]TokenType{
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"*":  MULT,
	"/":  DIV,
	"%":  REM,
	"(":  LPAREN,
	")":  RPAREN,
	":":  COLUMN,
	",":  COMMA,
	";":  SEMI,
	"<":  LESS,
	">":  GREATER,
	">=": GREATER_OR_EQUAL,
	"<=": LESS_OR_EQUAL,
	"!":  NOT,
	"==": EQUAL,
	"!=": NOT_EQUAL,
}
