package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

func evaluateNumericOperation(
	left, right ast.Expression,
	operator lexer.Token,
	intOp func(int, int) (ast.Expression, error),
	floatOp func(float64, float64) (ast.Expression, error),
) (ast.Expression, error) {
	switch l := left.(type) {
	case ast.IntExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return intOp(l.Value, r.Value)
		case ast.FloatExpression:
			return floatOp(float64(l.Value), r.Value)
		}
	case ast.FloatExpression:
		switch r := right.(type) {
		case ast.IntExpression:
			return floatOp(l.Value, float64(r.Value))
		case ast.FloatExpression:
			return floatOp(l.Value, r.Value)
		}
	}
	return nil, ErrUnsupportedBinaryOperation(operator, left, right)
}
