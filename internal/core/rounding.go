package core

import (
	"math"

	"github.com/pblazh/tabula/internal/ast"
)

func Int(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard("INT(number)", ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	result, _ := ast.ToInt(&(values[0]))
	return result, nil
}

func Round(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	first, second, err := extractRoundingArguments(format, call, values)
	if err != nil {
		return nil, err
	}

	if math.Round(second) == second {
		return ast.IntExpression{Value: int(roundPrecise(first, second))}, nil
	}

	return ast.FloatExpression{Value: roundPrecise(first, second)}, nil
}

func RoundUp(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	first, second, err := extractRoundingArguments(format, call, values)
	if err != nil {
		return nil, err
	}

	if math.Round(second) == second {
		return ast.IntExpression{Value: int(roundUpPrecise(first, second))}, nil
	}
	return ast.FloatExpression{Value: roundUpPrecise(first, second)}, nil
}

func RoundDown(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	first, second, err := extractRoundingArguments(format, call, values)
	if err != nil {
		return nil, err
	}
	if math.Round(second) == second {
		return ast.IntExpression{Value: int(roundDownPrecise(first, second))}, nil
	}
	return ast.FloatExpression{Value: roundDownPrecise(first, second)}, nil
}

func extractRoundingArguments(format string, call ast.CallExpression, values []ast.Expression) (float64, float64, error) {
	var callGuard CallGuard

	if len(values) == 1 {
		callGuard = MakeExactTypesGuard(format, ast.IsNumeric)
	} else {
		callGuard = MakeExactTypesGuard(format, ast.IsNumeric, ast.IsNumeric)
	}

	if err := callGuard(call, values...); err != nil {
		return 0, 0, err
	}

	first, _ := ast.ToFloat(&(values[0]))
	if len(values) == 1 {
		return first.Value, 1, nil
	}

	second, _ := ast.ToFloat(&(values[1]))
	return first.Value, second.Value, nil
}
