package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestMainWithStdinStdout(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		stdin  string
		stdout string
		stderr string
	}{
		{
			name:   "help flag",
			args:   []string{"-h"},
			stdout: usageMessage + "\n",
			stderr: "",
		},
		{
			name:   "conflicting output flags",
			args:   []string{"-s", "script.file", "-o", "output.csv", "-u", "input.csv"},
			stdout: "",
			stderr: `conflicting output flags: -o and -u cannot be used together
exit status 1
`,
		},
		{
			name:   "no arguments",
			args:   []string{},
			stdout: "",
			stderr: `either script or data has to be read from a file
exit status 1
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build command with go run
			cmd := []string{"go", "run", "."}
			cmd = append(cmd, tt.args...)

			// Execute command with stdin input
			var stdout, stderr bytes.Buffer

			// Use exec.Command to run go run with stdin/stdout
			goCmd := exec.Command(cmd[0], cmd[1:]...)
			goCmd.Stdin = strings.NewReader(tt.stdin)
			goCmd.Stdout = &stdout
			goCmd.Stderr = &stderr

			_ = goCmd.Run()

			stdoutStr := stdout.String()
			stderrStr := stderr.String()

			// Check stdout
			if stdoutStr != tt.stdout {
				t.Errorf("Expected stdout %q but got %q", tt.stdout, stdoutStr)
			}

			// Check stderr
			if stderrStr != tt.stderr {
				t.Errorf("Expected stderr %q but got %q", tt.stderr, stderrStr)
			}
		})
	}
}
