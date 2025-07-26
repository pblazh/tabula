package functions

import (
	"github.com/pblazh/csvss/internal/ast"
	"golang.org/x/exp/constraints"
)

func Sum(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	if len(values) == 0 {
		return ast.IntExpression{Value: 0}, nil
	}

	first := values[0]
	switch first.(type) {
	case ast.StringExpression:
		var args []string
		for _, arg := range values {
			switch a := arg.(type) {
			case ast.StringExpression:
				args = append(args, a.Value)
			default:
				return nil, ErrUnsupportedArgument(call, a)
			}
		}
		return ast.StringExpression{Value: sum(args...)}, nil
	default:
		return callNumbersFunction(sum, sum, EmptyGuard, call, values...)
	}
}

type Sumable interface {
	constraints.Integer | constraints.Float | string
}

func sum[T Sumable](values ...T) T {
	var result T
	for _, n := range values {
		result += n
	}
	return result
}
