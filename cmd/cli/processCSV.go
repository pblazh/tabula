package main

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/pblazh/csvss/internal/evaluator"
)

func processCSV(scriptPath string, scriptReader io.Reader, csvReader io.Reader, csvWriter io.Writer) error {
	// Read and parse CSV
	reader := csv.NewReader(csvReader)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV: %v", err)
	}

	// Parse script
	program, err := evaluator.ParseProgram(scriptReader, scriptPath)
	if err != nil {
		return fmt.Errorf("error parsing script: %v", err)
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return fmt.Errorf("error evaluating script %s: %v", scriptPath, err)
	}

	// Output result in the expected format
	writer := csv.NewWriter(csvWriter)
	defer writer.Flush()

	for _, row := range result {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing CSV output: %v", err)
		}
	}

	return nil
}
