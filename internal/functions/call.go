package functions

import (
	"os/exec"
	"strings"

	"github.com/pblazh/csvss/internal/ast"
)

func Call(format string, call ast.CallExpression, values ...ast.Expression) (ast.Expression, error) {
	callGuard := MakeSameTypeGuard(format, ast.IsString)
	if err := callGuard(call, values...); err != nil {
		return nil, err
	}

	name, _ := ast.ToString(&(values[0]))

	args := make([]string, 0, len(values)-1)
	for _, arg := range values[1:] {
		str, _ := ast.ToString(&arg)
		args = append(args, str.Value)
	}

	// Execute external program
	cmd := exec.Command(name.Value, args...)
	output, err := cmd.Output()
	if err != nil {
		return ast.StringExpression{Value: err.Error(), Token: call.Token}, err
	}

	// Return stdout output as one string (trimmed of trailing whitespace)
	result := strings.TrimSpace(strings.ReplaceAll(string(output), "\n", " "))

	return ast.StringExpression{Value: result, Token: call.Token}, nil
}
