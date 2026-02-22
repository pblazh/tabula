package main

import (
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
  -m           Parse input as markdown
  -h           Show this help
  -v           Show version

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

	# Update markdown file in place
  tabula -m -s script.tbl -u data.md
`
	outputConflictMessage = "conflicting output flags: -o and -u cannot be used together"
	inputConflictMessage  = "either script or data has to be read from a file"
)

func parseArgs() (*Config, error) {
	var script string
	var execute string
	var output string
	var input string
	var update string
	var markdown bool
	var align bool
	var sort bool
	var help bool
	var showVersion bool

	flag.StringVar(&input, "i", "", "read CSV file")
	flag.StringVar(&script, "s", "", "path to a script file")
	flag.StringVar(&execute, "e", "", "execute code directly")
	flag.StringVar(&output, "o", "", "output CSV file")
	flag.StringVar(&update, "u", "", "update CSV file in place")
	flag.BoolVar(&markdown, "m", false, "input file format \"markdown\" or \"csv\" (default)")
	flag.BoolVar(&align, "a", false, "Align CSV output")
	flag.BoolVar(&sort, "t", false, "Sort statements topologically")
	flag.BoolVar(&help, "h", false, "usage")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()

	if help {
		fmt.Println(usageMessage)
		return nil, nil //nolint:nilnil
	}

	if showVersion {
		fmt.Println(VERSION)
		return nil, nil //nolint:nilnil
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
		Script:   script,
		Execute:  execute,
		Input:    input,
		Output:   output,
		Align:    align,
		Sort:     sort,
		Markdown: markdown,
	}
	return &config, nil
}
