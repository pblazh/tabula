package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
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

	// Setup output writer
	csvWriter, cleanup, err := setupOutputWriter(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	// Setup CSV input reader
	csvReader, comments, err := setupCSVReader(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Setup script reader
	scriptReader, err := setupScriptReader(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// Process CSV with script
	if err := processCSV(config, scriptReader, csvReader, csvWriter, comments); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// setupOutputWriter configures the output destination and returns writer and cleanup function
func setupOutputWriter(config *Config) (io.Writer, func(), error) {
	noop := func() {}

	// Update in place - write to temp file then rename
	if config.Input == config.Output && config.Input != "" {
		tempFile, err := os.CreateTemp("", "csvss_temp_*.csv")
		if err != nil {
			return nil, noop, fmt.Errorf("error creating temp file: %v", err)
		}

		cleanup := func() {
			tempFile.Close()
			if err := os.Rename(tempFile.Name(), config.Input); err != nil {
				fmt.Fprintf(os.Stderr, "Error updating file: %v\n", err)
				os.Exit(1)
			}
		}

		return tempFile, cleanup, nil
	}

	// Write to specific output file
	if config.Input != config.Output && config.Output != "" {
		file, err := os.Create(config.Output)
		if err != nil {
			return nil, noop, fmt.Errorf("error creating output file: %v", err)
		}

		cleanup := func() {
			dclose(file)
		}

		return file, cleanup, nil
	}

	// Default: write to stdout
	return os.Stdout, noop, nil
}

// setupCSVReader configures CSV data and comments reader
func setupCSVReader(config *Config) (io.Reader, map[int]string, error) {
	// Read from file
	if config.Input != "" {
		file, err := os.Open(config.Input)
		if err != nil {
			return nil, nil, fmt.Errorf("error opening CSV file: %v", err)
		}
		defer dclose(file)

		// Extract comments and embedded script references
		scriptComment, comments := readComments(file)

		// Reset file position to beginning
		if _, err = file.Seek(0, 0); err != nil {
			return nil, nil, fmt.Errorf("error seeking CSV file: %v", err)
		}

		// Use embedded script if no explicit script provided
		if scriptComment != "" && config.Script == "" && config.Execute == "" {
			config.Script = path.Join(path.Dir(config.Input), scriptComment)
		}

		// Re-open file for actual reading (since we need a fresh reader)
		csvFile, err := os.Open(config.Input)
		if err != nil {
			return nil, nil, fmt.Errorf("error reopening CSV file: %v", err)
		}

		return csvFile, comments, nil
	}

	// Read from stdin
	stdinContent, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading stdin: %v", err)
	}

	// Parse comments from stdin content
	_, comments := readComments(bytes.NewReader(stdinContent))

	// Create reader from the content
	csvReader := bytes.NewReader(stdinContent)

	return csvReader, comments, nil
}

// setupScriptReader configures script input source
func setupScriptReader(config *Config) (io.Reader, error) {
	// Execute inline code
	if config.Execute != "" {
		return strings.NewReader(config.Execute), nil
	}

	// Read script from file
	if config.Script != "" {
		file, err := os.Open(config.Script)
		if err != nil {
			return nil, fmt.Errorf("error opening script file: %v", err)
		}
		return file, nil
	}

	// Default: read script from stdin (this should not happen due to validation)
	return os.Stdin, nil
}

// readComments extracts comments and embedded script references from CSV content
func readComments(f io.Reader) (string, map[int]string) {
	const (
		commentPrefix = "#"
		scriptPrefix  = "#csvss:"
	)

	scanner := bufio.NewScanner(f)
	comments := make(map[int]string)
	var embeddedScript string
	lineNum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for embedded script reference
		if strings.HasPrefix(line, scriptPrefix) {
			embeddedScript = line[len(scriptPrefix):]
		}

		// Store all comment lines
		if strings.HasPrefix(line, commentPrefix) {
			comments[lineNum] = line
		}

		lineNum++
	}

	return embeddedScript, comments
}

