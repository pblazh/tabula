package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := parseArgs()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	if config == nil {
		os.Exit(0)
	}

	// Setup CSV input reader
	csvReader, embedded, comments, err := setupCSVReader(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Setup script reader
	scriptReader, err := setupScriptReader(config, embedded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Setup output writer
	csvWriter, cleanup, err := setupOutputWriter(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	// Process CSV with script
	if err := processCSV(config, scriptReader, csvReader, csvWriter, comments); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
