package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func parseNumberWithoutFormat(value string) ast.Expression {
	// Trim whitespace as the first step
	value = strings.TrimSpace(value)

	// Check if it's a number with a dot
	floatRegex := regexp.MustCompile(`^[+-]*\d+\.\d+`)
	if floatRegex.Match([]byte(value)) {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			return ast.FloatExpression{Value: parsed, Token: lexer.Token{Literal: value}}
		}
	}

	// Check if it's a number
	intRegex := regexp.MustCompile(`^[+-]*\d+`)
	if intRegex.Match([]byte(value)) {
		if intVal, err := strconv.Atoi(value); err == nil {
			return ast.IntExpression{Value: intVal, Token: lexer.Token{Literal: value}}
		}
	}

	return nil
}
