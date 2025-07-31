package main

import (
	"bufio"
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

	// For update in place, we'll need to write to a temp file first
	csvWriter := os.Stdout
	if config.Input == config.Output {
		tempFile, err := os.CreateTemp("", "csvss_temp_*.csv")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
			os.Exit(1)
		}

		// For update replace original file
		defer func() {
			if err := os.Rename(tempFile.Name(), config.Input); err != nil {
				fmt.Fprintf(os.Stderr, "Error updating file: %v\n", err)
				os.Exit(1)
			}
		}()

		csvWriter = tempFile
	}

	if config.Input != config.Output && config.Output != "" {
		// Open CSV destination
		file, err := os.Create(config.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer dclose(file)
		csvWriter = file
	}

	var comments map[int]string
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

		scriptComment, c := readComments(csvReader)
		comments = c

		_, err = file.Seek(0, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
			os.Exit(1)
		}

		if scriptComment != "" && config.Script == "" {
			config.Script = path.Join(path.Dir(config.Input), scriptComment)
		}
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

	// Process CSV
	if err := processCSV(config, scriptReader, csvReader, csvWriter, comments); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func readComments(f io.Reader) (string, map[int]string) {
	prefix := "#"
	scriptPrefix := "#csvss:"
	scanner := bufio.NewScanner(f)
	var i int

	comments := make(map[int]string)
	var script string

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, scriptPrefix) {
			script = trimmed[len(scriptPrefix):]
		}
		if strings.HasPrefix(trimmed, prefix) {
			comments[i] = trimmed
		}
		i++
	}
	return script, comments
}
