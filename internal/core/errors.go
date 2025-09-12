// Package core provides built-in functions for the CSV spreadsheet language
package core

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
)

func ErrUnsupportedArity(format string, function ast.CallExpression, expected, given int) error {
	if expected == 1 {
		return fmt.Errorf("%s expects %d argument, got %d in %s, at %v", format, expected, given, function, function.Token)
	}
	return fmt.Errorf("%s expects %d arguments, got %d in %s, at %v", format, expected, given, function, function.Token)
}

func ErrUnsupportedArgument(format string, function ast.CallExpression, argument ast.Expression) error {
	return fmt.Errorf("%s received invalid argument %s in %s, at %v", format, argument, function, function.Token)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s at %v", function, function.Token)
}

func ErrExecuting(format string, function ast.CallExpression, err error) error {
	return fmt.Errorf("failed %s with %v at %v", format, function.Token, err)
}

func ErrParseWithFormat(input, format, reason string) error {
	return fmt.Errorf("failed to parse %q with format %q:%s", input, format, reason)
}

func ErrUnsupportedExpressionType(expr ast.Expression) error {
	return fmt.Errorf("unsupported expression type: %T", expr)
}
