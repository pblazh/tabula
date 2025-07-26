// Package functions provides built-in functions for the CSV spreadsheet language
package functions

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

func ErrUnsupportedArity(format string, function ast.CallExpression, expected, given int) error {
	if expected == 1 {
		return fmt.Errorf("%s expected %d argument, but got %d in %s", format, expected, given, function)
	}
	return fmt.Errorf("%s expected %d arguments, but got %d in %s", format, expected, given, function)
}

func ErrUnsupportedArgument(format string, function ast.CallExpression, argument ast.Expression) error {
	return fmt.Errorf("%s got a wrong argument %s in %s", format, argument, function)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s", function)
}
