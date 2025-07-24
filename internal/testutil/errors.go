package testutil

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

func ErrNoStatementsParsed() error {
	return fmt.Errorf("no statements parsed")
}

func ErrExpectedExpressionStatement(actual ast.Statement) error {
	return fmt.Errorf("expected ExpressionStatement, got %T", actual)
}
