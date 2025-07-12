// Package evaluator provides evaluation functionality for the CSV spreadsheet language,
// converting an abstract syntax tree into computed values.
package evaluator

import (
	"io"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/lexer"
	"github.com/pblazh/csvss/internal/parser"
)

// Evaluate evaluates a program (list of statements) with the given context
// input is a two-dimensional array of strings representing the spreadsheet data
func Evaluate(program ast.Program, input [][]string) ([][]string, error) {
	context := make(map[string]string)
	formats := make(map[string]string)
	for _, statement := range program {
		error := EvaluateStatement(statement, context, input, formats)
		if error != nil {
			return nil, ErrStatementExecution(statement.String(), error)
		}
	}
	return input, nil
}

func ParseProgram(r io.Reader, name string) (ast.Program, error) {
	lex := lexer.New(r, name)
	p := parser.New(lex)
	program, _, err := p.Parse()
	return program, err
}
