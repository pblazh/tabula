package functions

import (
	"strings"

	"github.com/pblazh/csvss/internal/ast"
)

func Concat(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeSameTypeGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	var result strings.Builder
	for _, arg := range values {
		a := arg.(ast.StringExpression)
		result.WriteString(a.Value)
	}
	return ast.StringExpression{Value: result.String()}, nil
}

func Len(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.IntExpression{Value: len(a.Value)}, nil
}

func Lower(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.StringExpression{Value: strings.ToLower(a.Value)}, nil
}

func Upper(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.StringExpression{Value: strings.ToUpper(a.Value)}, nil
}
