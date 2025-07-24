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

	// Open CSV source
	csvReader := os.Stdin
	if config.Input != "" {
		file, err := os.Open(config.Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(file)
		csvReader = file
	}

	// Open script source
	scriptReader := os.Stdin
	if config.Script != "" {
		file, err := os.Open(config.Script)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening script file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(file)
		scriptReader = file
	}

	// Open CSV destination
	csvWriter := os.Stdout
	if config.Output != "" {
		file, err := os.Create(config.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(file)
		csvWriter = file
	}

	// For update in place, we'll need to write to a temp file first
	var tempFile *os.File
	if config.Input == config.Output {
		tempFile, err = os.CreateTemp("", "csvss_temp_*.csv")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
			os.Exit(1)
		}

		defer dremove(tempFile)
		csvWriter = tempFile
	}

	// Process CSV
	if err := processCSV(csvReader, scriptReader, csvWriter); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// For update replace original file
	if config.Input == config.Output {
		if err := os.Rename(tempFile.Name(), config.Input); err != nil {
			fmt.Fprintf(os.Stderr, "Error updating file: %v\n", err)
			os.Exit(1)
		}
	}
}
