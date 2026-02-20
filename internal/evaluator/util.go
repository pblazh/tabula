package evaluator

import (
	"strings"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func evaluateNumericOperation(
	left, right ast.Node,
	operator lexer.Token,
	intOp func(int, int) (ast.Node, error),
	floatOp func(float64, float64) (ast.Node, error),
) (ast.Node, error) {
	switch l := left.(type) {
	case ast.IntExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return intOp(l.Value, r.Value)
		case ast.FloatExpression:
			return floatOp(float64(l.Value), r.Value)
		}
	case ast.FloatExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return floatOp(l.Value, float64(r.Value))
		case ast.FloatExpression:
			return floatOp(l.Value, r.Value)
		}
	}
	return nil, ErrUnsupportedBinaryOperation(operator, left, right)
}

func ifCellInBounds(s ast.IdentifierExpression, input [][]string) error {
	col, row := ast.ParseCell(s.Value)
	if row < 0 || row >= len(input) || col < 0 || col >= len(input[row]) {
		return ErrRelOutOfBounds(s)
	}
	return nil
}

func EnsureProgramDimensions(identifiers []string, records [][]string) [][]string {
	for i, row := range records {
		for j, cel := range row {
			records[i][j] = strings.TrimSpace(cel)
		}
	}

	requiredWidth, requiredHeight := getProgramDimensions(identifiers)

	for i, row := range records {
		diff := requiredWidth - len(row)
		if diff > 0 {
			records[i] = append(records[i], make([]string, requiredWidth-len(row))...)
		}
	}
	for range requiredHeight - len(records) {
		records = append(records, make([]string, requiredWidth))
	}
	return records
}

func getProgramDimensions(identifiers []string) (int, int) {
	requiredWidth := 0
	requiredHeight := 0

	for _, id := range identifiers {
		if ast.IsCellIdentifier(id) {
			col, row := ast.ParseCell(id)

			if col > requiredWidth {
				requiredWidth = col
			}
			if row > requiredHeight {
				requiredHeight = row
			}
		}
	}

	return requiredWidth + 1, requiredHeight + 1
}
