package parser

import (
	"fmt"
	"text/scanner"

	"github.com/pblazh/tabula/internal/lexer"
)

func ErrExpectedIdentifier(literal string, position scanner.Position) error {
	return fmt.Errorf(
		"expected identifier, got %s at %s",
		literal,
		lexer.FormatPosition(position),
	)
}

func ErrExpectedToken(expected string, actual lexer.Token) error {
	return fmt.Errorf(
		"expected %s, got %s at %s",
		expected,
		actual.Literal,
		lexer.FormatPosition(actual.Position),
	)
}

func ErrUnexpectedToken(literal string, position scanner.Position) error {
	return fmt.Errorf(
		"unexpected %s at %s",
		literal,
		lexer.FormatPosition(position),
	)
}

func ErrExpectedPrefix(actual lexer.Token) error {
	return fmt.Errorf("expected prefix, got %s", actual)
}

func ErrParseInclude(path string, err error) error {
	return fmt.Errorf("%s: %w", path, err)
}
