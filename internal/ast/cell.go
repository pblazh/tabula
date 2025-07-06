package ast

import (
	"regexp"
	"strconv"
	"strings"
)

// parseCell parses a cell reference like "A1" into column and row components
func parseCell(cell string) (string, int, bool) {
	cellRegex := regexp.MustCompile(`^([A-Za-z]+)([0-9]+)$`)
	matches := cellRegex.FindStringSubmatch(cell)

	if len(matches) != 3 {
		return "", 0, false
	}

	col := strings.ToUpper(matches[1]) // Convert column to uppercase for consistency
	rowStr := matches[2]

	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return "", 0, false
	}

	return col, row, true
}
