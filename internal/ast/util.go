package ast

import (
	"regexp"
	"strconv"
	"strings"
)

// IsCellIdentifier returns true if the identifier matches A1 cell name format
func IsCellIdentifier(identifier string) bool {
	return regexp.MustCompile(`^[A-Za-z]+[0-9]+$`).MatchString(identifier)
}

// ColumnToNumber converts column letters like "A" to 1, "B" to 2, etc.
func ColumnToNumber(column string) int {
	result := 0
	for _, c := range column {
		result = result*26 + int(c-'A') + 1
	}
	return result
}

// NumberToColumn converts numbers like 1 to "A", 2 to "B", etc.
func NumberToColumn(num int) string {
	var result string
	for num > 0 {
		num-- // Convert to 0-based
		result = string(rune('A'+num%26)) + result
		num /= 26
	}
	return result
}

// parseCell parses a cell reference like "A1" into column and row components
func parseCell(cell string) (string, int) {
	cellRegex := regexp.MustCompile(`^([A-Za-z]+)([0-9]+)$`)
	matches := cellRegex.FindStringSubmatch(cell)

	col := strings.ToUpper(matches[1]) // Convert column to uppercase for consistency
	rowStr := matches[2]

	row, _ := strconv.Atoi(rowStr)
	return col, row
}
