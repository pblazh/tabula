package ast

type RangeBounds struct {
	width  int
	height int
}

func NewRangeBounds(records [][]string) RangeBounds {
	bounds := RangeBounds{height: len(records)}
	if len(records) > 0 {
		bounds.width = len(records[0])
	}

	return bounds
}

func (bounds RangeBounds) LastCellInColumn(columnName string) (string, bool) {
	column := ParseColumn(columnName)
	if column < 0 || column >= bounds.width || bounds.height == 0 {
		return "", false
	}

	return ToCell(column, bounds.height-1), true
}

func (bounds RangeBounds) LastCellInRow(rowName string) (string, bool) {
	row := ParseRow(rowName)
	if row < 0 || row >= bounds.height || bounds.width == 0 {
		return "", false
	}

	return ToCell(bounds.width-1, row), true
}

func (bounds RangeBounds) FirstCellInLastRow() (string, bool) {
	if bounds.height == 0 {
		return "", false
	}

	return ToCell(0, bounds.height-1), true
}

func (bounds RangeBounds) LastCell() (string, bool) {
	if bounds.width == 0 || bounds.height == 0 {
		return "", false
	}

	return ToCell(bounds.width-1, bounds.height-1), true
}

func (bounds RangeBounds) ContainsCell(cell string) bool {
	column, row := ParseCell(cell)
	return column >= 0 && row >= 0 && column < bounds.width && row < bounds.height
}

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

func IsValidRangeSyntax(start, end string) bool {
	return IsCellIdentifier(start) && IsCellIdentifier(end)
}

func IsValidOpenRangeSyntax(start, end string) bool {
	if IsValidRangeSyntax(start, end) {
		return true
	}

	if start == "" {
		return IsCellIdentifier(end)
	}
	if !IsCellIdentifier(start) {
		return false
	}
	if end == "" {
		return true
	}
	return IsColumnIdentifier(end) || IsRowIdentifier(end)
}

func ExpandRangeWithBounds(start, end string, bounds RangeBounds) ([]string, error) {
	resolvedStart, resolvedEnd, ok := bounds.resolveRange(start, end)
	if !ok {
		return nil, ErrInvalidRange(start, end)
	}

	cells, err := ExpandRange(resolvedStart, resolvedEnd)
	if err != nil {
		return nil, err
	}

	return cells, nil
}

func (bounds RangeBounds) resolveRange(start, end string) (string, string, bool) {
	switch {
	case IsValidRangeSyntax(start, end):
		return start, end, true
	case start == "" && IsCellIdentifier(end) && bounds.ContainsCell(end):
		resolvedStart, ok := bounds.FirstCellInLastRow()
		return resolvedStart, end, ok
	case IsCellIdentifier(start) && bounds.ContainsCell(start):
		resolvedEnd, ok := bounds.resolveOpenEnd(end)
		return start, resolvedEnd, ok
	default:
		return "", "", false
	}
}

func (bounds RangeBounds) resolveOpenEnd(end string) (string, bool) {
	switch {
	case end == "":
		return bounds.LastCell()
	case IsColumnIdentifier(end):
		return bounds.LastCellInColumn(end)
	case IsRowIdentifier(end):
		return bounds.LastCellInRow(end)
	default:
		return "", false
	}
}
