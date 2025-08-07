package functions

import (
	"time"

	"github.com/pblazh/csvss/internal/ast"
)

func ToDate(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsString, ast.IsString)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	layout := values[0].(ast.StringExpression)
	value := values[1].(ast.StringExpression)

	parsed, err := time.Parse(layout.Value, value.Value)
	if err != nil {
		return nil, ErrExecuting(format, call, err)
	}

	return ast.DateExpression{Value: parsed, Token: call.Token}, nil
}

func FromDate(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsString, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	layout := values[0].(ast.StringExpression)
	value := values[1].(ast.DateExpression)

	formated := value.Value.Format(layout.Value)

	return ast.StringExpression{Value: formated, Token: call.Token}, nil
}

func Day(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: value.Value.Day(), Token: call.Token}, nil
}

func Hour(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: value.Value.Hour(), Token: call.Token}, nil
}

func Minute(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: value.Value.Minute(), Token: call.Token}, nil
}

func Month(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: int(value.Value.Month()), Token: call.Token}, nil
}
