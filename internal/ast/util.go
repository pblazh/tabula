package ast

import (
	"regexp"
	"strconv"
	"strings"
)

// Pre-compiled regex patterns for better performance
var (
	cellIdentifierRegex = regexp.MustCompile(`^[A-Za-z]+[0-9]+$`)
	cellParseRegex      = regexp.MustCompile(`^([A-Za-z]+)([0-9]+)$`)
)

// IsCellIdentifier returns true if the identifier matches A1 cell name format
func IsCellIdentifier(identifier string) bool {
	return cellIdentifierRegex.MatchString(identifier)
}

// ParseCell parses a cell reference like "A1" into column and row components
func ParseCell(cell string) (string, int) {
	matches := cellParseRegex.FindStringSubmatch(cell)
	if len(matches) != 3 {
		return "", 0
	}

	col := strings.ToUpper(matches[1]) // Convert column to uppercase for consistency
	row, _ := strconv.Atoi(matches[2])
	return col, row
}

// ColumnToIndex converts column letters like "A" to 0, "B" to 1, etc. (0-based)
func ColumnToIndex(column string) int {
	result := 0
	for _, c := range column {
		result = result*26 + int(c-'A') + 1
	}
	return result
}

// IndexToColumn converts 0-based index to column letters like 0 to "A", 1 to "B", etc.
func IndexToColumn(index int) string {
	var result string
	for index > 0 {
		index--
		result = string(rune('A'+index%26)) + result
		index = index / 26
	}
	return result
}

// TypeName returns a human-readable name for the expression type.
func TypeName(expr Expression) string {
	switch expr.(type) {
	case IntExpression:
		return "integer"
	case FloatExpression:
		return "float"
	case BooleanExpression:
		return "boolean"
	case StringExpression:
		return "string"
	case IdentifierExpression:
		return "identifier"
	case PrefixExpression:
		return "prefix expression"
	case InfixExpression:
		return "infix expression"
	case CallExpression:
		return "function call"
	case RangeExpression:
		return "range"
	default:
		return "unknown"
	}
}

func IsInt(expr Expression) bool {
	switch expr.(type) {
	case IntExpression:
		return true
	default:
		return false
	}
}

func IsFloat(expr Expression) bool {
	switch expr.(type) {
	case FloatExpression:
		return true
	default:
		return false
	}
}

// IsNumeric returns true if the expression is a numeric type (int or float)
func IsNumeric(expr Expression) bool {
	return IsInt(expr) || IsFloat(expr)
}

func IsString(expr Expression) bool {
	switch expr.(type) {
	case StringExpression:
		return true
	default:
		return false
	}
}

func IsBoolean(expr Expression) bool {
	switch expr.(type) {
	case BooleanExpression:
		return true
	default:
		return false
	}
}

// IsLiteral returns true if the expression is a literal value type
func IsLiteral(expr Expression) bool {
	switch expr.(type) {
	case IntExpression, FloatExpression, BooleanExpression, StringExpression:
		return true
	default:
		return false
	}
}
