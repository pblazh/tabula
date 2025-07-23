package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	usageMessage = `Usage: csvss [OPTIONS]

Options:
  -s <file>    Script file path (default: stdin)
  -o <file>    Output CSV file (default: stdout)
  -u           Update input CSV file in place
  -h           Show this help

Examples:
	# CSV from file, script from stdin → stdout
  csvss -i data.csv

	# CSV from file, script from file → stdout
  csvss -i data.csv -s script.file

	# CSV from file, script from file → file
  csvss -i data.csv -s script.file -o output.csv

	# CSV from file, script from file → update in place
  csvss -s script.file -u data.csv

	# CSV from file, script from stdin → update in place
  csvss -u data.csv
`
	outputConflictMessage = "conflicting output flags: -o and -u cannot be used together"
	inputConflictMessage  = "either script or data has to be read from a file"
)

type Config struct {
	Script string
	Input  string
	Output string
}

func parseArgs() (*Config, error) {
	var script string
	var output string
	var input string
	var update string
	var help bool

	flag.StringVar(&script, "s", "", "path to a script file")
	flag.StringVar(&output, "o", "", "output CSV file")
	flag.StringVar(&input, "i", "", "read CSV file")
	flag.StringVar(&update, "u", "", "update CSV file in place")
	flag.BoolVar(&help, "h", false, "usage")
	flag.Parse()

	if help {
		fmt.Println(usageMessage)
		return nil, nil
	}

	// Check conflicting output flags
	if output != "" && update != "" {
		return nil, errors.New(outputConflictMessage)
	}

	// Handle update flag - when -u is used, it specifies both input and output
	if update != "" {
		input = update
		output = update
	}

	// Basic validation - need either input file or script file
	if input == "" && script == "" {
		return nil, errors.New(inputConflictMessage)
	}

	return &Config{
		Script: script,
		Input:  input,
		Output: output,
	}, nil
}
