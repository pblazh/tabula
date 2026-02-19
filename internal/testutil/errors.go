package testutil

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
)

func ErrNoStatementsParsed() error {
	return fmt.Errorf("no statements parsed")
}

func ErrExpectedExpressionStatement(actual ast.Node) error {
	return fmt.Errorf("expected ExpressionStatement, got %T", actual)
}

func ErrParse(input string, err error) error {
	return fmt.Errorf("failed to parse %s, %s", input, err)
}
