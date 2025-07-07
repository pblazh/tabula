package evaluator

import (
	"fmt"

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
		return nil, fmt.Errorf("%s is not supported for type: %s", token, ast.TypeName(expr))
	}
}

func evaluateNot(expr ast.Expression, token lexer.Token) (ast.Expression, error) {
	switch r := expr.(type) {
	case ast.BooleanExpression:
		return ast.BooleanExpression{Value: !r.Value, Token: token}, nil
	default:
		return nil, fmt.Errorf("%s is not supported for type: %s", token, ast.TypeName(expr))
	}
}

func evaluateAddition(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	if l, ok := left.(ast.StringExpression); ok {
		if r, ok := right.(ast.StringExpression); ok {
			return ast.StringExpression{Token: lexer.Token{Type: lexer.STRING, Literal: l.Token.Literal + r.Token.Literal}}, nil
		}
	}

	return evaluateArithmetic(left, right, operator,
		func(a, b int) int { return a + b },
		func(a, b float64) float64 { return a + b })
}

func evaluateSubtraction(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateArithmetic(left, right, operator,
		func(a, b int) int { return a - b },
		func(a, b float64) float64 { return a - b })
}

func evaluateMultiplication(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateArithmetic(left, right, operator,
		func(a, b int) int { return a * b },
		func(a, b float64) float64 { return a * b })
}

func evaluateDivision(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	switch r := right.(type) {
	case ast.IntExpression:
		if r.Value == 0 {
			return nil, fmt.Errorf("division by zero at %s", operator)
		}
	case ast.FloatExpression:
		if r.Value == 0 {
			return nil, fmt.Errorf("division by zero at %s", operator)
		}
	}

	switch l := left.(type) {
	case ast.IntExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			// int / int returns int
			return ast.IntExpression{Value: l.Value / r.Value, Token: operator}, nil
		case ast.FloatExpression:
			// int / float returns float
			return ast.FloatExpression{Value: float64(l.Value) / r.Value, Token: operator}, nil
		}
	case ast.FloatExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			// float / int returns float
			return ast.FloatExpression{Value: l.Value / float64(r.Value), Token: operator}, nil
		case ast.FloatExpression:
			// float / float returns float
			return ast.FloatExpression{Value: l.Value / r.Value, Token: operator}, nil
		}
	}
	return nil, fmt.Errorf("cannot divide %s by %s", ast.TypeName(left), ast.TypeName(right))
}

func evaluateEquality(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	// Handle numeric comparisons
	if _, ok := left.(ast.IntExpression); ok {
		return evaluateNumericComparison(left, right, operator,
			func(a, b int) bool { return a == b },
			func(a, b float64) bool { return a == b })
	}
	if _, ok := left.(ast.FloatExpression); ok {
		return evaluateNumericComparison(left, right, operator,
			func(a, b int) bool { return a == b },
			func(a, b float64) bool { return a == b })
	}

	// Handle string comparisons
	if l, ok := left.(ast.StringExpression); ok {
		if r, ok := right.(ast.StringExpression); ok {
			return ast.BooleanExpression{Value: l.Token.Literal == r.Token.Literal, Token: operator}, nil
		}
	}

	// Handle boolean comparisons
	if l, ok := left.(ast.BooleanExpression); ok {
		if r, ok := right.(ast.BooleanExpression); ok {
			return ast.BooleanExpression{Value: l.Value == r.Value, Token: operator}, nil
		}
	}

	return ast.BooleanExpression{Value: false, Token: operator}, nil
}

func evaluateInequality(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	result, err := evaluateEquality(left, right, operator)
	if err != nil {
		return nil, err
	}
	if b, ok := result.(ast.BooleanExpression); ok {
		return ast.BooleanExpression{Value: !b.Value, Token: operator}, nil
	}
	return nil, fmt.Errorf("inequality evaluation failed at %s", operator)
}

func evaluateLessThan(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	// Handle numeric comparisons
	if _, ok := left.(ast.IntExpression); ok {
		return evaluateNumericComparison(left, right, operator,
			func(a, b int) bool { return a < b },
			func(a, b float64) bool { return a < b })
	}
	if _, ok := left.(ast.FloatExpression); ok {
		return evaluateNumericComparison(left, right, operator,
			func(a, b int) bool { return a < b },
			func(a, b float64) bool { return a < b })
	}

	// Handle string comparisons
	if l, ok := left.(ast.StringExpression); ok {
		if r, ok := right.(ast.StringExpression); ok {
			return ast.BooleanExpression{Value: l.Token.Literal < r.Token.Literal, Token: operator}, nil
		}
	}

	return nil, fmt.Errorf("cannot compare %s and %s at %s", ast.TypeName(left), ast.TypeName(right), operator)
}

func evaluateGreaterThan(left, right ast.Expression, operator lexer.Token) (ast.Expression, error) {
	return evaluateLessThan(right, left, operator)
}
