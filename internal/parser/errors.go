package parser

import (
	"fmt"
	"text/scanner"

	"github.com/pblazh/csvss/internal/lexer"
)

func ErrExpectedIdentifier(literal string, position scanner.Position) error {
	return fmt.Errorf("expected an identifier, but got %s at %v", literal, position)
}

func ErrExpectedToken(expected string, actual lexer.Token) error {
	return fmt.Errorf("expected %s, but got %v", expected, actual)
}

func ErrUnexpectedToken(literal string, position scanner.Position) error {
	return fmt.Errorf("unexpected %s at %v", literal, position)
}

func ErrExpectedRightParen(actual lexer.Token) error {
	return fmt.Errorf("expected right paren, but got %v", actual)
}

func ErrExpectedPrefix(actual lexer.Token) error {
	return fmt.Errorf("expected prefix, but got %v", actual)
}
