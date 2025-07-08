package evaluator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func ReadValue(name string, context map[string]string, format map[string]string) (ast.Expression, error) {
	value, exists := context[name]
	if !exists {
		return nil, fmt.Errorf("%s not found in context", name)
	}

	formatSpec, hasFormat := format[name]
	if !hasFormat {
		return parseWithoutFormat(value)
	}
	// Check if format is correct
	if err := validateFormatString(formatSpec); err != nil {
		return nil, fmt.Errorf("invalid format %q: %w", formatSpec, err)
	}

	// Use scanf with the format specification
	placeholderType := detectPlaceholderType(formatSpec)

	switch placeholderType {
	case "int":
		return parseInt(value, formatSpec)
	case "float":
		return parseFloat(value, formatSpec)
	case "string":
		return parseString(value, formatSpec)
	case "bool":
		return parseBool(value, formatSpec)
	default:
		return nil, fmt.Errorf("failed to parse %q with format %q", value, formatSpec)
	}
}

// parseInt parses an integer value using the specified format
func parseInt(value, formatSpec string) (ast.Expression, error) {
	var resultInt int
	_, err := fmt.Sscanf(value, formatSpec, &resultInt)
	if err != nil {
		return nil, newParseError(value, formatSpec, err)
	}
	return ast.IntExpression{Value: resultInt, Token: lexer.Token{Literal: fmt.Sprintf("%d", resultInt)}}, nil
}

// parseFloat parses a float value using the specified format
func parseFloat(value, formatSpec string) (ast.Expression, error) {
	var resultFloat float64
	_, err := fmt.Sscanf(value, formatSpec, &resultFloat)
	if err != nil {
		return nil, newParseError(value, formatSpec, err)
	}
	return ast.FloatExpression{Value: resultFloat, Token: lexer.Token{Literal: fmt.Sprintf("%g", resultFloat)}}, nil
}

// parseString parses a string value using the specified format
func parseString(value, formatSpec string) (ast.Expression, error) {
	var resultString string
	_, err := fmt.Sscanf(value, formatSpec, &resultString)
	if err != nil {
		return nil, newParseError(value, formatSpec, err)
	}
	return ast.StringExpression{Token: lexer.Token{Literal: "\"" + resultString + "\""}}, nil
}

// parseBool parses a boolean value using the specified format
func parseBool(value, formatSpec string) (ast.Expression, error) {
	var resultBool bool
	_, err := fmt.Sscanf(value, formatSpec, &resultBool)
	if err != nil {
		return nil, newParseError(value, formatSpec, err)
	}
	return ast.BooleanExpression{Value: resultBool, Token: lexer.Token{Literal: fmt.Sprintf("%t", resultBool)}}, nil
}

// parseWithoutFormat parses a value without format specification using default rules
func parseWithoutFormat(value string) (ast.Expression, error) {
	// Trim whitespace as the first step
	value = strings.TrimSpace(value)

	// Check if it's a string (enclosed in quotes) - return content as string
	quotedStringRegex := regexp.MustCompile(`^".*"$`)
	if quotedStringRegex.Match([]byte(value)) {
		return ast.StringExpression{Token: lexer.Token{Literal: value}}, nil
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
		return ast.BooleanExpression{Value: true, Token: lexer.Token{Literal: "true"}}, nil
	}
	if value == "false" {
		return ast.BooleanExpression{Value: false, Token: lexer.Token{Literal: "false"}}, nil
	}

	// Otherwise return content as string
	return ast.StringExpression{Token: lexer.Token{Literal: "\"" + value + "\""}}, nil
}

// WriteValue writes an AST expression to context with optional format specification
func WriteValue(name string, value ast.Expression, context map[string]string, format map[string]string) error {
	formatSpec, hasFormat := format[name]
	if !hasFormat {
		formatted, err := formatWithoutSpec(value)
		if err != nil {
			return err
		}
		context[name] = formatted
		return nil
	}

	if err := validateFormatString(formatSpec); err != nil {
		return fmt.Errorf("invalid format %q: %w", formatSpec, err)
	}

	// Format the value using the format specification
	formattedValue, err := formatWithSpec(value, formatSpec)
	if err != nil {
		return err
	}
	context[name] = formattedValue

	return nil
}

// formatWithSpec formats an AST expression using the specified format
func formatWithSpec(value ast.Expression, formatSpec string) (string, error) {
	switch expr := value.(type) {
	case ast.IntExpression:
		return fmt.Sprintf(formatSpec, expr.Value), nil
	case ast.FloatExpression:
		return fmt.Sprintf(formatSpec, expr.Value), nil
	case ast.StringExpression:
		// Extract content from quoted string literal
		content := expr.Token.Literal
		if len(content) >= 2 && content[0] == '"' && content[len(content)-1] == '"' {
			content = content[1 : len(content)-1]
		}
		return fmt.Sprintf(formatSpec, content), nil
	case ast.BooleanExpression:
		return fmt.Sprintf(formatSpec, expr.Value), nil
	default:
		return "", fmt.Errorf("unsupported expression type: %T", value)
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
		return expr.Token.Literal, nil
	case ast.BooleanExpression:
		return fmt.Sprintf("%t", expr.Value), nil
	default:
		return "", fmt.Errorf("unsupported expression type: %T", value)
	}
}

// newParseError creates an error message for format parsing failures
func newParseError(value, formatSpec string, err error) error {
	return fmt.Errorf("failed to parse %q with format %q: %w", value, formatSpec, err)
}

// detectPlaceholderType detects the type of scanf placeholder in the format string
func detectPlaceholderType(format string) string {
	// Regex patterns for different placeholder types
	intRegex := regexp.MustCompile(`%[#+ -]?(?:\*|\d+)?[diouxX]`)
	floatRegex := regexp.MustCompile(`%[#+ -]?(?:\*|\d+)?(?:\.(?:\*|\d+))?[eEfFgGaA]`)
	stringRegex := regexp.MustCompile(`%[#+ -]?(?:\*|\d+)?[sc]`)
	boolRegex := regexp.MustCompile(`%[#+ -]?(?:\*|\d+)?[t]`)

	if intRegex.MatchString(format) {
		return "int"
	}
	if floatRegex.MatchString(format) {
		return "float"
	}
	if stringRegex.MatchString(format) {
		return "string"
	}
	if boolRegex.MatchString(format) {
		return "bool"
	}

	// Unsupported format
	return "unsupported"
}

// validateFormatString checks if format contains exactly one scanf placeholder using regex
func validateFormatString(format string) error {
	// Regex to match scanf placeholders: % followed by optional flags/width/precision and a format specifier
	// This matches patterns like %d, %f, %s, %10s, %.2f, %t, etc.
	// But excludes escaped %% (literal %)
	placeholderRegex := regexp.MustCompile(`%(?:%|[#+ -]?(?:\*|\d+)?(?:\.(?:\*|\d+))?[diouxXeEfFgGaAcspvt])`)

	matches := placeholderRegex.FindAllString(format, -1)

	// Filter out escaped %% which are literal % characters
	actualPlaceholders := 0
	for _, match := range matches {
		if match != "%%" {
			actualPlaceholders++
		}
	}

	if actualPlaceholders == 0 {
		return fmt.Errorf("no scanf placeholder found")
	}
	if actualPlaceholders > 1 {
		return fmt.Errorf("multiple scanf placeholders found (%d), expected exactly one", actualPlaceholders)
	}
	return nil
}
