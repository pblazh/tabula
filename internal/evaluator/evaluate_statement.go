package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
)

func EvaluateStatement(statement ast.Statement, context map[string]string, input [][]string, formats map[string]string) error {
	switch s := statement.(type) {
	case ast.LetStatement:
		value, error := EvaluateExpression(s.Value, context, input, formats)
		if error != nil {
			return error
		}

		format := formats[s.Identifier.Token.Literal]
		output, error := WriteValue(value, format)
		if error != nil {
			return error
		}

		if ast.IsCellIdentifier(s.Identifier.Token.Literal) {
			context[s.Identifier.Token.Literal] = output
			break
		}

		context[s.Identifier.Token.Literal] = output
	case ast.FmtStatement:
		value, error := EvaluateExpression(s.Value, context, input, formats)
		if error != nil {
			return error
		}

		switch val := value.(type) {
		case ast.StringExpression:
			formats[s.Identifier.Token.Literal] = val.Token.Literal
		default:
			return ErrFmtExpectedString(s.Identifier.Token, val.String())
		}

	}
	return nil
}
