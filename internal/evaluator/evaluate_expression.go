package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
)

// EvaluateExpression evaluates any AST expression and returns the result
func EvaluateExpression(expr ast.Expression, context map[string]string, input [][]string, formats map[string]string) (ast.Expression, error) {
	switch node := expr.(type) {
	case ast.IntExpression, ast.FloatExpression, ast.BooleanExpression, ast.StringExpression:
		return node, nil
	case ast.IdentifierExpression:
		if ast.IsCellIdentifier(node.Token.Literal) {
			return evaluateCellExpression(node, input, formats)
		}
		return evaluateVariableExpression(node, context, formats)
	case ast.PrefixExpression:
		return evaluatePrefixExpression(node, context, input, formats)
	case ast.InfixExpression:
		return evaluateInfixExpression(node, context, input, formats)
	case ast.CallExpression:
		return evaluateCallExpression(node, context, input, formats)
	case ast.RangeExpression:
		return node, nil
	default:
		return nil, ErrUnknownExpressionType(expr)
	}
}

func evaluatePrefixExpression(expr ast.PrefixExpression, context map[string]string, input [][]string, formats map[string]string) (ast.Expression, error) {
	value, err := EvaluateExpression(expr.Value, context, input, formats)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case lexer.MINUS:
		return evaluateNegation(value, expr.Operator)
	case lexer.NOT:
		return evaluateNot(value, expr.Operator)
	default:
		return nil, ErrUnsupportedPrefixOperator(expr.Operator.Literal)
	}
}

func evaluateInfixExpression(expr ast.InfixExpression, context map[string]string, input [][]string, formats map[string]string) (ast.Expression, error) {
	left, err := EvaluateExpression(expr.Left, context, input, formats)
	if err != nil {
		return nil, err
	}
	right, err := EvaluateExpression(expr.Right, context, input, formats)
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
		return evaluateEquality(expr.Operator, left, right)
	case lexer.NOT_EQUAL:
		return evaluateInequality(expr.Operator, left, right)
	case lexer.LESS:
		return evaluateLessThan(expr.Operator, left, right)
	case lexer.GREATER:
		return evaluateGreaterThan(expr.Operator, left, right)
	default:
		return nil, ErrUnsupportedOperator(expr.Operator.Literal)
	}
}

func evaluateCallExpression(expr ast.CallExpression, context map[string]string, input [][]string, formats map[string]string) (ast.Expression, error) {
	args := make([]ast.Expression, 0, len(expr.Arguments))
	for _, arg := range expr.Arguments {
		switch a := arg.(type) {
		case ast.RangeExpression:
			evaluated, err := EvaluateRangeExpression(a, input, formats)
			if err != nil {
				return nil, err
			}
			args = append(args, evaluated...)
		default:
			evaluated, err := EvaluateExpression(arg, context, input, formats)
			if err != nil {
				return nil, err
			}
			args = append(args, evaluated)
		}
	}

	identifier := expr.Identifier.String()
	internalFunction, ok := dispatchMap[identifier]
	if !ok {
		return nil, ErrUnsupportedFunctions(identifier)
	}
	return internalFunction(expr, args...)
}

func evaluateVariableExpression(expr ast.IdentifierExpression, context map[string]string, formats map[string]string) (ast.Expression, error) {
	name := expr.Token.Literal
	value, exists := context[name]
	if !exists {
		return nil, ErrVariableNotFound(expr)
	}

	format := formats[name]
	return ReadValue(value, format)
}

// evaluateCellExpression evaluates a cell reference (like A1, B2) and returns the value from the CSV input
func evaluateCellExpression(expr ast.IdentifierExpression, input [][]string, formats map[string]string) (ast.Expression, error) {
	cellRef := expr.Token.Literal
	col, row := ast.ParseCell(cellRef)

	// Check bounds
	if row < 0 || row >= len(input) {
		return nil, ErrCellOutOfBounds(cellRef, "row", row)
	}
	if col < 0 || col >= len(input[row]) {
		return nil, ErrCellOutOfBounds(cellRef, "column", col)
	}

	// Get the value from the CSV input
	value := input[row][col]
	return ReadValue(value, formats[cellRef])
}

// EvaluateRangeExpression evaluates a range cell reference (like A1:A2, A1:B2) and returns the value from the CSV input
func EvaluateRangeExpression(expr ast.RangeExpression, input [][]string, formats map[string]string) ([]ast.Expression, error) {
	cells := make([]ast.IdentifierExpression, len(expr.Value))
	for i, cell := range expr.Value {
		cells[i] = ast.IdentifierExpression{Token: lexer.Token{Literal: cell}}
	}

	result := make([]ast.Expression, len(expr.Value))

	for i, cell := range cells {
		res, err := evaluateCellExpression(cell, input, formats)
		if err != nil {
			return nil, err
		}
		result[i] = res
	}
	return result, nil
}
