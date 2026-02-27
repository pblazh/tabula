package csv

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func Read(input io.Reader) (records [][]string, comments map[int]string, script string, err error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return nil, nil, "", ErrReadCSV(err)
	}

	script, comments, err = readComments(bytes.NewReader(data))
	if err != nil {
		return nil, nil, "", err
	}

	reader := csv.NewReader(bytes.NewReader(data))
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.Comment = '#'

	records, err = reader.ReadAll()
	if err != nil {
		return nil, nil, "", ErrReadCSV(err)
	}

	return records, comments, script, nil
}

// readComments extracts comments and embedded script references and embedded script from CSV content
func readComments(f io.Reader) (string, map[int]string, error) {
	const (
		commentPrefix     = "#"
		tabulaEmbedPrefix = "#tabula"
	)

	scanner := bufio.NewScanner(f)
	comments := make(map[int]string)
	var script strings.Builder

	lineNum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, tabulaEmbedPrefix) {
			script.WriteString(line[len(tabulaEmbedPrefix):] + "\n")
		}

		// Store all comment lines
		if strings.HasPrefix(line, commentPrefix) {
			comments[lineNum] = line
		}

		lineNum++
	}

	return script.String(), comments, nil
}

// setupScriptReader configures script input source
func setupScriptReader(config *Config, embedded string) (io.Reader, error) {
	// Execute inline code
	if config.Execute != "" {
		return strings.NewReader(config.Execute), nil
	}

	// Use embedded script if available (from CSV comments)
	if embedded != "" {
		// Set config.Name to CSV file path so parser can resolve relative includes
		return strings.NewReader(embedded), nil
	}

	if config.Script != "" {
		// config.Script is a file path, read the file content
		file, err := os.Open(config.Script)
		if err != nil {
			return nil, fmt.Errorf("failed to open script file %s: %w", config.Script, err)
		}
		return file, nil
	}

	// Default: read script from stdin (this should not happen due to validation)
	return os.Stdin, nil
}

func HasEmbeddedScript(csv string) bool {
	embeddedRe := regexp.MustCompile(`#tabula`)
	return embeddedRe.MatchString(csv)
}
