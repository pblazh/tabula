package ast

import (
	"fmt"
)

func ErrInvalidRange(start, end string) error {
	return fmt.Errorf("range must contain valid cell references (like A1:B2), got %s:%s", start, end)
}

func ErrCircularDependency() error {
	return fmt.Errorf("circular dependency detected")
}
