package ast

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/pblazh/tabula/internal/lexer"
)

func ErrInvalidRange(start, end string) error {
	return fmt.Errorf("invalid range %s:%s", start, end)
}

func ErrCircularDependency() error {
	return fmt.Errorf("circular dependency detected")
}

func ErrIncludeFileNotFound(path string, position scanner.Position) error {
	return fmt.Errorf("include file not found: %s at %s", path, lexer.FormatPosition(position))
}

func ErrCircularInclude(chain []string) error {
	return fmt.Errorf("circular include: %s", strings.Join(chain, " → "))
}

func ErrIncludeReadError(path string, err error) error {
	return fmt.Errorf("cannot read %s: %w", path, err)
}
