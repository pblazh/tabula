package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func IsNumber(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeArityGuard(format, 1)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	return ast.BooleanExpression{Value: ast.IsNumeric(values[0]), Token: call.Token}, nil
}

func IsText(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeArityGuard(format, 1)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	return ast.BooleanExpression{Value: ast.IsString(values[0]), Token: call.Token}, nil
}

func IsBoolean(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeArityGuard(format, 1)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	return ast.BooleanExpression{Value: ast.IsBoolean(values[0]), Token: call.Token}, nil
}
