package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pblazh/tabula/internal/csv"
	"github.com/pblazh/tabula/internal/markdown"
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

	var dataReader io.Reader

	// Setup CSV input reader
	if config.Input == "" {
		dataReader = os.Stdin
	}

	if config.Input != "" {
		dataReader, err = os.Open(config.Input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}

	if config.Markdown {
		err = doMarkdownProcessing(config, dataReader)
	} else {
		err = doCsvProcessing(config, dataReader)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func doCsvProcessing(
	config *Config,
	reader io.Reader,
) error {
	writer, cleanup, err := setupOutputWriter(config)
	if err != nil {
		return err
	}
	defer cleanup()

	csvConfig := csv.Config{
		Align:   config.Align,
		Execute: config.Execute,
		Input:   config.Input,
		Output:  config.Output,
		Script:  config.Script,
		Sort:    config.Sort,
	}

	return csv.Process(&csvConfig, reader, writer) //nolint:wrapcheck
}

func doMarkdownProcessing(
	config *Config,
	reader io.Reader,
) error {
	writer, cleanup, err := setupOutputWriter(config)
	if err != nil {
		return err
	}
	defer cleanup()

	markdownConfig := markdown.Config{
		Align:   config.Align,
		Execute: config.Execute,
		Input:   config.Input,
		Output:  config.Output,
		Script:  config.Script,
		Sort:    config.Sort,
	}
	return markdown.Process(&markdownConfig, reader, writer) //nolint:wrapcheck
}
