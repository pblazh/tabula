package evaluator

import (
	"text/scanner"

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
	args := make([]ast.Expression, len(expr.Arguments))
	for i, arg := range expr.Arguments {
		evaluated, err := EvaluateExpression(arg, context, input, formats)
		if err != nil {
			return nil, err
		}
		args[i] = evaluated
	}

	// For now, just return the identifier as we haven't implemented function calls yet
	return expr.Identifier, nil // TODO: Remove this placeholder
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

	// Convert column letters to column index (A=0, B=1, etc.)
	colIndex := ast.ColumnToIndex(col)
	rowIndex := row - 1 // Convert to 0-based index

	// Check bounds
	if rowIndex < 0 || rowIndex >= len(input) {
		return nil, ErrCellOutOfBounds(cellRef, "row", row)
	}
	if colIndex < 0 || colIndex >= len(input[rowIndex]) {
		return nil, ErrCellOutOfBounds(cellRef, "column", colIndex+1)
	}

	// Get the value from the CSV input
	value := input[rowIndex][colIndex]

	// Return as StringExpression for now
	return ast.StringExpression{
		Value: value,
		Token: lexer.Token{
			Literal: value,
			Position: scanner.Position{
				Filename: "input",
				Line:     rowIndex,
				Column:   colIndex,
			},
		},
	}, nil
}
