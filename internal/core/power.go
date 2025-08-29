package core

import (
	"math"

	"github.com/pblazh/tabula/internal/ast"
)

func Power(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsNumeric, ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	first, _ := ast.ToFloat(&(values[0]))
	second, _ := ast.ToFloat(&(values[1]))

	return ast.FloatExpression{Value: math.Pow(first.Value, second.Value)}, nil
}

func Sqrt(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsNumeric)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	value := values[0]
	switch v := value.(type) {
	case ast.IntExpression:
		if v.Value < 0 {
			return nil, ErrUnsupportedArgument(format, call, value)
		}
	case ast.FloatExpression:
		if v.Value < 0 {
			return nil, ErrUnsupportedArgument(format, call, value)
		}
	}

	return callNumbersFunction(format, sqrt, sqrt, EmptyGuard, call, values...)
}
