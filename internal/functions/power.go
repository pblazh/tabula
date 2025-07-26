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
