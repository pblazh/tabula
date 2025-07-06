// Package evaluator provides evaluation functionality for the CSV spreadsheet language,
// converting an abstract syntax tree into computed values.
package evaluator

import (
	"github.com/pblazh/csvss/internal/ast"
)

// Evaluate evaluates a program (list of statements) with the given context
// context is a two-dimensional array of strings representing the spreadsheet data
func Evaluate(program ast.Program, context [][]string) ([][]string, error) {
	// TODO: Implement evaluation logic
	return nil, nil
}
