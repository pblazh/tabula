package lexer

import (
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	scanner          *scanner.Scanner
	lastErrorMessage string
	errorCount       int
}

func (l *Lexer) Next() (Lexem, error) {
	tok := l.scanner.Scan()
	literal := l.scanner.TokenText()

	if l.scanner.ErrorCount > l.errorCount {
		l.errorCount = l.scanner.ErrorCount

		return Lexem{
			Type:     ERROR,
			Position: l.scanner.Position,
		}, fmt.Errorf("%s at %v", l.lastErrorMessage, l.scanner.Position)
	}

	if tok == scanner.EOF {
		return Lexem{
			Type:     EOF,
			Position: l.scanner.Position,
		}, nil
	}

	if literal[0] == '"' || literal[0] == '\'' {
		lexem := Lexem{
			Type:     STRING,
			Literal:  literal,
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lexem, nil
	}

	if literal == "=" && l.scanner.Peek() == '=' {
		lexem := Lexem{
			Type:     EQUAL,
			Literal:  "==",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lexem, nil
	}

	if literal == "!" && l.scanner.Peek() == '=' {
		lexem := Lexem{
			Type:     NOT_EQUAL,
			Literal:  "!=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lexem, nil
	}

	if literal == ">" && l.scanner.Peek() == '=' {
		lexem := Lexem{
			Type:     GREATER_OR_EQUAL,
			Literal:  ">=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lexem, nil
	}

	if literal == "<" && l.scanner.Peek() == '=' {
		lexem := Lexem{
			Type:     LESS_OR_EQUAL,
			Literal:  "<=",
			Position: l.scanner.Position,
		}
		l.scanner.Scan()
		return lexem, nil
	}

	if literal == "let" {
		return Lexem{
			Type:     LET,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if literal == "true" {
		return Lexem{
			Type:     TRUE,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if literal == "false" {
		return Lexem{
			Type:     FALSE,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if unicode.IsDigit(rune(literal[0])) && strings.Contains(literal, ".") {
		return Lexem{
			Type:     FLOAT,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if unicode.IsDigit(rune(literal[0])) && !strings.Contains(literal, ".") {
		return Lexem{
			Type:     INT,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if unicode.IsLetter(rune(literal[0])) {
		return Lexem{
			Type:     IDENT,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	if t, ok := lexems[literal]; ok {
		return Lexem{
			Type:     t,
			Literal:  literal,
			Position: l.scanner.Position,
		}, nil
	}

	return Lexem{
		Type:     ILLEGAL,
		Literal:  literal,
		Position: l.scanner.Position,
	}, nil
}

func New(r io.Reader, filename string) *Lexer {
	var s scanner.Scanner
	s.Filename = filename
	s.Init(r)

	lexer := &Lexer{
		scanner:          &s,
		lastErrorMessage: "",
		errorCount:       0,
	}

	s.Error = func(s *scanner.Scanner, msg string) {
		lexer.lastErrorMessage = msg
	}

	return lexer
}
