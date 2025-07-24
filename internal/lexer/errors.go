package lexer

import (
	"fmt"
	"text/scanner"
)

func ErrLexerError(message string, position scanner.Position) error {
	return fmt.Errorf("Lexer error: %s at %v", message, position)
}
