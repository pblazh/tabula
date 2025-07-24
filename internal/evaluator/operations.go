package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func evaluateNegation(expr ast.Expression, token lexer.Token) (ast.Expression, error) {
	switch r := expr.(type) {
	case ast.IntExpression:
		return ast.IntExpression{Value: -r.Value, Token: token}, nil
	case ast.FloatExpression:
		return ast.FloatExpression{Value: -r.Value, Token: token}, nil
	default:
		return nil, ErrUnsupportedOperation(token, expr)
	}
}

func evaluateNot(expr ast.Expression, token lexer.Token) (ast.Expression, error) {
	switch r := expr.(type) {
	case ast.BooleanExpression:
		return ast.BooleanExpression{Value: !r.Value, Token: token}, nil
	default:
		return nil, ErrUnsupportedOperation(token, expr)
	}
}

func evaluateAddition(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	if l, ok := left.(ast.StringExpression); ok {
		if r, ok := right.(ast.StringExpression); ok {
			return ast.StringExpression{Value: l.Value + r.Value, Token: operator}, nil
		}
	}

	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Expression, error) {
			return ast.IntExpression{Value: a + b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.FloatExpression{Value: a + b, Token: operator}, nil
		})
}

func evaluateSubtraction(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Expression, error) {
			return ast.IntExpression{Value: a - b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.FloatExpression{Value: a - b, Token: operator}, nil
		})
}

func evaluateMultiplication(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Expression, error) {
			return ast.IntExpression{Value: a * b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			return ast.FloatExpression{Value: a * b, Token: operator}, nil
		})
}

func evaluateDivision(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Expression, error) {
			if b == 0 {
				return nil, ErrDivisionByZero(operator)
			}
			return ast.IntExpression{Value: a / b, Token: operator}, nil
		},
		func(a, b float64) (ast.Expression, error) {
			if b == 0 {
				return nil, ErrDivisionByZero(operator)
			}
			return ast.FloatExpression{Value: a / b, Token: operator}, nil
		})
}
