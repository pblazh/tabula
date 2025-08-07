package functions

import (
	"time"

	"github.com/pblazh/csvss/internal/ast"
)

func ToDate(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsString, ast.IsString)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	layout := values[0].(ast.StringExpression)
	value := values[1].(ast.StringExpression)

	parsed, err := time.Parse(layout.Value, value.Value)
	if err != nil {
		return nil, ErrExecuting(format, call, err)
	}

	return ast.DateExpression{Value: parsed, Token: call.Token}, nil
}
