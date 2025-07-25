package functions

import (
	"math"

	"github.com/pblazh/csvss/internal/ast"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type MathFunction[T Number] func(values ...T) T

func callMathFunction(
	intFunction MathFunction[int],
	floatFunction MathFunction[float64],
	callGuard CallGuard,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return ast.IntExpression{Value: int(intFunction())}, nil
	}

	first := values[0]
	switch first.(type) {
	case ast.IntExpression:
		var args []int
		for _, arg := range values {
			switch a := arg.(type) {
			case ast.IntExpression:
				args = append(args, a.Value)
			case ast.FloatExpression:
				args = append(args, int(a.Value))
			default:
				return nil, ErrUnsupportedArgument(call, a)
			}
		}
		return ast.IntExpression{Value: intFunction(args...)}, nil
	case ast.FloatExpression:
		var args []float64
		for _, arg := range values {
			switch a := arg.(type) {
			case ast.IntExpression:
				args = append(args, float64(a.Value))
			case ast.FloatExpression:
				args = append(args, a.Value)
			default:
				return nil, ErrUnsupportedArgument(call, a)
			}
		}
		return ast.FloatExpression{Value: floatFunction(args...)}, nil
	default:
		return nil, ErrUnsupportedFunction(call)
	}
}

func product[T Number](values ...T) T {
	result := T(1)
	for _, n := range values {
		result *= n
	}
	return result
}

func average[T Number](values ...T) T {
	if len(values) == 0 {
		return T(0)
	}

	total := T(0)
	for _, n := range values {
		total += n
	}
	return T(total / T(len(values)))
}

func abs[T Number](values ...T) T {
	return T(math.Abs(float64(values[0])))
}
