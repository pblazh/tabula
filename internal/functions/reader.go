package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

// ParseWithoutFormat parses a value without format specification using default rules
func ParseWithoutFormat(value string) (ast.Expression, error) {
	// Trim whitespace as the first step
	value = strings.TrimSpace(value)

	// Check if it's a string (enclosed in quotes) - return content as string
	quotedStringRegex := regexp.MustCompile(`^".*"$`)
	if quotedStringRegex.Match([]byte(value)) {
		return ast.StringExpression{Value: value, Token: lexer.Token{Literal: value}}, nil
	}

	datetime, err := ParseDateWithoutFormat(value)
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

func ReadValue(value string, format string) (ast.Expression, error) {
	if format == "" {
		return ParseWithoutFormat(value)
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
		return parseFormatedDate(value, format)
	}
}

const (
	intPlacehoder = iota
	floatPlacehoder
	stringPlacehoder
	boolPlacehoder
	datePlacehoder
)

// detectPlaceholderType detects the type of scanf placeholder in the format string
func detectPlaceholderType(format string) int {
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[diouxX]`, format); matched {
		return intPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?(?:\.(?:\*|\d+))?[eEfFgGaA]`, format); matched {
		return floatPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[sc]`, format); matched {
		return stringPlacehoder
	}
	if matched, _ := regexp.MatchString(`%[#+ -]?(?:\*|\d+)?[t]`, format); matched {
		return boolPlacehoder
	}
	return datePlacehoder
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

func parseFormatedDate(value, format string) (ast.DateExpression, error) {
	date, err := time.Parse(format, value)
	if err != nil {
		return ast.DateExpression{}, err
	}

	return ast.DateExpression{Value: date, Token: lexer.Token{Literal: value}}, nil
}
