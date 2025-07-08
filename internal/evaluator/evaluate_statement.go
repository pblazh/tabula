package evaluator

import "github.com/pblazh/csvss/internal/ast"

func EvaluateStatement(statement ast.Statement, context map[string]string) error {
	switch s := statement.(type) {
	case ast.LetStatement:
		value, error := EvaluateExpression(s.Value, context)
		if error != nil {
			return error
		}

		context[s.Identifier.Token.Literal] = value.String()
	}
	return nil
}
