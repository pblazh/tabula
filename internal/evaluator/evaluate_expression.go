package evaluator

import (
	"fmt"
	"strconv"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

// EvaluateExpression evaluates any AST expression and returns the result
func EvaluateExpression(expr ast.Expression, context map[string]string, format map[string]string) (ast.Expression, error) {
	switch node := expr.(type) {
	case ast.IntExpression, ast.FloatExpression, ast.BooleanExpression, ast.StringExpression:
		return node, nil
	case ast.IdentifierExpression:
		value, ok := context[node.Token.Literal]
		if !ok {
			return nil, fmt.Errorf("%s not found in context", node)
		}
		// TODO: Replace this hardcoded integer parsing with proper type inference
		val, error := strconv.Atoi(value)
		if error != nil {
			return nil, error
		}
		return ast.IntExpression{Value: val}, nil
	case ast.PrefixExpression:
		return evaluatePrefixExpression(node, context, format)
	case ast.InfixExpression:
		return evaluateInfixExpression(node, context, format)
	case ast.CallExpression:
		return evaluateCallExpression(node, context, format)
	case ast.RangeExpression:
		return node, nil
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}

func evaluatePrefixExpression(expr ast.PrefixExpression, context map[string]string, format map[string]string) (ast.Expression, error) {
	value, err := EvaluateExpression(expr.Value, context, format)
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

func evaluateInfixExpression(expr ast.InfixExpression, context map[string]string, format map[string]string) (ast.Expression, error) {
	left, err := EvaluateExpression(expr.Left, context, format)
	if err != nil {
		return nil, err
	}
	right, err := EvaluateExpression(expr.Right, context, format)
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

func evaluateCallExpression(expr ast.CallExpression, context map[string]string, format map[string]string) (ast.Expression, error) {
	args := make([]ast.Expression, len(expr.Arguments))
	for i, arg := range expr.Arguments {
		evaluated, err := EvaluateExpression(arg, context, format)
		if err != nil {
			return nil, err
		}
		args[i] = evaluated
	}

	// For now, just return the identifier as we haven't implemented function calls yet
	return expr.Identifier, nil // TODO: Remove this placeholder
}
