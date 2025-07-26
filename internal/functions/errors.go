// Package functions provides built-in functions for the CSV spreadsheet language
package functions

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

func ErrUnsupportedArity(function ast.CallExpression, expected, given int) error {
	return fmt.Errorf("%s expected %d arguments, but got %d", function, expected, given)
}

func ErrUnsupportedArgument(function ast.CallExpression, argument ast.Expression) error {
	return fmt.Errorf("unsupported argument %s for %s", argument, function)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s", function)
}
