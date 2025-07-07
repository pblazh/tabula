// Package testutil provides shared testing utilities for the csvss project.
package testutil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
	"github.com/pblazh/csvss/internal/parser"
)

// ParseExpression parses an input string as an expression and returns the AST.
func ParseExpression(t *testing.T, input string) ast.Expression {
	t.Helper()
	expr, err := parseExpression(input)
	if err != nil {
		t.Fatalf("Failed to parse expression %q: %v", input, err)
	}
	return expr
}

func parseExpression(input string) (ast.Expression, error) {
	lex := lexer.New(strings.NewReader(input+";"), "test")
	p := parser.New(lex)
	program, _, err := p.Parse()
	if err != nil {
		return nil, err
	}

	if len(program) == 0 {
		return nil, fmt.Errorf("no statements parsed")
	}

	stmt, ok := program[0].(ast.ExpressionStatement)
	if !ok {
		return nil, fmt.Errorf("expected ExpressionStatement, got %T", program[0])
	}

	return stmt.Value, nil
}
