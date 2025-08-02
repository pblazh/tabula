package ast

import (
	"fmt"
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

// ParseCell parses a cell reference like "A1" into a zero based column and row components
func ParseCell(cell string) (int, int) {
	matches := cellParseRegex.FindStringSubmatch(cell)
	if len(matches) != 3 {
		return -1, -1
	}

	columnName := strings.ToUpper(matches[1])
	row, _ := strconv.Atoi(matches[2])

	letters := int('Z'-'A') + 1
	column := 0
	for _, c := range columnName {
		column = column*letters + int(c-'A') + 1
	}
	return column - 1, row - 1
}

// ToCell converts 0-based column, row to cell like 0, 0 to "A1", 1,1 to "B2", etc.
func ToCell(column, row int) string {
	letters := int('Z'-'A') + 1
	column++
	row++
	var result string
	for column > 0 {
		column--
		result = string(rune('A'+column%letters)) + result
		column = column / letters
	}
	return fmt.Sprintf("%s%d", result, row)
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

func ToInt(expr *Expression) (*IntExpression, bool) {
	switch e := (*expr).(type) {
	case IntExpression:
		return &e, true
	case FloatExpression:
		return &IntExpression{Value: int(e.Value), Token: e.Token}, true
	default:
		return nil, false
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

func ToFloat(expr *Expression) (*FloatExpression, bool) {
	switch e := (*expr).(type) {
	case IntExpression:
		return &FloatExpression{Value: float64(e.Value), Token: e.Token}, true
	case FloatExpression:
		return &e, true
	default:
		return nil, false
	}
}

// IsNumeric returns true if the expression is a numeric type (int or float)
func IsNumeric(expr Expression) bool {
	return IsInt(expr) || IsFloat(expr)
}

func IsIdentifier(expr Expression) bool {
	switch expr.(type) {
	case IdentifierExpression:
		return true
	default:
		return false
	}
}

func IsFunction(expr Expression) bool {
	switch expr.(type) {
	case CallExpression:
		return true
	default:
		return false
	}
}

func IsString(expr Expression) bool {
	switch expr.(type) {
	case StringExpression:
		return true
	default:
		return false
	}
}

func ToString(expr *Expression) (*StringExpression, bool) {
	switch e := (*expr).(type) {
	case StringExpression:
		return &e, true
	default:
		return nil, false
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

func ToBoolean(expr *Expression) (*BooleanExpression, bool) {
	switch e := (*expr).(type) {
	case BooleanExpression:
		return &e, true
	default:
		return nil, false
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
