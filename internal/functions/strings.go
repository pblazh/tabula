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

func Find(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	var guard CallGuard
	var start int
	if len(values) == 2 {
		guard = MakeExactTypesGuard(format, ast.IsString, ast.IsString)
		start = 0
	} else {
		guard = MakeExactTypesGuard(format, ast.IsString, ast.IsString, ast.IsInt)
	}

	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	b := values[1].(ast.StringExpression)

	if len(values) == 3 {
		start = values[2].(ast.IntExpression).Value
	}

	// Handle edge cases for start position
	if start < 0 || start > len(a.Value) {
		return ast.IntExpression{Value: -1, Token: call.Token}, nil
	}

	result := strings.Index(a.Value[start:], b.Value)
	if result == -1 {
		return ast.IntExpression{Value: -1, Token: call.Token}, nil
	}
	return ast.IntExpression{Value: result + start, Token: call.Token}, nil
}

func Left(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	var guard CallGuard
	var n int
	if len(values) == 1 {
		guard = MakeExactTypesGuard(format, ast.IsString)
	} else {
		guard = MakeExactTypesGuard(format, ast.IsString, ast.IsInt)
	}

	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)

	if len(values) == 2 {
		n = values[1].(ast.IntExpression).Value
	} else {
		n = 1
	}

	// Handle edge cases for count
	if n < 0 {
		n = 0
	}
	if n > len(a.Value) {
		n = len(a.Value)
	}

	return ast.StringExpression{Value: a.Value[0:n], Token: call.Token}, nil
}
