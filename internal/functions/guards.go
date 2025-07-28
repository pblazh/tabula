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

func MakeExactTypesGuard(format string, guards ...typeGuard) CallGuard {
	checkTypes := func(function ast.CallExpression, values ...ast.Expression) error {
		if len(values) != len(guards) {
			return ErrUnsupportedArity(format, function, len(guards), len(values))
		}

		for i, value := range values {
			if !guards[i](value) {
				return ErrUnsupportedArgument(format, function, values[i])
			}
		}

		return nil
	}

	return checkTypes
}

func MakeSameTypeGuard(format string, guard typeGuard) CallGuard {
	checkTypes := func(function ast.CallExpression, values ...ast.Expression) error {
		for _, value := range values {
			if !guard(value) {
				return ErrUnsupportedArgument(format, function, value)
			}
		}
		return nil
	}

	return checkTypes
}
