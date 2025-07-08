package evaluator

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

func EvaluateStatement(statement ast.Statement, context map[string]string, format map[string]string) error {
	switch s := statement.(type) {
	case ast.LetStatement:
		value, error := EvaluateExpression(s.Value, context, format)
		if error != nil {
			return error
		}

		context[s.Identifier.Token.Literal] = value.String()
	case ast.FmtStatement:
		value, error := EvaluateExpression(s.Value, context, format)
		if error != nil {
			return error
		}

		switch val := value.(type) {
		case ast.StringExpression:
			format[s.Identifier.Token.Literal] = value.String()
		default:
			return fmt.Errorf("fmt %s accepts only strings, but got %s", s.Identifier.Token, val)
		}

	}
	return nil
}
