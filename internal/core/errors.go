// Package core provides built-in functions for the CSV spreadsheet language
package core

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
)

func ErrUnsupportedArity(format string, function ast.CallExpression, expected, given int) error {
	if expected == 1 {
		return fmt.Errorf(
			"%s expects %d argument, got %d at %v",
			format,
			expected,
			given,
			function.Token,
		)
	}
	return fmt.Errorf(
		"%s expects %d arguments, got %d at %v",
		format,
		expected,
		given,
		function.Token,
	)
}

func ErrUnsupportedArgument(
	format string,
	function ast.CallExpression,
	argument ast.Node,
) error {
	return fmt.Errorf(
		"%s invalid argument %s at %v",
		format,
		argument,
		function.Token,
	)
}

func ErrUnsupportedFunction(function ast.CallExpression) error {
	return fmt.Errorf("unsupported function call %s at %v", function, function.Token)
}

func ErrExecuting(format string, function ast.CallExpression, err error) error {
	return fmt.Errorf("failed %s with %v at %w", format, function.Token, err)
}

func ErrParseWithFormat(input, format, reason string) error {
	return fmt.Errorf("cannot parse %q with format %q", input, format)
}

func ErrUnsupportedExpressionType(expr ast.Node) error {
	return fmt.Errorf("unsupported expression type: %T", expr)
}

func ErrExecute(err error) error {
	return fmt.Errorf("cannot execute: %w", err)
}

func ErrExpand(err error) error {
	return fmt.Errorf("cannot expand: %w", err)
}

func ErrParse(typeName string, value string, err error) error {
	return fmt.Errorf("cannot parse %q as %s: %w", value, typeName, err)
}

func ErrParseBoolean(err error) error {
	return fmt.Errorf("cannot parse boolean: %w", err)
}

func ErrParseInt(err error) error {
	return fmt.Errorf("cannot parse int: %w", err)
}

func ErrParseFloat(err error) error {
	return fmt.Errorf("cannot parse float: %w", err)
}

func ErrParseDate(value, format string) error {
	return fmt.Errorf("cannot parse date: %s with format %s", value, format)
}
