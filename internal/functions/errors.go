// Package functions provides built-in functions for the CSV spreadsheet language
package functions

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

func ErrUnsupportedArgument(function ast.CallExpression, argument ast.Expression) error {
	return fmt.Errorf("unsupported argument %s for for %s", argument, function)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s", function)
}
