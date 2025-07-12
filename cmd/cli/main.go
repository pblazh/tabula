package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pblazh/csvss/internal/evaluator"
)

var (
	usageMessage      = "Usage: 'csvss -s ./script.file ./table.csv'"
	csvPathMessage    = "Path to a csv file is required"
	scriptPathMessage = "Path to a script file is required"
)

func main() {
	var script string
	var inPlace bool
	var help bool

	flag.StringVar(&script, "s", "", "path to a script file")
	flag.BoolVar(&inPlace, "i", false, "update CSV file in place")
	flag.BoolVar(&help, "h", false, "usage")
	flag.Parse()

	if help {
		_, _ = os.Stdout.WriteString(usageMessage + "\n")
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) == 0 {
		_, _ = os.Stderr.WriteString(
			strings.Join([]string{csvPathMessage, usageMessage, ""}, "\n"),
		)
		os.Exit(1)
	}

	if script == "" {
		_, _ = os.Stderr.WriteString(
			strings.Join([]string{scriptPathMessage, usageMessage, ""}, "\n"),
		)
		os.Exit(1)
	}

	// Read and parse CSV file
	csvFile, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.LazyQuotes = true
	csvReader.TrimLeadingSpace = true

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
		os.Exit(1)
	}

	// Read and parse script file
	scriptFile, err := os.Open(script)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening script file: %v\n", err)
		os.Exit(1)
	}
	defer scriptFile.Close()

	program, err := evaluator.ParseProgram(scriptFile, script)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing script: %v\n", err)
		os.Exit(1)
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating script: %v\n", err)
		os.Exit(1)
	}

	// Output result as CSV to stdout
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, row := range result {
		if err := writer.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing CSV output: %v\n", err)
			os.Exit(1)
		}
	}
}
