package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
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
		{
			name: "succcess",
			args: []string{"-i", "../../examples/basic/file.csv", "-s", "../../examples/basic/script.csvs"},
			stdout: `# Header
Full Name,Age,Grade
"Dow, Bob",25,170
"Dow, Alice",30,184
#csvss:./script.csvs
`,
			stderr: "",
		},
		{
			name: "script path from CSV comment",
			args: []string{"-i", "../../examples/basic/file.csv"},
			stdout: `# Header
Full Name,Age,Grade
"Dow, Bob",25,170
"Dow, Alice",30,184
#csvss:./script.csvs
`,
			stderr: "",
		},
		{
			name: "align flag",
			args: []string{"-i", "../../examples/basic/file.csv", "-a"},
			stdout: `# Header
Full Name    , Age , Grade
"Dow, Bob"   , 25  , 170
"Dow, Alice" , 30  , 184
#csvss:./script.csvs
`,
			stderr: "",
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

func TestUpdateInPlace(t *testing.T) {
	// Create a temporary CSV file
	tempDir := os.TempDir()

	// Create test CSV file (copy the working example)
	csvFile := filepath.Join(tempDir, "test.csv")
	csvContent := "1, 2, 0\n"
	err := os.WriteFile(csvFile, []byte(csvContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test CSV file: %v", err)
	}
	defer dremove(csvFile)

	// Create test script file (copy the working example)
	scriptFile := filepath.Join(tempDir, "script.csvs")
	scriptContent := "let C1 = A1 + B1;"

	err = os.WriteFile(scriptFile, []byte(scriptContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test script file: %v", err)
	}
	defer dremove(scriptFile)

	// Execute command with -u flag
	cmd := exec.Command("go", "run", ".", "-s", scriptFile, "-u", csvFile)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
	}

	// Read the updated file content
	updatedContent, err := os.ReadFile(csvFile)
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}

	// Expected output after running the script (matching the working example)
	expectedContent := "1,2,3\n"

	if string(updatedContent) != expectedContent {
		t.Errorf("Expected file content:\n%s\nBut got:\n%s", expectedContent, string(updatedContent))
	}

	// Ensure stdout is empty (since we're updating in place)
	if stdout.String() != "" {
		t.Errorf("Expected empty stdout but got: %q", stdout.String())
	}

	// Ensure stderr is empty (no errors)
	if stderr.String() != "" {
		t.Errorf("Expected empty stderr but got: %q", stderr.String())
	}
}
