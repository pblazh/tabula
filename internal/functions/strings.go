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
	n = max(0, n)
	n = min(n, len(a.Value))

	return ast.StringExpression{Value: a.Value[0:n], Token: call.Token}, nil
}

func Right(format string,
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
	n = max(0, n)
	n = min(n, len(a.Value))

	// Calculate start position for right-side extraction
	start := len(a.Value) - n
	start = max(0, start)

	return ast.StringExpression{Value: a.Value[start:], Token: call.Token}, nil
}

func Mid(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString, ast.IsInt, ast.IsInt)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	str := values[0].(ast.StringExpression).Value
	start := values[1].(ast.IntExpression).Value - 1
	start = max(0, start)

	end := start + values[2].(ast.IntExpression).Value
	end = min(len(str), end)

	return ast.StringExpression{Value: str[start:end], Token: call.Token}, nil
}

func Substitute(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	var guard CallGuard
	if len(values) == 3 {
		guard = MakeExactTypesGuard(format, ast.IsString, ast.IsString, ast.IsString)
	} else {
		guard = MakeExactTypesGuard(format, ast.IsString, ast.IsString, ast.IsString, ast.IsInt)
	}

	if err := guard(call, values...); err != nil {
		return nil, err
	}

	source := values[0].(ast.StringExpression).Value
	target := values[1].(ast.StringExpression).Value
	replacement := values[2].(ast.StringExpression).Value

	if target == "" {
		return ast.StringExpression{Value: source, Token: call.Token}, nil
	}

	n := 0
	if len(values) == 4 {
		n = values[3].(ast.IntExpression).Value
		if n < 0 {
			return nil, ErrUnsupportedArgument(format, call, values[3])
		}
	}

	if n == 0 {
		return ast.StringExpression{Value: strings.ReplaceAll(source, target, replacement), Token: call.Token}, nil
	}

	indices := []int{}
	searchStr := source
	offset := 0

	// Find all occurrences
	for {
		idx := strings.Index(searchStr, target)
		if idx == -1 {
			break
		}
		indices = append(indices, offset+idx)
		offset += idx + len(target)
		searchStr = source[offset:]
	}

	// Replace required one
	if n > 0 && n <= len(indices) {
		replaceIdx := indices[n-1]
		source = source[:replaceIdx] + replacement + source[replaceIdx+len(target):]
	}

	return ast.StringExpression{Value: source, Token: call.Token}, nil
}

func Value(format string,
	call ast.CallExpression, values ...ast.Expression,
) (ast.Expression, error) {
	guard := MakeExactTypesGuard(format, ast.IsString)
	if err := guard(call, values...); err != nil {
		return nil, err
	}

	a := values[0].(ast.StringExpression)
	return ParseWithoutFormat(a.Value)
}
