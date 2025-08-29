package core

import (
	"github.com/pblazh/tabula/internal/ast"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type MathFunction[T Number] func(values ...T) T

func callNumbersFunction(
	format string,
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

	// Check if any argument is a float to determine result type
	hasFloat := false
	for _, arg := range values {
		switch arg.(type) {
		case ast.FloatExpression:
			hasFloat = true
		case ast.IntExpression:
			// continue checking
		default:
			return nil, ErrUnsupportedArgument(format, call, arg)
		}
	}

	if hasFloat {
		// If any argument is float, use float processing
		var args []float64
		for _, arg := range values {
			switch a := arg.(type) {
			case ast.IntExpression:
				args = append(args, float64(a.Value))
			case ast.FloatExpression:
				args = append(args, a.Value)
			}
		}
		return ast.FloatExpression{Value: floatFunction(args...)}, nil
	} else {
		// All arguments are integers, use int processing
		var args []int
		for _, arg := range values {
			switch a := arg.(type) {
			case ast.IntExpression:
				args = append(args, a.Value)
			}
		}
		return ast.IntExpression{Value: intFunction(args...)}, nil
	}
}
