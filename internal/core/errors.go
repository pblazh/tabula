// Package core provides built-in functions for the CSV spreadsheet language
package core

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
)

func ErrUnsupportedArity(format string, function ast.CallExpression, expected, given int) error {
	if expected == 1 {
		return fmt.Errorf(
			"%s expects %d argument, got %d in %s, at %v",
			format,
			expected,
			given,
			function,
			function.Token,
		)
	}
	return fmt.Errorf(
		"%s expects %d arguments, got %d in %s, at %v",
		format,
		expected,
		given,
		function,
		function.Token,
	)
}

func ErrUnsupportedArgument(
	format string,
	function ast.CallExpression,
	argument ast.Node,
) error {
	return fmt.Errorf(
		"%s received an invalid argument %s in %s, at %v",
		format,
		argument,
		function,
		function.Token,
	)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s at %v", function, function.Token)
}

func ErrExecuting(format string, function ast.CallExpression, err error) error {
	return fmt.Errorf("failed %s with %v at %v", format, function.Token, err)
}

func ErrParseWithFormat(input, format, reason string) error {
	return fmt.Errorf("failed to parse %q with format %q, %s", input, format, reason)
}

func ErrUnsupportedExpressionType(expr ast.Node) error {
	return fmt.Errorf("unsupported expression type: %T", expr)
}

func ErrExecute(err error) error {
	return fmt.Errorf("fails to execute,  %s", err)
}

func ErrExpand(err error) error {
	return fmt.Errorf("failed to expand, %s", err)
}

func ErrParseString(err error) error {
	return fmt.Errorf("failed to parse string, %s", err)
}

func ErrParseBoolean(err error) error {
	return fmt.Errorf("failed to parse boolean, %s", err)
}

func ErrParseDate(err error) error {
	return fmt.Errorf("failed to parse date, %s", err)
}

func ErrParseInt(err error) error {
	return fmt.Errorf("failed to parse int, %s", err)
}

func ErrParseFloat(err error) error {
	return fmt.Errorf("failed to parse float, %s", err)
}
