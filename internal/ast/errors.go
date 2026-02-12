package ast

import (
	"fmt"
	"strings"
	"text/scanner"
)

func ErrInvalidRange(start, end string) error {
	return fmt.Errorf("range must contain valid cell references (like A1:B2), got %s:%s", start, end)
}

func ErrCircularDependency() error {
	return fmt.Errorf("circular dependency detected")
}

func ErrIncludeFileNotFound(path string, position scanner.Position) error {
	return fmt.Errorf("include file not found: %s at %v", path, position)
}

func ErrCircularInclude(chain []string) error {
	return fmt.Errorf("circular include dependency detected: %s", strings.Join(chain, " → "))
}

func ErrIncludeReadError(path string, err error) error {
	return fmt.Errorf("failed to read include file %s: %w", path, err)
}
