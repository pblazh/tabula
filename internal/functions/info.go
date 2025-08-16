package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func IsNumber(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard("ISNUMBER(number)", ast.IsString)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	result, err := ParseWithoutFormat(values[0].(ast.StringExpression).Value)
	if err != nil {
		return nil, err
	}

	return ast.BooleanExpression{Value: ast.IsNumeric(result), Token: call.Token}, nil
}
