package evaluator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func ReadValue(value string, format string) (ast.Expression, error) {
	if format == "" {
		return parseWithoutFormat(value)
	}
	// Check if format is correct
	if err := validateFormatString(format); err != nil {
		return nil, ErrInvalidFormat(format, err.Error())
	}

	// Use scanf with the format specification
	placeholderType := detectPlaceholderType(format)

	switch placeholderType {
	case intPlacehoder:
		return parseInt(value, format)
	case floatPlacehoder:
		return parseFloat(value, format)
	case stringPlacehoder:
		return parseString(value, format)
	case boolPlacehoder:
		return parseBool(value, format)
	default:
		return nil, ErrParseWithFormat(value, format, "")
	}
}

// parseInt parses an integer value using the specified format
func parseInt(value, formatSpec string) (ast.Expression, error) {
	var resultInt int
	_, err := fmt.Sscanf(value, formatSpec, &resultInt)
	if err != nil {
		return nil, ErrParseWithFormat(value, formatSpec, err.Error())
	}
	return ast.IntExpression{Value: resultInt, Token: lexer.Token{Literal: value}}, nil
}

// cleanFormat removes the width and precision parts from format strings (e.g., %-9s → %s, %6f → %f, %6.2f → %f)
func cleanFormat(format string) string {
	re := regexp.MustCompile(`%([-+0# ]*)(\d+)?(\.\d+)?([a-zA-Z])`)
	return re.ReplaceAllString(format, `%$4`)
}

// parseFloat parses a float value using the specified format
func parseFloat(value, formatSpec string) (ast.Expression, error) {
	var resultFloat float64
	cleaned := cleanFormat(formatSpec)
	_, err := fmt.Sscanf(value, cleaned, &resultFloat)
	if err != nil {
		return nil, ErrParseWithFormat(value, formatSpec, err.Error())
	}
	return ast.FloatExpression{Value: resultFloat, Token: lexer.Token{Literal: value}}, nil
}

// parseString parses a string value using the specified format
func parseString(value, format string) (ast.Expression, error) {
	var resultString string
	cleaned := cleanFormat(format)
	_, err := fmt.Sscanf(value, cleaned, &resultString)
	if err != nil {
		return nil, ErrParseWithFormat(value, cleaned, err.Error())
	}
	return ast.StringExpression{Value: resultString, Token: lexer.Token{Literal: value}}, nil
}

// parseBool parses a boolean value using the specified format
func parseBool(value, formatSpec string) (ast.Expression, error) {
	var resultBool bool
	_, err := fmt.Sscanf(value, formatSpec, &resultBool)
	if err != nil {
		return nil, ErrParseWithFormat(value, formatSpec, err.Error())
	}
	return ast.BooleanExpression{Value: resultBool, Token: lexer.Token{Literal: value}}, nil
}

// parseWithoutFormat parses a value without format specification using default rules
func parseWithoutFormat(value string) (ast.Expression, error) {
	// Trim whitespace as the first step
	value = strings.TrimSpace(value)

	// Check if it's a string (enclosed in quotes) - return content as string
	quotedStringRegex := regexp.MustCompile(`^".*"$`)
	if quotedStringRegex.Match([]byte(value)) {
		return ast.StringExpression{Value: value, Token: lexer.Token{Literal: value}}, nil
	}

	datetime, err := parseDateWithoutFormat(value)
	if err != nil {
		return nil, err
	}
	if datetime != nil {
		return ast.DateExpression{Value: *datetime, Token: lexer.Token{Literal: value}}, nil
	}

	// Check if it's a number with a dot
	floatRegex := regexp.MustCompile(`^[+-]*\d+\.\d+`)
	if floatRegex.Match([]byte(value)) {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			return ast.FloatExpression{Value: parsed, Token: lexer.Token{Literal: value}}, nil
		}
	}

	// Check if it's a number
	intRegex := regexp.MustCompile(`^[+-]*\d+`)
	if intRegex.Match([]byte(value)) {
		if intVal, err := strconv.Atoi(value); err == nil {
			return ast.IntExpression{Value: intVal, Token: lexer.Token{Literal: value}}, nil
		}
	}

	// Check for boolean values
	if value == "true" {
		return ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: value}}, nil
	}
	if value == "false" {
		return ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: value}}, nil
	}

	// Otherwise return content as string
	return ast.StringExpression{Value: value, Token: lexer.Token{Literal: value}}, nil
}

func parseDateWithoutFormat(value string) (*time.Time, error) {
	formats := []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"02.01.2006",
		"02.01.2006 15:04",
		"02.01.2006 15:04:05",
		"01/02/2006",
		"01/02/2006 15:04",
		"01/02/2006 15:04:05",
		time.DateTime,
		time.TimeOnly,
		time.Kitchen,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			return &t, nil
		}
	}

	return nil, nil
}

// WriteValue writes an AST expression to context with optional format specification
func WriteValue(value ast.Expression, format string) (string, error) {
	if format == "" {
		formatted, err := formatWithoutSpec(value)
		if err != nil {
			return "", err
		}
		return formatted, nil
	}

	if err := validateFormatString(format); err != nil {
		return "", ErrInvalidFormatWrapper(format, err)
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
	default:
		return "", ErrUnsupportedExpressionType(value)
	}
}
