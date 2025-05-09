package lexer

import (
	"io"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	scanner  scanner.Scanner
	Filename string
}

func (l *Lexer) Next() Lexem {
	tok := l.scanner.Scan()
	literal := l.scanner.TokenText()

	if tok == scanner.EOF {
		return Lexem{
			Type:     EOF,
			Position: l.scanner.Position,
		}
	}

	if literal == "=" && l.scanner.Peek() == '=' {
		lex := Lexem{
			Type:     EQUAL,
			Literal:  "==",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lex
	}

	if literal == "!" && l.scanner.Peek() == '=' {
		lex := Lexem{
			Type:     NOT_EQUAL,
			Literal:  "!=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lex
	}

	if literal == ">" && l.scanner.Peek() == '=' {
		lex := Lexem{
			Type:     GREATER_OR_EQUAL,
			Literal:  ">=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lex
	}

	if literal == "<" && l.scanner.Peek() == '=' {
		lex := Lexem{
			Type:     LESS_OR_EQUAL,
			Literal:  "<=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lex
	}

	if literal == "let" {
		return Lexem{
			Type:     LET,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if literal == "true" {
		return Lexem{
			Type:     TRUE,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if literal == "false" {
		return Lexem{
			Type:     FALSE,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if unicode.IsDigit(rune(literal[0])) {
		return Lexem{
			Type:     NUMBER,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if unicode.IsLetter(rune(literal[0])) {
		return Lexem{
			Type:     IDENT,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if t, ok := lexems[literal]; ok {
		return Lexem{
			Type:     t,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	return Lexem{
		Type:     ILLEGAL,
		Literal:  literal,
		Position: l.scanner.Position,
	}
}

func New(r io.Reader, filename string) *Lexer {
	var s scanner.Scanner
	s.Init(r)
	s.Filename = filename

	return &Lexer{s, filename}
}
