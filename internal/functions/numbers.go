package functions

import (
	"math"

	"github.com/pblazh/csvss/internal/ast"
)

func sum[T Number](values ...T) T {
	var result T
	for _, n := range values {
		result += n
	}
	return result
}

func product[T Number](values ...T) T {
	result := T(1)
	for _, n := range values {
		result *= n
	}
	return result
}

func average[T Number](values ...T) T {
	if len(values) == 0 {
		return T(0)
	}

	total := T(0)
	for _, n := range values {
		total += n
	}
	return T(total / T(len(values)))
}

func max[T Number](values ...T) T {
	if len(values) == 0 {
		return T(math.Inf(-1))
	}

	total := values[0]
	for _, n := range values {
		if total < n {
			total = n
		}
	}
	return T(total)
}

func min[T Number](values ...T) T {
	if len(values) == 0 {
		return T(math.Inf(1))
	}

	total := values[0]
	for _, n := range values {
		if total > n {
			total = n
		}
	}
	return T(total)
}

func abs[T Number](values ...T) T {
	return T(math.Abs(float64(values[0])))
}

func mod[T Number](values ...T) T {
	a := values[0]
	b := values[1]
	return T(math.Mod(float64(a), float64(b)))
}

func roundPrecise(value, precision float64) float64 {
	return math.Round(value/precision) * precision
}

func roundUpPrecise(value, precision float64) float64 {
	return math.Ceil(value/precision) * precision
}

func roundDownPrecise(value, precision float64) float64 {
	return math.Floor(value/precision) * precision
}

func parseStringExpressions(call ast.CallExpression, values []ast.Expression) ([]ast.Expression, ast.Expression) {
	var converted []ast.Expression
	for _, v := range values {
		switch expr := v.(type) {
		case ast.IntExpression:
			converted = append(converted, expr)
		case ast.FloatExpression:
			converted = append(converted, expr)
		case ast.StringExpression:
			res := parseNumberWithoutFormat(expr.Value)
			if res == nil {
				return nil, expr
			}
			converted = append(converted, res)
		default:
			return nil, expr
		}
	}
	return converted, nil
}

func count(call ast.CallExpression, values ...ast.Expression) ast.IntExpression {
	var result int
	for _, value := range values {
		if ast.IsNumeric(value) || ast.IsDate(value) {
			result++
		}
	}
	return ast.IntExpression{Value: result, Token: call.Token}
}

func counta(call ast.CallExpression, values ...ast.Expression) ast.IntExpression {
	var result int
	for _, value := range values {
		switch v := value.(type) {
		case ast.StringExpression:
			if v.Value != "" {
				result++
			}
		default:
			result++
		}
	}
	return ast.IntExpression{Value: result, Token: call.Token}
}
