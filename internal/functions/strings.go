package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func Concat(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeSameTypeGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	args := make([]string, len(values))
	for i, arg := range values {
		a := arg.(ast.StringExpression)
		args[i] = a.Value
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
