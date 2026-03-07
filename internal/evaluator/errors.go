package evaluator

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func ErrDivisionByZero(token lexer.Token) error {
	return fmt.Errorf("division by zero at %s", token)
}

func ErrUnsupportedCall(expr ast.Node, target string) error {
	return fmt.Errorf("invalid argument %s for %s", target, expr)
}

func ErrCellOutOfBounds(cellRef, dimension string, index int) error {
	return fmt.Errorf("%s %d out of bounds for cell %s", dimension, index, cellRef)
}

func ErrRelOutOfBounds(expr ast.Node) error {
	return fmt.Errorf("%s is out of bounds", expr)
}

func ErrVariableNotFound(expr ast.Node) error {
	return fmt.Errorf("undefined: %s", expr)
}

func ErrUnknownExpressionType(expr ast.Node) error {
	return fmt.Errorf("unknown expression type %T", expr)
}

func ErrUnsupportedOperation(operator lexer.Token, expr ast.Node) error {
	return fmt.Errorf(
		"operator %s not supported for %s at %s",
		operator.Literal,
		ast.TypeName(expr),
		lexer.FormatPosition(operator.Position),
	)
}

func ErrUnsupportedType(receiver ast.Node, expr ast.Node) error {
	return fmt.Errorf("%s not supported by %s", ast.TypeName(expr), receiver)
}

func ErrUnsupportedBinaryOperation(operator lexer.Token, left, right ast.Node) error {
	return fmt.Errorf(
		"operator %s not supported for %s and %s at %s",
		operator.Literal,
		ast.TypeName(left),
		ast.TypeName(right),
		lexer.FormatPosition(operator.Position),
	)
}

func ErrUnsupportedPrefixOperator(operator lexer.Token) error {
	return fmt.Errorf(
		"unsupported prefix operator %s at %s",
		operator.Literal,
		lexer.FormatPosition(operator.Position),
	)
}

func ErrUnsupportedOperator(operator lexer.Token) error {
	return fmt.Errorf(
		"unsupported operator %s at %s",
		operator.Literal,
		lexer.FormatPosition(operator.Position),
	)
}

func ErrFmtExpectedString(identifier lexer.Token, actualValue string) error {
	return fmt.Errorf(
		"fmt accepts only strings, got %s at %s",
		actualValue,
		lexer.FormatPosition(identifier.Position),
	)
}

func ErrStatementExecution(statement string, err error) error {
	return fmt.Errorf("%s: %w", statement, err)
}

func ErrUnsupportedFunctions(identifier string) error {
	return fmt.Errorf("unsupported function: %s", identifier)
}

func ErrParsing(scriptName string, err error) error {
	if scriptName == "" {
		return err
	}
	return fmt.Errorf("%s: %w", scriptName, err)
}

func ErrWriting(err error) error {
	return fmt.Errorf("can not write value: %w", err)
}
