package functions

import "github.com/pblazh/csvss/internal/ast"

type CallGuard func(ast.CallExpression, ...ast.Expression) error

func EmptyGuard(function ast.CallExpression, values ...ast.Expression) error {
	return nil
}

func MakeArityGuard(arity int) CallGuard {
	checkArity := func(function ast.CallExpression, values ...ast.Expression) error {
		if arity >= 0 && len(values) != arity {
			return ErrUnsupportedArity(function, arity, len(values))
		}
		return nil
	}

	return checkArity
}

func SameTypeGuard(function ast.CallExpression, values ...ast.Expression) error {
	if len(values) == 0 {
		return nil
	}

	first := values[0]
	switch first.(type) {
	case ast.IntExpression:
		for _, arg := range values[1:] {
			switch a := arg.(type) {
			case ast.IntExpression:
				continue
			default:
				return ErrUnsupportedArgument(function, a)
			}
		}
	case ast.FloatExpression:
		for _, arg := range values[1:] {
			switch a := arg.(type) {
			case ast.FloatExpression:
				continue
			default:
				return ErrUnsupportedArgument(function, a)
			}
		}
	case ast.StringExpression:
		for _, arg := range values[1:] {
			switch a := arg.(type) {
			case ast.StringExpression:
				continue
			default:
				return ErrUnsupportedArgument(function, a)
			}
		}
	}

	return nil
}
