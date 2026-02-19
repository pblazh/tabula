package evaluator

import (
	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/lexer"
)

func evaluateNegation(expr ast.Node, token lexer.Token) (ast.Node, error) {
	switch r := expr.(type) {
	case ast.IntExpression:
		return ast.IntExpression{Value: -r.Value, Token: token}, nil
	case ast.FloatExpression:
		return ast.FloatExpression{Value: -r.Value, Token: token}, nil
	default:
		return nil, ErrUnsupportedOperation(token, expr)
	}
}

func evaluateNot(expr ast.Node, token lexer.Token) (ast.Node, error) {
	switch r := expr.(type) {
	case ast.BooleanExpression:
		return ast.BooleanExpression{Value: !r.Value, Token: token}, nil
	default:
		return nil, ErrUnsupportedOperation(token, expr)
	}
}

func evaluateAddition(left, right ast.Node, operator lexer.Token) (ast.Node, error) {
	if l, ok := left.(ast.StringExpression); ok {
		if r, ok := right.(ast.StringExpression); ok {
			return ast.StringExpression{Value: l.Value + r.Value, Token: operator}, nil
		}
	}

	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Node, error) {
			return ast.IntExpression{Value: a + b, Token: operator}, nil
		},
		func(a, b float64) (ast.Node, error) {
			return ast.FloatExpression{Value: a + b, Token: operator}, nil
		})
}

func evaluateSubtraction(left, right ast.Node, operator lexer.Token) (ast.Node, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Node, error) {
			return ast.IntExpression{Value: a - b, Token: operator}, nil
		},
		func(a, b float64) (ast.Node, error) {
			return ast.FloatExpression{Value: a - b, Token: operator}, nil
		})
}

func evaluateMultiplication(
	left, right ast.Node,
	operator lexer.Token,
) (ast.Node, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Node, error) {
			return ast.IntExpression{Value: a * b, Token: operator}, nil
		},
		func(a, b float64) (ast.Node, error) {
			return ast.FloatExpression{Value: a * b, Token: operator}, nil
		})
}

func evaluateDivision(left, right ast.Node, operator lexer.Token) (ast.Node, error) {
	return evaluateNumericOperation(left, right, operator,
		func(a, b int) (ast.Node, error) {
			if b == 0 {
				return nil, ErrDivisionByZero(operator)
			}
			return ast.IntExpression{Value: a / b, Token: operator}, nil
		},
		func(a, b float64) (ast.Node, error) {
			if b == 0 {
				return nil, ErrDivisionByZero(operator)
			}
			return ast.FloatExpression{Value: a / b, Token: operator}, nil
		})
}
