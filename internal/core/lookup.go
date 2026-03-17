package core

import (
	"strings"

	"github.com/pblazh/tabula/internal/ast"
)

func Address(format string,
	call ast.CallExpression, values ...ast.Node,
) (ast.Node, error) {
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
	call ast.CallExpression, values ...ast.Node,
) (ast.Node, error) {
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
	call ast.CallExpression, values ...ast.Node,
) (ast.Node, error) {
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
	call ast.CallExpression, values ...ast.Node,
) (ast.Node, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	addressExpression := values[0].(ast.StringExpression)
	address := strings.TrimSpace(addressExpression.Value)
	valueFormat := formats[address]

	value, ok := context[address]
	if ok {
		return ReadValue(value, valueFormat)
	}

	if ast.IsCellIdentifier(address) {
		col, row := ast.ParseCell(address)
		return ReadValue(input[row][col], valueFormat)
	}

	var cells []string
	for segment := range strings.SplitSeq(addressExpression.Value, ",") {
		segment = strings.TrimSpace(segment)

		if ast.IsCellIdentifier(segment) {
			cells = append(cells, segment)
			continue
		}

		if strings.Contains(segment, ":") {
			parts := strings.SplitN(segment, ":", 2)
			start := strings.TrimSpace(parts[0])
			end := strings.TrimSpace(parts[1])
			expanded, err := ast.ExpandRange(start, end)
			if err != nil {
				return nil, ErrExpand(err)
			}
			cells = append(cells, expanded...)
			continue
		}

		return nil, ErrUnsupportedArgument(format, call, addressExpression)
	}

	return ast.RangeExpression{Value: cells, Token: call.Token}, nil
}
