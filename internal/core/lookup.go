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
