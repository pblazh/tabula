package functions

import (
	"math"

	"github.com/pblazh/csvss/internal/ast"
)

func Power(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard("POWER(number, number)", ast.IsNumeric, ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	first, _ := ast.ToFloat(&(values[0]))
	second, _ := ast.ToFloat(&(values[1]))

	return ast.FloatExpression{Value: math.Pow(first.Value, second.Value)}, nil
}

func Add(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard("ADD(number, number)", ast.IsNumeric, ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	first, _ := ast.ToFloat(&(values[0]))
	second, _ := ast.ToFloat(&(values[1]))

	result := first.Value + second.Value

	if ast.IsInt(first) || ast.IsInt(second) {
		return ast.IntExpression{Value: int(result)}, nil
	}

	return ast.FloatExpression{Value: result}, nil
}
