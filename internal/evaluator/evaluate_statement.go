package evaluator

import (
	"fmt"

	"github.com/pblazh/tabula/internal/ast"
	functions "github.com/pblazh/tabula/internal/core"
)

func EvaluateStatement(statement ast.Statement, context map[string]string, input [][]string, formats map[string]string) error {
	switch s := statement.(type) {
	case ast.LetStatement:

		value, error := EvaluateExpression(s.Value, context, input, formats, s.Identifier.Value)
		if error != nil {
			return error
		}

		format := formats[s.Identifier.Value]
		output, error := functions.WriteValue(value, format)
		if error != nil {
			return error
		}

		if ast.IsCellIdentifier(s.Identifier.Value) {
			col, row := ast.ParseCell(s.Identifier.Value)
			err := ifCellInBounds(s.Identifier, input)
			if err != nil {
				return err
			}
			input[row][col] = output
			break
		}

		context[s.Identifier.Value] = output
	case ast.FmtStatement:
		value, error := EvaluateExpression(s.Value, context, input, formats, s.Identifier.Value)
		if error != nil {
			return error
		}

		switch val := value.(type) {
		case ast.StringExpression:
			formats[s.Identifier.Value] = val.Value
		default:
			return ErrFmtExpectedString(s.Identifier.Token, val.String())
		}

	case ast.IncludeStatement:
		// Includes should already be resolved during parsing
		return fmt.Errorf("internal error: IncludeStatement should not reach evaluator")
	}
	return nil
}
