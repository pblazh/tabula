package functions

import (
	"errors"
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

func Second(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: int(value.Value.Second()), Token: call.Token}, nil
}

func Year(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: int(value.Value.Year()), Token: call.Token}, nil
}

func Weekday(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	value := values[0].(ast.DateExpression)

	return ast.IntExpression{Value: int(value.Value.Weekday()), Token: call.Token}, nil
}

func Now(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeArityGuard(format, 0)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	return ast.DateExpression{Value: time.Now(), Token: call.Token}, nil
}

func Date(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsInt, ast.IsInt, ast.IsInt)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	year := values[0].(ast.IntExpression).Value
	month := values[1].(ast.IntExpression).Value
	day := values[2].(ast.IntExpression).Value

	return ast.DateExpression{Value: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), Token: call.Token}, nil
}

func Days(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate, ast.IsDate)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	start := values[0].(ast.DateExpression).Value
	end := values[1].(ast.DateExpression).Value

	result, err := calculatesDatesDifference("D", start, end)
	if err != nil {
		return nil, ErrUnsupportedArgument(format, call, values[2])
	}

	return ast.IntExpression{Value: result, Token: call.Token}, nil
}

func DateDiff(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeExactTypesGuard(format, ast.IsDate, ast.IsDate, ast.IsString)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	start := values[0].(ast.DateExpression).Value
	end := values[1].(ast.DateExpression).Value
	unit := values[2].(ast.StringExpression).Value

	result, err := calculatesDatesDifference(unit, start, end)
	if err != nil {
		return nil, ErrUnsupportedArgument(format, call, values[2])
	}

	return ast.IntExpression{Value: result, Token: call.Token}, nil
}

func calculatesDatesDifference(unit string, start, end time.Time) (int, error) {
	if unit != "Y" && unit != "M" && unit != "D" && unit != "MD" && unit != "YM" && unit != "YD" {
		return 0, errors.New(unit)
	}

	var result int

	switch unit {
	case "Y":
		// Years between dates
		result = end.Year() - start.Year()
		if end.Month() < start.Month() || (end.Month() == start.Month() && end.Day() < start.Day()) {
			result--
		}
	case "M":
		// Months between dates
		result = int(end.Month()) - int(start.Month()) + (end.Year()-start.Year())*12
		if end.Day() < start.Day() {
			result--
		}
	case "D":
		// Days between dates
		duration := end.Sub(start)
		result = int(duration.Hours() / 24)
	case "MD":
		// Days between dates, ignoring months and years
		result = end.Day() - start.Day()
	case "YM":
		// Months between dates, ignoring years
		result = int(end.Month()) - int(start.Month())
	case "YD":
		// Days between dates, ignoring years
		result = end.Day() - start.Day()
	}
	return result, nil
}
