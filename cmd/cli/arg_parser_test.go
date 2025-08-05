package main

import (
	"flag"
	"os"
	"reflect"
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
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "", Output: "", Align: false, Sort: false},
		},
		{
			name:   "script from stdin, csv from file, output to stdout",
			args:   []string{"csvss", "-i", "csvFile"},
			config: &Config{Script: "", Execute: "", Name: "", Input: "csvFile", Output: "", Align: false, Sort: false},
		},
		{
			name:   "script from file, csv from file, output to stdout",
			args:   []string{"csvss", "-s", "scriptFile", "-i", "csvFile"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "", Align: false, Sort: false},
		},
		{
			name:   "script from file, csv from file, output to file",
			args:   []string{"csvss", "-s", "scriptFile", "-o", "output.csv", "-i", "csvFile"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "output.csv", Align: false, Sort: false},
		},
		{
			name:   "script from file, csv from file, update in place",
			args:   []string{"csvss", "-s", "scriptFile", "-u", "csvFile"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "csvFile", Align: false, Sort: false},
		},
		{
			name:   "script from stdin, csv from file, update in place",
			args:   []string{"csvss", "-u", "csvFile"},
			config: &Config{Script: "", Execute: "", Name: "", Input: "csvFile", Output: "csvFile", Align: false, Sort: false},
		},
		{
			name:   "script from file, csv from file, aligned output",
			args:   []string{"csvss", "-s", "scriptFile", "-i", "csvFile", "-a"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "", Align: true, Sort: false},
		},
		{
			name:   "script from file, csv from file, output to file, aligned",
			args:   []string{"csvss", "-s", "scriptFile", "-i", "csvFile", "-o", "output.csv", "-a"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "output.csv", Align: true, Sort: false},
		},
		{
			name:   "script from file, csv from file, update in place, aligned",
			args:   []string{"csvss", "-s", "scriptFile", "-u", "csvFile", "-a"},
			config: &Config{Script: "scriptFile", Execute: "", Name: "scriptFile", Input: "csvFile", Output: "csvFile", Align: true, Sort: false},
		},
		{
			name:   "execute inline code, csv from file, output to stdout",
			args:   []string{"csvss", "-e", "sum(amount)", "-i", "csvFile"},
			config: &Config{Script: "", Execute: "sum(amount)", Name: "<inline>", Input: "csvFile", Output: "", Align: false, Sort: false},
		},
		{
			name:   "execute inline code, csv from file, output to file",
			args:   []string{"csvss", "-e", "sum(amount)", "-i", "csvFile", "-o", "output.csv"},
			config: &Config{Script: "", Execute: "sum(amount)", Name: "<inline>", Input: "csvFile", Output: "output.csv", Align: false, Sort: false},
		},
		// Invalid combinations
		{
			name:   "conflicting output flags -o and -u",
			args:   []string{"csvss", "-s", "scriptFile", "-o", "output.csv", "-u", "csvFile"},
			errMsg: "conflicting output flags: -o and -u cannot be used together",
		},
		{
			name:   "conflicting script flags -s and -e",
			args:   []string{"csvss", "-s", "scriptFile", "-e", "sum(amount)", "-i", "csvFile"},
			errMsg: "conflicting script flags: -s and -e cannot be used together",
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

			// Call parseArgs function with buffer
			config, err := parseArgs()

			if tt.errMsg != "" {
				// Check that error message was written to buffer
				if err.Error() != tt.errMsg {
					t.Errorf("Expected output to contain %q, got %q", tt.errMsg, err.Error())
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
