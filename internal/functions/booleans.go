package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func anyExpression(exp ast.Expression) bool {
	return true
}

func If(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsBoolean, anyExpression, anyExpression)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	a := values[0].(ast.BooleanExpression)
	if a.Value {
		return values[1], nil
	}
	return values[2], nil
}

func Not(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsBoolean)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	a := values[0].(ast.BooleanExpression)
	return ast.BooleanExpression{Value: !a.Value, Token: call.Token}, nil
}

func And(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsBoolean, ast.IsBoolean)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	a := values[0].(ast.BooleanExpression)
	b := values[1].(ast.BooleanExpression)
	return ast.BooleanExpression{Value: a.Value && b.Value, Token: call.Token}, nil
}

func Or(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsBoolean, ast.IsBoolean)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	a := values[0].(ast.BooleanExpression)
	b := values[1].(ast.BooleanExpression)
	return ast.BooleanExpression{Value: a.Value || b.Value, Token: call.Token}, nil
}

func True(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeArityGuard(format, 0)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	return ast.BooleanExpression{Value: true, Token: call.Token}, nil
}

func False(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeArityGuard(format, 0)
	if err := guard(call, values...); err != nil {
		return nil, err
	}
	return ast.BooleanExpression{Value: false, Token: call.Token}, nil
}
