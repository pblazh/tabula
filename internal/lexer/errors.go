package lexer

import (
	"fmt"
	"text/scanner"
)

func ErrLexerError(message string, position scanner.Position) error {
	return fmt.Errorf("%s at %s", message, FormatPosition(position))
}
