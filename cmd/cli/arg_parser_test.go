package main

import (
	"bytes"
	"flag"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		errMsg string
		config *Config
	}{
		// Valid combinations
		{
			name:   "script from file, csv from stdin, output to stdout",
			args:   []string{"csvss", "-s", "scriptFile"},
			config: &Config{Script: "scriptFile", Input: "", Output: ""},
		},
		{
			name:   "script from stdin, csv from file, output to stdout",
			args:   []string{"csvss", "-i", "csvFile"},
			config: &Config{Script: "", Input: "csvFile", Output: ""},
		},
		{
			name:   "script from file, csv from file, output to stdout",
			args:   []string{"csvss", "-s", "scriptFile", "-i", "csvFile"},
			config: &Config{Script: "scriptFile", Input: "csvFile", Output: ""},
		},
		{
			name:   "script from file, csv from file, output to file",
			args:   []string{"csvss", "-s", "scriptFile", "-o", "output.csv", "-i", "csvFile"},
			config: &Config{Script: "scriptFile", Input: "csvFile", Output: "output.csv"},
		},
		{
			name:   "script from file, csv from file, update in place",
			args:   []string{"csvss", "-s", "scriptFile", "-u", "csvFile"},
			config: &Config{Script: "scriptFile", Input: "csvFile", Output: "csvFile"},
		},
		{
			name:   "script from stdin, csv from file, update in place",
			args:   []string{"csvss", "-u", "csvFile"},
			config: &Config{Script: "", Input: "csvFile", Output: "csvFile"},
		},
		// Invalid combinations
		{
			name:   "conflicting output flags -o and -u",
			args:   []string{"csvss", "-s", "scriptFile", "-o", "output.csv", "-u", "csvFile"},
			errMsg: "conflicting output",
		},
		{
			name:   "update flag without CSV file",
			args:   []string{"csvss", "-u"},
			errMsg: "either script or data has to be read from a file",
		},
		{
			name:   "script flag without script file",
			args:   []string{"csvss", "-s"},
			errMsg: "either script or data has to be read from a file",
		},
		{
			name:   "input flag without input file",
			args:   []string{"csvss", "-i"},
			errMsg: "either script or data has to be read from a file",
		},
		{
			name:   "without flags",
			args:   []string{"csvss"},
			errMsg: "either script or data has to be read from a file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset flag package for each test
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			// Mock os.Args for this test
			oldArgs := os.Args
			os.Args = tt.args

			defer func() {
				os.Args = oldArgs
			}()

			// Create buffer to capture output
			var outputBuffer bytes.Buffer

			// Call parseArgs function with buffer
			config := parseArgs(&outputBuffer)

			if tt.errMsg != "" {
				// Check that error message was written to buffer
				if !strings.Contains(outputBuffer.String(), tt.errMsg) {
					t.Errorf("Expected output to contain %q, got %q", tt.errMsg, outputBuffer.String())
				}
				return
			}

			if config == nil {
				t.Error("Expected config but got nil")
				return
			}

			if !reflect.DeepEqual(config, tt.config) {
				t.Errorf("Expected config %+v, got %+v", tt.config, config)
			}
		})
	}
}
