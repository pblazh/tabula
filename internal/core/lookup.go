package core

import (
	"github.com/pblazh/tabula/internal/ast"
)

func Address(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsInt, ast.IsInt)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	row := values[0].(ast.IntExpression).Value
	column := values[1].(ast.IntExpression).Value

	address := ast.ToCell(column-1, row-1)

	return ast.IdentifierExpression{Value: address, Token: call.Token}, nil
}

func Row(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	identifierGuard := MakeExactTypesGuard(format, ast.IsIdentifier)
	identifierErr := identifierGuard(call, values...)

	rangeGuard := MakeExactTypesGuard(format, ast.IsRange)
	rangeErr := rangeGuard(call, values...)

	if identifierErr != nil && rangeErr != nil {
		return nil, identifierErr
	}

	var cell string

	if identifierErr == nil {
		cell = values[0].(ast.IdentifierExpression).Value
	} else {
		cell = values[0].(ast.RangeExpression).Value[0]
	}

	_, row := ast.ParseCell(cell)

	return ast.IntExpression{Value: row + 1, Token: call.Token}, nil
}

func Column(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	identifierGuard := MakeExactTypesGuard(format, ast.IsIdentifier)
	identifierErr := identifierGuard(call, values...)

	rangeGuard := MakeExactTypesGuard(format, ast.IsRange)
	rangeErr := rangeGuard(call, values...)

	if identifierErr != nil && rangeErr != nil {
		return nil, identifierErr
	}

	var cell string

	if identifierErr == nil {
		cell = values[0].(ast.IdentifierExpression).Value
	} else {
		cell = values[0].(ast.RangeExpression).Value[0]
	}

	column, _ := ast.ParseCell(cell)

	return ast.IntExpression{Value: column + 1, Token: call.Token}, nil
}

func Ref(
	context map[string]string, input [][]string, formats map[string]string,
	format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	argsGuard := MakeExactTypesGuard(format, ast.IsString)
	identifierErr := argsGuard(call, values...)
	if identifierErr != nil {
		return nil, identifierErr
	}

	identifier := values[0].(ast.StringExpression)
	valueFormat := formats[identifier.Value]

	if ast.IsCellIdentifier(identifier.Value) {
		col, row := ast.ParseCell(identifier.Value)
		if col >= len(input) || col >= len(input[row]) {
			return ast.StringExpression{Value: "", Token: call.Token}, nil
		}

		return ReadValue(input[row][col], valueFormat)
	}

	value, ok := context[identifier.Value]
	if !ok {
		return nil, ErrUnsupportedArgument(format, call, identifier)
	}

	return ReadValue(value, valueFormat)
}
