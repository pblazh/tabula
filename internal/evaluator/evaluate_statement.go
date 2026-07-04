package evaluator

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
	functions "github.com/pblazh/tabula/internal/core"
)

func EvaluateStatement(
	statement ast.Node,
	context map[string]string,
	input [][]string,
	formats map[string]string,
) error {
	bounds := ast.NewRangeBounds(input)
	switch s := statement.(type) {
	case ast.LetStatement:
		targets, err := evaluateStatementTargets(statementTarget(s.Target, s.Identifier), bounds)
		if err != nil {
			return err
		}

		for _, target := range targets {
			value, err := EvaluateExpression(s.Value, context, input, formats, target.Value)
			if err != nil {
				return err
			}

			format := formats[target.Value]
			output, err := functions.WriteValue(value, format)
			if err != nil {
				return ErrWriting(err)
			}

			if ast.IsCellIdentifier(target.Value) {
				col, row := ast.ParseCell(target.Value)
				err := ifCellInBounds(target, input)
				if err != nil {
					return err
				}
				input[row][col] = output
				continue
			}

			context[target.Value] = output
		}

	case ast.FmtStatement:
		targets, err := evaluateStatementTargets(statementTarget(s.Target, s.Identifier), bounds)
		if err != nil {
			return err
		}

		for _, target := range targets {
			value, err := EvaluateExpression(s.Value, context, input, formats, target.Value)
			if err != nil {
				return err
			}

			switch val := value.(type) {
			case ast.StringExpression:
				formats[target.Value] = val.Value
			default:
				return ErrFmtExpectedString(target.Token, val.String())
			}
		}

	case ast.IncludeStatement:
		// Includes should already be resolved during parsing
		return fmt.Errorf("internal error: IncludeStatement should not reach evaluator")
	}
	return nil
}

func statementTarget(target ast.Node, identifier ast.IdentifierExpression) ast.Node {
	if target != nil {
		return target
	}
	return identifier
}

func evaluateStatementTargets(
	target ast.Node,
	rangeBounds ast.RangeBounds,
) ([]ast.IdentifierExpression, error) {
	switch t := target.(type) {
	case ast.IdentifierExpression:
		return []ast.IdentifierExpression{t}, nil
	case ast.RangeExpression:
		expanded, err := expandRangeExpression(t, rangeBounds)
		if err != nil {
			return nil, err
		}
		targets := make([]ast.IdentifierExpression, 0, len(expanded.Value))
		for _, cell := range expanded.Value {
			targets = append(targets, ast.IdentifierExpression{Token: t.Token, Value: cell})
		}
		return targets, nil
	default:
		return nil, fmt.Errorf("unsupported statement target %s", target)
	}
}
