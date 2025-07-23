package main

import (
	"fmt"
	"log"
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
		csvReader, err := os.Open(config.Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(csvReader)
	}

	// Open CSV source
	scriptReader := os.Stdin
	if config.Script != "" {
		scriptReader, err := os.Open(config.Script)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening script file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(scriptReader)
	}

	// Open CSV destionation
	csvWriter := os.Stdout
	if config.Output != "" {
		csvWriter, err := os.Open(config.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(csvWriter)
	}

	// For update in place, we'll need to write to a temp file first
	var tempFile os.File
	if config.Input == config.Output {
		tempFile, err := os.CreateTemp("", "csvss_temp_*.csv")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
			os.Exit(1)
		}

		defer dclose(tempFile)
		defer func() {
			err := os.Remove(tempFile.Name())
			if err != nil {
				log.Fatal(err)
			}
		}()
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
