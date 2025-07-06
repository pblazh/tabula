package ast

import (
	"fmt"
	"strconv"
)

func ExpandRange(start, end string) ([]string, error) {
	if !IsCellIdentifier(start) || !IsCellIdentifier(end) {
		return nil, fmt.Errorf("range must contain valid cell references (like A1:B2), got %s:%s", start, end)
	}

	startCol, startRow := parseCell(start)
	endCol, endRow := parseCell(end)

	startColNum := ColumnToNumber(startCol)
	endColNum := ColumnToNumber(endCol)

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
			col := NumberToColumn(colNum)
			result = append(result, col+strconv.Itoa(row))
		}
	}

	return result, nil
}
