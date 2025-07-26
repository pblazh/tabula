package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func Concat(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	if len(values) == 0 {
		return ast.StringExpression{Value: ""}, nil
	}

	args := make([]string, len(values))
	for i, arg := range values {
		switch a := arg.(type) {
		case ast.StringExpression:
			args[i] = a.Value
		default:
			return nil, ErrUnsupportedArgument(format, call, a)
		}
	}
	return ast.StringExpression{Value: concat(args...)}, nil
}

func concat[T string](values ...T) T {
	var result T
	for _, n := range values {
		result += n
	}
	return result
}
