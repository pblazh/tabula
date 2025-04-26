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

func (l *Lexer) Next() Token {
	tok := l.scanner.Scan()
	literal := l.scanner.TokenText()

	if tok == scanner.EOF {
		return Token{
			Type:     EOF,
			Position: l.scanner.Position,
		}
	}

	if unicode.IsDigit(rune(literal[0])) {
		return Token{
			Type:     INT,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if unicode.IsLetter(rune(literal[0])) {
		return Token{
			Type:     IDENT,
			Literal:  literal,
			Position: l.scanner.Position,
		}
	}

	if t, ok := lexems[tok]; ok {
		return Token{
			Type:     t,
			Position: l.scanner.Position,
		}
	}

	return Token{
		Type: ILLEGAL,
	}
}

func New(r io.Reader, filename string) *Lexer {
	var s scanner.Scanner
	s.Init(r)
	s.Filename = filename

	return &Lexer{s, filename}
}
