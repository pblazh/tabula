package main

import (
	"encoding/csv"
	"io"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/evaluator"
)

func processCSV(config *Config, scriptReader io.Reader, csvReader io.Reader, csvWriter io.Writer, comments map[int]string) error {
	// Read and parse CSV
	reader := csv.NewReader(csvReader)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return ErrReadCSV(err)
	}

	program, identifiers, err := evaluator.ParseProgram(scriptReader, config.Name)
	records = ensureProgramDimmensions(identifiers, records)

	if err != nil {
		return ErrParseScript(err)
	}

	// Sort program topologically if Sort flag is set
	if config.Sort {
		program, err = ast.SortProgram(program)
		if err != nil {
			return ErrSortScriptStatements(err)
		}
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return ErrEvaluateScript(config.Name, err)
	}

	if config.Align {
		return writeAligned(csvWriter, result, comments)
	}

	// Output result in the expected format
	return writeCompact(csvWriter, result, comments)
}
