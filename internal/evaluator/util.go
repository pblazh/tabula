package evaluator

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func evaluateArithmetic(left, right ast.Expression, operator lexer.Token, intOp func(int, int) int, floatOp func(float64, float64) float64) (ast.Expression, error) {
	switch l := left.(type) {
	case ast.IntExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return ast.IntExpression{Value: intOp(l.Value, r.Value), Token: operator}, nil
		case ast.FloatExpression:
			return ast.FloatExpression{Value: floatOp(float64(l.Value), r.Value), Token: operator}, nil
		}
	case ast.FloatExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return ast.FloatExpression{Value: floatOp(l.Value, float64(r.Value)), Token: operator}, nil
		case ast.FloatExpression:
			return ast.FloatExpression{Value: floatOp(l.Value, r.Value), Token: operator}, nil
		}
	}
	return nil, fmt.Errorf("operator %s is not supported for type: %s and %s", operator, ast.TypeName(left), ast.TypeName(right))
}

func evaluateNumericComparison(left, right ast.Expression, operator lexer.Token, intOp func(int, int) bool, floatOp func(float64, float64) bool) (ast.Expression, error) {
	switch l := left.(type) {
	case ast.IntExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return ast.BooleanExpression{Value: intOp(l.Value, r.Value), Token: operator}, nil
		case ast.FloatExpression:
			return ast.BooleanExpression{Value: floatOp(float64(l.Value), r.Value), Token: operator}, nil
		}
	case ast.FloatExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return ast.BooleanExpression{Value: floatOp(l.Value, float64(r.Value)), Token: operator}, nil
		case ast.FloatExpression:
			return ast.BooleanExpression{Value: floatOp(l.Value, r.Value), Token: operator}, nil
		}
	}
	return nil, fmt.Errorf("numeric comparison %s is not supported for type: %s and %s", operator, ast.TypeName(left), ast.TypeName(right))
}
