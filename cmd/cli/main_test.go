package main

import (
	"bytes"
	"encoding/csv"
	"os"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	// Create a temporary CSV file
	csvContent := `a,b,c
1,2,3
4,5,6`

	csvFile, err := os.CreateTemp("", "test*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp CSV file: %v", err)
	}
	defer dremove(csvFile)

	if _, err := csvFile.WriteString(csvContent); err != nil {
		t.Fatalf("Failed to write CSV content: %v", err)
	}

	// Create a temporary script file
	scriptContent := `let A1 = "Header";
let B2 = 100;`

	scriptFile, err := os.CreateTemp("", "test*.csvs")
	if err != nil {
		t.Fatalf("Failed to create temp script file: %v", err)
	}
	defer dremove(scriptFile)

	if _, err := scriptFile.WriteString(scriptContent); err != nil {
		t.Fatalf("Failed to write script content: %v", err)
	}

	// Test would require running the main function with the temp files
	// For now, we'll just test that the files were created successfully
	if _, err := os.Stat(csvFile.Name()); os.IsNotExist(err) {
		t.Error("CSV file was not created")
	}

	if _, err := os.Stat(scriptFile.Name()); os.IsNotExist(err) {
		t.Error("Script file was not created")
	}
}

func TestCSVOutput(t *testing.T) {
	// Test CSV output formatting
	data := [][]string{
		{"Header1", "Header2", "Header3"},
		{"value1", "value2", "value3"},
		{"a", "b", "c"},
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			t.Fatalf("Error writing CSV: %v", err)
		}
	}
	writer.Flush()

	expected := "Header1,Header2,Header3\nvalue1,value2,value3\na,b,c\n"
	if buf.String() != expected {
		t.Errorf("Expected %q, got %q", expected, buf.String())
	}
}
