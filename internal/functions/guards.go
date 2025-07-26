package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

type CallGuard func(ast.CallExpression, ...ast.Expression) error

func EmptyGuard(function ast.CallExpression, values ...ast.Expression) error {
	return nil
}

func NumericGuard(function ast.CallExpression, values ...ast.Expression) error {
	for _, t := range values {
		if !ast.IsNumeric(t) {
			return ErrUnsupportedFunction(function)
		}
	}
	return nil
}

func MakeArityGuard(format string, arity int) CallGuard {
	checkArity := func(function ast.CallExpression, values ...ast.Expression) error {
		if arity >= 0 && len(values) != arity {
			return ErrUnsupportedArity(format, function, arity, len(values))
		}
		return nil
	}

	return checkArity
}

type typeGuard func(expr ast.Expression) bool

func MakeExactTypesGuard(format string, types ...typeGuard) CallGuard {
	checkTypes := func(function ast.CallExpression, values ...ast.Expression) error {
		if len(values) != len(types) {
			return ErrUnsupportedArity(format, function, len(types), len(values))
		}

		for i, t := range types {
			if !t(values[i]) {
				return ErrUnsupportedFunction(function)
			}
		}

		return nil
	}

	return checkTypes
}
