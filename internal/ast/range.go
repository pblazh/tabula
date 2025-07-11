package ast

import (
	"strconv"
)

func ExpandRange(start, end string) ([]string, error) {
	if !IsCellIdentifier(start) || !IsCellIdentifier(end) {
		return nil, ErrInvalidRange(start, end)
	}

	startCol, startRow := ParseCell(start)
	endCol, endRow := ParseCell(end)

	startColNum := ColumnToIndex(startCol)
	endColNum := ColumnToIndex(endCol)

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
			col := IndexToColumn(colNum)
			result = append(result, col+strconv.Itoa(row))
		}
	}

	return result, nil
}
