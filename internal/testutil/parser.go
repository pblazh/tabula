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

// ParseProgram is a helper function to parse a program from a string
func ParseProgram(input string) (ast.Program, error) {
	lex := lexer.New(strings.NewReader(input), "test")
	p := parser.New(lex)
	program, _, err := p.Parse()
	return program, err
}

// CompareExpressions compares two AST expressions for equality
func CompareExpressions(a, b ast.Expression) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	switch aExpr := a.(type) {
	case ast.IntExpression:
		if bExpr, ok := b.(ast.IntExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.FloatExpression:
		if bExpr, ok := b.(ast.FloatExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.StringExpression:
		if bExpr, ok := b.(ast.StringExpression); ok {
			return aExpr.Token.Literal == bExpr.Token.Literal
		}
	case ast.BooleanExpression:
		if bExpr, ok := b.(ast.BooleanExpression); ok {
			return aExpr.Value == bExpr.Value && aExpr.Token.Literal == bExpr.Token.Literal
		}
	}
	return false
}
