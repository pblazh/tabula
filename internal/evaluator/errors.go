package evaluator

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func ErrDivisionByZero(token lexer.Token) error {
	return fmt.Errorf("division by zero at %s", token)
}

func ErrUnsupportedCall(expr ast.Expression, target string) error {
	return fmt.Errorf("invalid argument %s for %s", target, expr)
}

func ErrCellOutOfBounds(cellRef, dimension string, index int) error {
	return fmt.Errorf("%s %d out of bounds for cell %s", dimension, index, cellRef)
}

func ErrRelOutOfBounds(expr ast.Expression) error {
	return fmt.Errorf("%s is outof bounds", expr)
}

func ErrVariableNotFound(expr ast.Expression) error {
	return fmt.Errorf("%s not found in context", expr)
}

func ErrUnknownExpressionType(expr ast.Expression) error {
	return fmt.Errorf("unknown expression type %T", expr)
}

func ErrUnsupportedOperation(operator lexer.Token, expr ast.Expression) error {
	return fmt.Errorf("operator %s is not supported for type %s at %v", operator.Literal, ast.TypeName(expr), operator.Position)
}

func ErrUnsupportedType(receiver ast.Expression, expr ast.Expression) error {
	return fmt.Errorf("%s is not supported by %s", ast.TypeName(expr), receiver)
}

func ErrUnsupportedBinaryOperation(operator lexer.Token, left, right ast.Expression) error {
	return fmt.Errorf("operator %s is not supported for type %s and %s at %v", operator.Literal, ast.TypeName(left), ast.TypeName(right), operator.Position)
}

func ErrUnsupportedPrefixOperator(operator lexer.Token) error {
	return fmt.Errorf("unsupported prefix operator %s at %v", operator.Literal, operator.Position)
}

func ErrUnsupportedOperator(operator lexer.Token) error {
	return fmt.Errorf("unsupported operator %s at %v", operator.Literal, operator.Position)
}

func ErrFmtExpectedString(identifier lexer.Token, actualValue string) error {
	return fmt.Errorf("fmt accepts only strings, got %s at %v", actualValue, identifier.Position)
}

func ErrStatementExecution(statement string, err error) error {
	return fmt.Errorf("%s caused %s", statement, err)
}

func ErrUnsupportedFunctions(identifier string) error {
	return fmt.Errorf("unsupported function: %s", identifier)
}
