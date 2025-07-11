package lexer

import (
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

func (lexer *Lexer) Next() (Token, error) {
	tok := lexer.scanner.Scan()
	literal := lexer.scanner.TokenText()

	if lexer.scanner.ErrorCount > lexer.errorCount {
		lexer.errorCount = lexer.scanner.ErrorCount

		return Token{
			Type:     ERROR,
			Position: lexer.scanner.Position,
		}, ErrLexerError(lexer.lastErrorMessage, lexer.scanner.Position)
	}

	if tok == scanner.EOF {
		return Token{
			Type:     EOF,
			Position: lexer.scanner.Position,
		}, nil
	}

	if literal[0] == '"' || literal[0] == '\'' {
		token := Token{
			Type:     STRING,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}
		lexer.scanner.Scan()
		return token, nil
	}

	if literal == "=" && lexer.scanner.Peek() == '=' {
		token := Token{
			Type:     EQUAL,
			Literal:  "==",
			Position: lexer.scanner.Position,
		}
		lexer.scanner.Scan()
		return token, nil
	}

	if literal == "!" && lexer.scanner.Peek() == '=' {
		token := Token{
			Type:     NOT_EQUAL,
			Literal:  "!=",
			Position: lexer.scanner.Position,
		}
		lexer.scanner.Scan()
		return token, nil
	}

	if literal == ">" && lexer.scanner.Peek() == '=' {
		token := Token{
			Type:     GREATER_OR_EQUAL,
			Literal:  ">=",
			Position: lexer.scanner.Position,
		}
		lexer.scanner.Scan()
		return token, nil
	}

	if literal == "<" && lexer.scanner.Peek() == '=' {
		token := Token{
			Type:     LESS_OR_EQUAL,
			Literal:  "<=",
			Position: lexer.scanner.Position,
		}
		lexer.scanner.Scan()
		return token, nil
	}

	if literal == "let" {
		return Token{
			Type:     LET,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if literal == "fmt" {
		return Token{
			Type:     FMT,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if literal == "true" {
		return Token{
			Type:     TRUE,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if literal == "false" {
		return Token{
			Type:     FALSE,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if unicode.IsDigit(rune(literal[0])) && strings.Contains(literal, ".") {
		return Token{
			Type:     FLOAT,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if unicode.IsDigit(rune(literal[0])) && !strings.Contains(literal, ".") {
		return Token{
			Type:     INT,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if unicode.IsLetter(rune(literal[0])) {
		return Token{
			Type:     IDENT,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	if t, ok := tokens[literal]; ok {
		return Token{
			Type:     t,
			Literal:  literal,
			Position: lexer.scanner.Position,
		}, nil
	}

	return Token{
		Type:     ILLEGAL,
		Literal:  literal,
		Position: lexer.scanner.Position,
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
