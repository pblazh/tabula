package functions

import (
	"math"

	"github.com/pblazh/csvss/internal/ast"
)

func Int(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(ast.IsNumeric)

	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	switch v := values[0].(type) {
	case ast.IntExpression:
		return v, nil
	case ast.FloatExpression:
		return ast.IntExpression{Value: int(v.Value)}, nil
	default:
		return nil, ErrUnsupportedFunction(call)
	}
}

func Round(up bool, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(ast.IsNumeric, ast.IsNumeric)

	if err := callGuard(call, values...); err != nil {
		return nil, err
	}
	switch v := values[0].(type) {
	case ast.IntExpression:
		switch p := values[1].(type) {
		case ast.IntExpression:
			return ast.FloatExpression{Value: roundPrecise(up, float64(v.Value), float64(p.Value))}, nil
		case ast.FloatExpression:
			return ast.FloatExpression{Value: (roundPrecise(up, float64(v.Value), p.Value))}, nil
		}
	case ast.FloatExpression:
		switch p := values[1].(type) {
		case ast.IntExpression:
			return ast.FloatExpression{Value: roundPrecise(up, v.Value, float64(p.Value))}, nil
		case ast.FloatExpression:
			return ast.FloatExpression{Value: (roundPrecise(up, v.Value, p.Value))}, nil
		}
	}

	return nil, nil
}

func roundPrecise(up bool, value, precision float64) float64 {
	if up {
		return math.Ceil(value/precision) * precision
	}

	return math.Floor(value/precision) * precision
}
