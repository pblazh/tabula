// Package evaluator provides evaluation functionality for the CSV spreadsheet language,
// converting an abstract syntax tree into computed values.
package evaluator

import (
	"fmt"

	"github.com/pblazh/csvss/internal/ast"
)

// Evaluate evaluates a program (list of statements) with the given context
// input is a two-dimensional array of strings representing the spreadsheet data
func Evaluate(program ast.Program, input [][]string) ([][]string, error) {
	context := make(map[string]string)
	format := make(map[string]string)
	for _, statement := range program {
		error := EvaluateStatement(statement, context, format)
		if error != nil {
			return nil, fmt.Errorf("%s caused %s", statement, error)
		}
	}
	fmt.Println(context)
	return nil, nil
}
