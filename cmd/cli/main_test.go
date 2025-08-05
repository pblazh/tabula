package main

import (
	"bytes"
	"io/fs"
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

func TestExecuteInlineCode(t *testing.T) {
	scriptPath := filepath.Join("..", "..", "examples", "apartment", "script.csvs")
	inputPath := filepath.Join("..", "..", "examples", "apartment", "input.csv")
	outputPath := filepath.Join("..", "..", "examples", "apartment", "output.csv")

	// Read expected output
	input, err := os.ReadFile(inputPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}
	csvIn := strings.ReplaceAll(string(input), "#csvss:./script.csvs", "")
	// Read expected output
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}

	// Read expected output
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read expected output: %v", err)
	}
	csvOut := strings.ReplaceAll(string(output), "#csvss:./script.csvs", "")

	cmd := exec.Command("go", "run", ".", "-e", string(script), "-a")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = strings.NewReader(csvIn)

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
	}

	// Remove the script comment line from expected output since we're using -e flag
	// expectedOutput := strings.ReplaceAll(string(output), "#csvss:./script.csvs", "")

	// Normalize outputs for comparison
	expectedStr := normalizeOutput(csvOut)
	actualStr := normalizeOutput(stdout.String())

	if expectedStr != actualStr {
		t.Errorf("Apartment example with -e flag: output mismatch\nExpected:\n%s\n\nActual:\n%s", expectedStr, actualStr)
	}

	// Ensure stderr is empty (no errors)
	if stderr.String() != "" {
		t.Errorf("Expected empty stderr but got: %q", stderr.String())
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

func TestScriptPathFromCSVComment(t *testing.T) {
	tempDir := os.TempDir()

	// Create subdirectory structure
	subDir := filepath.Join(tempDir, "csvss_test_subdir")
	err := os.MkdirAll(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test subdirectory: %v", err)
	}
	defer dremove(subDir)

	tests := []struct {
		name           string
		csvPath        string
		scriptPath     string
		scriptComment  string
		csvContent     string
		scriptContent  string
		expectedOutput string
	}{
		{
			name:           "parent directory script reference",
			csvPath:        filepath.Join(subDir, "test.csv"),
			scriptPath:     filepath.Join(tempDir, "parent_script.csvs"),
			scriptComment:  "../parent_script.csvs",
			csvContent:     "A,B\n1,2\n#csvss:../parent_script.csvs\n",
			scriptContent:  "let A1 = \"ParentScript\"; let B1 = \"Modified\";",
			expectedOutput: "ParentScript,Modified\n1,2\n#csvss:../parent_script.csvs\n",
		},
		{
			name:           "same directory script reference",
			csvPath:        filepath.Join(subDir, "test2.csv"),
			scriptPath:     filepath.Join(subDir, "local_script.csvs"),
			scriptComment:  "./local_script.csvs",
			csvContent:     "A,B\n1,2\n#csvss:./local_script.csvs\n",
			scriptContent:  "let A1 = \"LocalScript\"; let B1 = \"Local\";",
			expectedOutput: "LocalScript,Local\n1,2\n#csvss:./local_script.csvs\n",
		},
		{
			name:           "relative path without dot prefix",
			csvPath:        filepath.Join(subDir, "test3.csv"),
			scriptPath:     filepath.Join(subDir, "simple_script.csvs"),
			scriptComment:  "simple_script.csvs",
			csvContent:     "A,B\n1,2\n#csvss:simple_script.csvs\n",
			scriptContent:  "let A1 = \"SimpleScript\"; let B1 = \"Simple\";",
			expectedOutput: "SimpleScript,Simple\n1,2\n#csvss:simple_script.csvs\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create CSV file
			err := os.WriteFile(tt.csvPath, []byte(tt.csvContent), 0644)
			if err != nil {
				t.Fatalf("Failed to create CSV file: %v", err)
			}
			defer dremove(tt.csvPath)

			// Create script file
			err = os.WriteFile(tt.scriptPath, []byte(tt.scriptContent), 0644)
			if err != nil {
				t.Fatalf("Failed to create script file: %v", err)
			}
			defer dremove(tt.scriptPath)

			// Execute command - only specify CSV file, let it find script from comment
			cmd := exec.Command("go", "run", ".", "-i", tt.csvPath)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err = cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
			}

			// Check output
			if stdout.String() != tt.expectedOutput {
				t.Errorf("Expected output:\n%s\nBut got:\n%s", tt.expectedOutput, stdout.String())
			}

			// Ensure stderr is empty (no errors)
			if stderr.String() != "" {
				t.Errorf("Expected empty stderr but got: %q", stderr.String())
			}
		})
	}
}

func TestExamples(t *testing.T) {
	// Get the project root directory (go up from cmd/cli to project root)
	examplesDir := filepath.Join("..", "..", "examples")

	// Walk through all example directories
	err := filepath.WalkDir(examplesDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root examples directory and README.md
		if path == examplesDir || !d.IsDir() {
			return nil
		}

		// Skip if this is a subdirectory of an example (not a direct example folder)
		rel, _ := filepath.Rel(examplesDir, path)
		if strings.Contains(rel, string(filepath.Separator)) {
			return nil
		}

		exampleName := d.Name()

		// Define required file paths
		inputFile := filepath.Join(path, "input.csv")
		outputFile := filepath.Join(path, "output.csv")

		// Check if all required files exist
		if !fileExists(inputFile) {
			t.Errorf("Example %s: missing input.csv", exampleName)
			return nil
		}
		if !fileExists(outputFile) {
			t.Errorf("Example %s: missing output.csv", exampleName)
			return nil
		}

		// Run the test for this example
		t.Run(exampleName, func(t *testing.T) {
			testExample(t, exampleName, inputFile, outputFile)
		})

		return nil
	})
	if err != nil {
		t.Fatalf("Failed to walk examples directory: %v", err)
	}
}

func testExample(t *testing.T, exampleName, inputFile, outputFile string) {
	// Read expected output
	expectedOutput, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read expected output file %s: %v", outputFile, err)
	}

	// Execute csvss command: csvss -i input.csv -s script.csvs
	actualOutput, err := executeCSVSSCommand(inputFile)
	if err != nil {
		t.Fatalf("Failed to execute csvss command for example %s: %v", exampleName, err)
	}

	// Normalize whitespace for comparison
	expectedStr := normalizeOutput(string(expectedOutput))
	actualStr := normalizeOutput(string(actualOutput))

	if expectedStr != actualStr {
		t.Errorf("Example %s: output mismatch\nExpected:\n%s\n\nActual:\n%s",
			exampleName, expectedStr, actualStr)
	}
}

func executeCSVSSCommand(inputFile string) ([]byte, error) {
	cmd := exec.Command("go", "run", ".", "-i", inputFile, "-a")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return stdout.Bytes(), nil
}

func normalizeOutput(s string) string {
	// Normalize line endings and trim whitespace
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	s = strings.TrimSpace(s)
	return s
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
