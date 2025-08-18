package core

import (
	"fmt"
	"time"

	"github.com/pblazh/csvss/internal/ast"
)

// WriteValue writes an AST expression to context with optional format specification
func WriteValue(value ast.Expression, format string) (string, error) {
	if format == "" {
		formatted, err := formatWithoutSpec(value)
		if err != nil {
			return "", err
		}
		return formatted, nil
	}

	// Format the value using the format specification
	formattedValue, err := formatWithSpec(value, format)
	if err != nil {
		return "", err
	}
	return formattedValue, nil
}

// formatWithSpec formats an AST expression using the specified format
func formatWithSpec(value ast.Expression, format string) (string, error) {
	switch expr := value.(type) {
	case ast.IntExpression:
		return fmt.Sprintf(format, expr.Value), nil
	case ast.FloatExpression:
		return fmt.Sprintf(format, expr.Value), nil
	case ast.StringExpression:
		// Extract content from quoted string literal
		content := expr.Value
		if len(content) >= 2 && content[0] == '"' && content[len(content)-1] == '"' {
			content = content[1 : len(content)-1]
		}
		return fmt.Sprintf(format, content), nil
	case ast.BooleanExpression:
		return fmt.Sprintf(format, expr.Value), nil
	case ast.DateExpression:
		return expr.Value.Format(format), nil
	default:
		return "", ErrUnsupportedExpressionType(value)
	}
}

// formatWithoutSpec formats an AST expression without format specification
func formatWithoutSpec(value ast.Expression) (string, error) {
	switch expr := value.(type) {
	case ast.IntExpression:
		return fmt.Sprintf("%d", expr.Value), nil
	case ast.FloatExpression:
		return fmt.Sprintf("%g", expr.Value), nil
	case ast.StringExpression:
		return expr.Value, nil
	case ast.BooleanExpression:
		return fmt.Sprintf("%t", expr.Value), nil
	case ast.DateExpression:
		return expr.Value.Format(time.DateTime), nil
	default:
		return "", ErrUnsupportedExpressionType(value)
	}
}
