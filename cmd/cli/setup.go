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

// setupOutputWriter configures the output destination and returns writer and cleanup function
func setupOutputWriter(config *Config) (io.Writer, func(), error) {
	noop := func() {}

	// Update in place - write to temp file then rename
	if config.Input == config.Output && config.Input != "" {
		tempFile, err := os.CreateTemp("", "csvss_temp_*.csv")
		if err != nil {
			return nil, noop, ErrCreateTempFile(err)
		}

		cleanup := func() {
			_ = tempFile.Close()
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
			return nil, noop, ErrCreateOutputFile(err)
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
			return nil, nil, ErrOpenCSVFile(err)
		}
		defer dclose(file)

		// Extract comments and embedded script references
		scriptComment, comments := readComments(file)

		// Reset file position to beginning
		if _, err = file.Seek(0, 0); err != nil {
			return nil, nil, ErrSeekCSVFile(err)
		}

		// Use embedded script if no explicit script provided
		if scriptComment != "" && config.Script == "" && config.Execute == "" {
			config.Script = path.Join(path.Dir(config.Input), scriptComment)
		}

		// Re-open file for actual reading (since we need a fresh reader)
		csvFile, err := os.Open(config.Input)
		if err != nil {
			return nil, nil, ErrReopenCSVFile(err)
		}

		return csvFile, comments, nil
	}

	// Read from stdin
	stdinContent, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, nil, ErrReadStdin(err)
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
			return nil, ErrOpenScriptFile(err)
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
