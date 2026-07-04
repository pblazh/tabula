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

func programCellIdentifiers(program ast.Program, bounds ast.RangeBounds) []string {
	var identifiers []string
	for _, statement := range program {
		identifiers = append(identifiers, statementCellIdentifiers(statement, bounds)...)
	}
	return identifiers
}

func statementCellIdentifiers(statement ast.Node, bounds ast.RangeBounds) []string {
	switch stmt := statement.(type) {
	case ast.LetStatement:
		return append(
			targetCellIdentifiers(statementTarget(stmt.Target, stmt.Identifier), bounds),
			expressionCellIdentifiers(stmt.Value, bounds)...,
		)
	case ast.FmtStatement:
		return append(
			targetCellIdentifiers(statementTarget(stmt.Target, stmt.Identifier), bounds),
			expressionCellIdentifiers(stmt.Value, bounds)...,
		)
	case ast.ExpressionStatement:
		return expressionCellIdentifiers(stmt.Value, bounds)
	default:
		return nil
	}
}

func targetCellIdentifiers(target ast.Node, bounds ast.RangeBounds) []string {
	switch t := target.(type) {
	case ast.IdentifierExpression:
		if ast.IsCellIdentifier(t.Value) {
			return []string{t.Value}
		}
	case ast.RangeExpression:
		return rangeCellIdentifiers(t, bounds)
	}
	return nil
}

func expressionCellIdentifiers(expression ast.Node, bounds ast.RangeBounds) []string {
	switch expr := expression.(type) {
	case ast.IdentifierExpression:
		if ast.IsCellIdentifier(expr.Value) {
			return []string{expr.Value}
		}
	case ast.PrefixExpression:
		return expressionCellIdentifiers(expr.Value, bounds)
	case ast.InfixExpression:
		return append(
			expressionCellIdentifiers(expr.Left, bounds),
			expressionCellIdentifiers(expr.Right, bounds)...,
		)
	case ast.CallExpression:
		var identifiers []string
		for _, arg := range expr.Arguments {
			identifiers = append(identifiers, expressionCellIdentifiers(arg, bounds)...)
		}
		return identifiers
	case ast.RangeExpression:
		return rangeCellIdentifiers(expr, bounds)
	}
	return nil
}

func rangeCellIdentifiers(expr ast.RangeExpression, bounds ast.RangeBounds) []string {
	expanded, err := expandRangeExpression(expr, bounds)
	if err != nil {
		return nil
	}
	return expanded.Value
}

func getProgramDimensions(identifiers []string) (int, int) {
	if len(identifiers) == 0 {
		return 0, 0
	}
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
