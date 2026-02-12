package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
)

var (
	usageMessage = `Usage: tabula [OPTIONS]

Options:
  -i <file>    Input CSV path (default: stdin)
  -s <file>    Script file path (default: stdin)
  -e <code>    Execute code
  -o <file>    Output CSV file (default: stdout)
  -u           Update input CSV file in place
  -a           Align output
  -t           Sort statements topologically
  -h           Show this help

Examples:
	# CSV from file, script from stdin → stdout
  tabula -i data.csv

	# CSV from file, script from file → stdout
  tabula -i data.csv -s script.file

	# CSV from file, execute code directly → stdout
	tabula -i data.csv -e "let A1 = SUM(A2:A4)"

	# CSV from file, script from file → file
  tabula -i data.csv -s script.file -o output.csv

	# CSV from file, script from file → update in place
  tabula -s script.file -u data.csv

	# CSV from file, script from stdin → update in place
  tabula -u data.csv
`
	outputConflictMessage = "conflicting output flags: -o and -u cannot be used together"
	inputConflictMessage  = "either script or data has to be read from a file"
)

type Config struct {
	Script  string
	Execute string
	Name    string
	Input   string
	Output  string
	Align   bool
	Sort    bool
}

func (c *Config) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

func parseArgs() (*Config, error) {
	var script string
	var execute string
	var output string
	var input string
	var update string
	var align bool
	var sort bool
	var help bool

	flag.StringVar(&input, "i", "", "read CSV file")
	flag.StringVar(&script, "s", "", "path to a script file")
	flag.StringVar(&execute, "e", "", "execute code directly")
	flag.StringVar(&output, "o", "", "output CSV file")
	flag.StringVar(&update, "u", "", "update CSV file in place")
	flag.BoolVar(&align, "a", false, "Align CSV output")
	flag.BoolVar(&sort, "t", false, "Sort statements topologically")
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

	// Check conflicting script flags
	if script != "" && execute != "" {
		return nil, errors.New("conflicting script flags: -s and -e cannot be used together")
	}

	// Basic validation - need either input file or script source
	if input == "" && script == "" && execute == "" {
		return nil, errors.New(inputConflictMessage)
	}

	config := Config{
		Script:  script,
		Execute: execute,
		Input:   input,
		Output:  output,
		Align:   align,
		Sort:    sort,
	}

	if script != "" {
		config.Name = script
	}

	if execute != "" {
		config.Name = "<inline>"
	}

	return &config, nil
}
