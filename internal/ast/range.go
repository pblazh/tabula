package ast

import (
	"fmt"
	"strconv"
	"unicode"
)

// ParseCell parses a cell reference like "A1" into column and row components
func ParseCell(cell string) (string, int, bool) {
	if len(cell) == 0 {
		return "", 0, false
	}

	// Find where letters end and numbers begin
	i := 0
	for i < len(cell) && unicode.IsLetter(rune(cell[i])) {
		i++
	}

	if i == 0 || i == len(cell) {
		return "", 0, false
	}

	col := cell[:i]
	rowStr := cell[i:]

	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return "", 0, false
	}

	return col, row, true
}

// columnToNumber converts column letters like "A" to 1, "B" to 2, etc.
func columnToNumber(column string) int {
	result := 0
	for _, c := range column {
		result = result*26 + int(c-'A') + 1
	}
	return result
}

// numberToColumn converts numbers like 1 to "A", 2 to "B", etc.
func numberToColumn(num int) string {
	var result string
	for num > 0 {
		num-- // Convert to 0-based
		result = string(rune('A'+num%26)) + result
		num /= 26
	}
	return result
}

func ExpandRange(start, end string) ([]string, error) {
	startCol, startRow, startOk := ParseCell(start)
	endCol, endRow, endOk := ParseCell(end)

	if !startOk || !endOk {
		return nil, fmt.Errorf("range must contain valid cell references (like A1:B2), got %s:%s", start, end)
	}

	startColNum := columnToNumber(startCol)
	endColNum := columnToNumber(endCol)

	colStep := 1
	if startColNum > endColNum {
		colStep = -1
	}

	rowStep := 1
	if startRow > endRow {
		rowStep = -1
	}

	var result []string
	for row := startRow; (rowStep > 0 && row <= endRow) || (rowStep < 0 && row >= endRow); row += rowStep {
		for colNum := startColNum; (colStep > 0 && colNum <= endColNum) || (colStep < 0 && colNum >= endColNum); colNum += colStep {
			col := numberToColumn(colNum)
			result = append(result, col+strconv.Itoa(row))
		}
	}

	return result, nil
}
