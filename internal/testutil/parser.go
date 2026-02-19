// Package testutil provides shared testing utilities for the Tabula project.
package testutil

import (
	"os"
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
	"github.com/pblazh/tabula/internal/parser"
)

// ParseExpression parses an input string as an expression and returns the AST.
func ParseExpression(t *testing.T, input string) ast.Expression {
	t.Helper()
	expr, err := parseExpression(input)
	if err != nil {
		t.Fatalf("fatal, %s", ErrParse(input, err))
	}
	return expr
}

func parseExpression(input string) (ast.Expression, error) {
	lex := lexer.New(strings.NewReader(input+";"), "test")
	p := parser.New(lex)
	program, _, err := p.Parse()
	if err != nil {
		return nil, ErrParse(input, err)
	}

	if len(program) == 0 {
		return nil, ErrNoStatementsParsed()
	}

	stmt, ok := program[0].(ast.ExpressionStatement)
	if !ok {
		return nil, ErrExpectedExpressionStatement(program[0])
	}

	return stmt.Value, nil
}

// ParseProgram is a helper function to parse a program from a string
func ParseProgram(input string) (ast.Program, error) {
	lex := lexer.New(strings.NewReader(input), "test")
	p := parser.New(lex)
	program, _, err := p.Parse()
	if err != nil {
		return nil, ErrParse(input, err)
	}
	return program, nil
}

// ParseProgramFromFile is a helper function to parse a program from a file
func ParseProgramFromFile(filename string) (ast.Program, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, ErrParse(filename, err)
	}
	return ParseProgram(string(content))
}
