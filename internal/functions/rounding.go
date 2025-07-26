package functions

import (
	"github.com/pblazh/csvss/internal/ast"
)

func Int(call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard("INT(number)", ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	result, _ := ast.ToInt(&(values[0]))
	return result, nil
}

func Round(up bool, format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsNumeric, ast.IsNumeric)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	first, _ := ast.ToFloat(&(values[0]))
	second, _ := ast.ToFloat(&(values[1]))
	return ast.FloatExpression{Value: roundPrecise(up, first.Value, second.Value)}, nil
}
