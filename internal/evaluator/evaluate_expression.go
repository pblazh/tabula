package evaluator

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

// EvaluateExpression evaluates any AST expression and returns the result
func EvaluateExpression(expr ast.Expression) (ast.Expression, error) {
	switch node := expr.(type) {
	case ast.IntExpression, ast.FloatExpression, ast.BooleanExpression, ast.StringExpression, ast.IdentifierExpression:
		return node, nil
	case ast.PrefixExpression:
		return evaluatePrefixExpression(node)
	case ast.InfixExpression:
		return evaluateInfixExpression(node)
	case ast.CallExpression:
		return evaluateCallExpression(node)
	case ast.RangeExpression:
		return node, nil
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}

func evaluatePrefixExpression(expr ast.PrefixExpression) (ast.Expression, error) {
	value, err := EvaluateExpression(expr.Value)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case lexer.MINUS:
		return evaluateNegation(value, expr.Operator)
	case lexer.NOT:
		return evaluateNot(value, expr.Operator)
	default:
		return nil, fmt.Errorf("unsupported prefix operator: %s", expr.Operator.Literal)
	}
}

func evaluateInfixExpression(expr ast.InfixExpression) (ast.Expression, error) {
	left, err := EvaluateExpression(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := EvaluateExpression(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case lexer.PLUS:
		return evaluateAddition(left, right, expr.Operator)
	case lexer.MINUS:
		return evaluateSubtraction(left, right, expr.Operator)
	case lexer.MULT:
		return evaluateMultiplication(left, right, expr.Operator)
	case lexer.DIV:
		return evaluateDivision(left, right, expr.Operator)
	case lexer.EQUAL:
		return evaluateEquality(left, right, expr.Operator)
	case lexer.NOT_EQUAL:
		return evaluateInequality(left, right, expr.Operator)
	case lexer.LESS:
		return evaluateLessThan(left, right, expr.Operator)
	case lexer.GREATER:
		return evaluateGreaterThan(left, right, expr.Operator)
	default:
		return nil, fmt.Errorf("unsupported operator: %s", expr.Operator.Literal)
	}
}

func evaluateCallExpression(expr ast.CallExpression) (ast.Expression, error) {
	args := make([]ast.Expression, len(expr.Arguments))
	for i, arg := range expr.Arguments {
		evaluated, err := EvaluateExpression(arg)
		if err != nil {
			return nil, err
		}
		args[i] = evaluated
	}

	// For now, just return the identifier as we haven't implemented function calls yet
	return expr.Identifier, nil
}
