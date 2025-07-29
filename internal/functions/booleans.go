package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func False(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsBoolean)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	a := values[0].(ast.BooleanExpression)
	return ast.BooleanExpression{Value: !a.Value, Token: call.Token}, nil
}
