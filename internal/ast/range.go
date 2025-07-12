package ast

func ExpandRange(start, end string) ([]string, error) {
	if !IsCellIdentifier(start) || !IsCellIdentifier(end) {
		return nil, ErrInvalidRange(start, end)
	}

	startCol, startRow := ParseCell(start)
	endCol, endRow := ParseCell(end)

	colStep := 1
	if startCol > endCol {
		colStep = -1
	}

	rowStep := 1
	if startRow > endRow {
		rowStep = -1
	}

	var result []string
	for row := startRow; (rowStep > 0 && row <= endRow) || (rowStep < 0 && row >= endRow); row += rowStep {
		for colNum := startCol; (colStep > 0 && colNum <= endCol) || (colStep < 0 && colNum >= endCol); colNum += colStep {
			result = append(result, ToCell(colNum, row))
		}
	}

	return result, nil
}
