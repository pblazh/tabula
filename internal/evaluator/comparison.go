package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func evaluateComparison(operator lexer.Token, left, right ast.Expression, intOp func(int, int) (ast.Expression, error), floatOp func(float64, float64) (ast.Expression, error), stringOp func(string, string) (ast.Expression, error), boolOp func(bool, bool) (ast.Expression, error)) (ast.Expression, error) {
	if ast.IsNumeric(left) && ast.IsNumeric(right) {
		return evaluateNumericOperation(left, right, operator, intOp, floatOp)
	}

	if ast.IsString(left) && ast.IsString(right) {
		l := left.(ast.StringExpression)
		r := right.(ast.StringExpression)
		return stringOp(l.Value, r.Value)
	}

	if ast.IsBoolean(left) && ast.IsBoolean(right) {
		l := left.(ast.BooleanExpression)
		r := right.(ast.BooleanExpression)
		return boolOp(l.Value, r.Value)
	}

	return nil, ErrUnsupportedBinaryOperation(operator, left, right)
}

func evaluateEquality(operator lexer.Token, left, right ast.Expression) (ast.Expression, error) {
	return evaluateComparison(operator, left, right,
		func(a, b int) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a == b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a == b, Token: operator}, nil
		},
		func(a, b string) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a == b, Token: operator}, nil
		},
		func(a, b bool) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a == b, Token: operator}, nil
		})
}

func evaluateInequality(operator lexer.Token, left, right ast.Expression) (ast.Expression, error) {
	return evaluateComparison(operator, left, right,
		func(a, b int) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a != b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a != b, Token: operator}, nil
		},
		func(a, b string) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a != b, Token: operator}, nil
		},
		func(a, b bool) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a != b, Token: operator}, nil
		})
}

func evaluateLessThan(operator lexer.Token, left, right ast.Expression) (ast.Expression, error) {
	return evaluateComparison(operator, left, right,
		func(a, b int) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a < b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a < b, Token: operator}, nil
		},
		func(a, b string) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a < b, Token: operator}, nil
		},
		func(a, b bool) (ast.Expression, error) {
			return nil, ErrUnsupportedBinaryOperation(operator, ast.BooleanExpression{}, ast.BooleanExpression{})
		})
}

func evaluateGreaterThan(operator lexer.Token, left, right ast.Expression) (ast.Expression, error) {
	return evaluateComparison(operator, left, right,
		func(a, b int) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a > b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a > b, Token: operator}, nil
		},
		func(a, b string) (ast.Expression, error) {
			return ast.BooleanExpression{Value: a > b, Token: operator}, nil
		},
		func(a, b bool) (ast.Expression, error) {
			return nil, ErrUnsupportedBinaryOperation(operator, ast.BooleanExpression{}, ast.BooleanExpression{})
		})
}
