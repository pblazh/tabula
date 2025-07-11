package evaluator

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func ErrDivisionByZero(token lexer.Token) error {
	return fmt.Errorf("division by zero at %s", token)
}

func ErrCellOutOfBounds(cellRef, dimension string, index int) error {
	return fmt.Errorf("%s %d out of bounds for cell %s", dimension, index, cellRef)
}

func ErrVariableNotFound(expr ast.Expression) error {
	return fmt.Errorf("%s not found in context", expr)
}

func ErrUnknownExpressionType(expr ast.Expression) error {
	return fmt.Errorf("unknown expression type: %T", expr)
}

func ErrUnsupportedOperation(operator lexer.Token, expr ast.Expression) error {
	return fmt.Errorf("%s is not supported for type: %s", operator, ast.TypeName(expr))
}

func ErrUnsupportedBinaryOperation(operator lexer.Token, left, right ast.Expression) error {
	return fmt.Errorf("operator %s is not supported for type: %s and %s", operator, ast.TypeName(left), ast.TypeName(right))
}

func ErrParseWithFormat(input, format, reason string) error {
	return fmt.Errorf("failed to parse %q with format %q:%s", input, format, reason)
}

func ErrInvalidFormat(format, reason string) error {
	return fmt.Errorf("invalid format %q: %s", format, reason)
}

func ErrInvalidFormatWrapper(format string, err error) error {
	return fmt.Errorf("invalid format %q: %w", format, err)
}

func ErrMissedPlaceholder() error {
	return fmt.Errorf("no scanf placeholder found")
}

func ErrManyPlaceholders(n int) error {
	return fmt.Errorf("multiple scanf placeholders found (%d), expected exactly one", n)
}

func ErrUnsupportedPrefixOperator(operator string) error {
	return fmt.Errorf("unsupported prefix operator: %s", operator)
}

func ErrUnsupportedOperator(operator string) error {
	return fmt.Errorf("unsupported operator: %s", operator)
}

func ErrUnsupportedExpressionType(expr ast.Expression) error {
	return fmt.Errorf("unsupported expression type: %T", expr)
}

func ErrFmtExpectedString(identifier lexer.Token, actualValue string) error {
	return fmt.Errorf("fmt %s accepts only strings, but got %s", identifier, actualValue)
}

func ErrStatementExecution(statement string, err error) error {
	return fmt.Errorf("%s caused %s", statement, err)
}
