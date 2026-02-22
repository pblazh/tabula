package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// setupOutputWriter configures the output destination and returns writer and cleanup function
func setupOutputWriter(config *Config) (io.Writer, func(), error) {
	noop := func() {}

	// Update in place - use memory buffer then write to original file
	if config.Input == config.Output && config.Input != "" {
		// Read original file permissions
		fileInfo, err := os.Stat(config.Input)
		if err != nil {
			return nil, noop, ErrOpenDataFile(err)
		}
		perm := fileInfo.Mode().Perm()

		var buffer bytes.Buffer

		cleanup := func() {
			// Write buffer contents to original file with original permissions
			if err := os.WriteFile(config.Input, buffer.Bytes(), perm); err != nil {
				fmt.Fprint(os.Stderr, ErrWriteDataOutput(err))
				os.Exit(1)
			}
		}

		return &buffer, cleanup, nil
	}

	// Write to specific output file
	if config.Output != "" {
		file, err := os.Create(config.Output)
		if err != nil {
			return nil, noop, ErrCreateOutputFile(err)
		}

		cleanup := func() {
			closeOrFatal(file)
		}

		return file, cleanup, nil
	}

	// Default: write to stdout
	return os.Stdout, noop, nil
}
