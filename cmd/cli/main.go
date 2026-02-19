package main

import (
	"fmt"
	"io"
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

	err = doProcessing(config, scriptReader, csvReader, comments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func doProcessing(
	config *Config,
	scriptReader io.Reader,
	csvReader io.Reader,
	comments map[int]string,
) error {
	csvWriter, cleanup, err := setupOutputWriter(config)
	if err != nil {
		return err
	}
	defer cleanup()

	return processCSV(config, scriptReader, csvReader, csvWriter, comments)
}
