package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	usageMessage = `Usage: tabula [OPTIONS]

Options:
  -i, --input <file>    Input CSV path (default: stdin)
  -s, --script <file>   Script file path (default: stdin)
  -e, --execute <code>  Execute code
  -o, --output <file>   Output CSV file (default: stdout)
  -u, --update          Update input CSV file in place
  -a, --align           Align output
  -t, --topologically   Sort statements topologically
  -m, --markdown        Parse input as markdown
  -h, --help            Show this help
  -v, --version         Show version

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

	flag.StringVar(&input, "i", "", "")
	flag.StringVar(&input, "input", "", "input CSV or MD file")

	flag.StringVar(&script, "s", "", "")
	flag.StringVar(&script, "script", "", "path to a script file")

	flag.StringVar(&execute, "e", "", "")
	flag.StringVar(&execute, "execute", "", "code to execute")

	flag.StringVar(&output, "o", "", "")
	flag.StringVar(&output, "output", "", "output file")

	flag.StringVar(&update, "u", "", "")
	flag.StringVar(&update, "update", "", "update file in place")

	flag.BoolVar(&markdown, "m", false, "")
	flag.BoolVar(&markdown, "markdown", false, "input/output file format is \"markdown\"")

	flag.BoolVar(&align, "a", false, "")
	flag.BoolVar(&align, "align", false, "align CSV output")

	flag.BoolVar(&sort, "t", false, "")
	flag.BoolVar(&sort, "topologically", false, "sort statements topologically")

	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "help")

	flag.BoolVar(&showVersion, "v", false, "")
	flag.BoolVar(&showVersion, "version", false, "version")
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
