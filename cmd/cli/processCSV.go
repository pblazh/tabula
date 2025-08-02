package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/pblazh/csvss/internal/ast"
	"github.com/pblazh/csvss/internal/evaluator"
)

func processCSV(config *Config, scriptReader io.Reader, csvReader io.Reader, csvWriter io.Writer, comments map[int]string) error {
	// Read and parse CSV
	reader := csv.NewReader(csvReader)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.Comment = '#'

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %v", err)
	}

	for i, row := range records {
		for j, cel := range row {
			records[i][j] = strings.TrimSpace(cel)
		}
	}

	// Parse script
	program, err := evaluator.ParseProgram(scriptReader, config.Script)
	if err != nil {
		return fmt.Errorf("error parsing script: %v", err)
	}

	// Sort program topologically if Sort flag is set
	if config.Sort {
		program, err = ast.SortProgram(program)
		if err != nil {
			return fmt.Errorf("error sorting script statements: %v", err)
		}
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return fmt.Errorf("error evaluating script %s: %v", config.Script, err)
	}

	if config.Align {
		return writeAligned(csvWriter, result, comments)
	}

	// Output result in the expected format
	return writeCompact(csvWriter, result, comments)
}
