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
	return ast.StringExpression{Value: result.String(), Token: call.Token}, nil
}

func Len(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.IntExpression{Value: len(a.Value), Token: call.Token}, nil
}

func Lower(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.StringExpression{Value: strings.ToLower(a.Value), Token: call.Token}, nil
}

func Upper(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.StringExpression{Value: strings.ToUpper(a.Value), Token: call.Token}, nil
}

func Trim(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ast.StringExpression{Value: strings.TrimSpace(a.Value), Token: call.Token}, nil
}

func Exact(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	b := values[1].(ast.StringExpression)
	return ast.BooleanExpression{Value: strings.Compare(a.Value, b.Value) == 0, Token: call.Token}, nil
}
