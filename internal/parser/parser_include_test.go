package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pblazh/tabula/internal/lexer"
)

func TestIncludeBasic(t *testing.T) {
	tmpDir := t.TempDir()

	// Create included file
	includedFile := filepath.Join(tmpDir, "included.tbl")
	includedContent := "let B1 = 42; let B2 = 100;"
	err := os.WriteFile(includedFile, []byte(includedContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create main file
	mainFile := filepath.Join(tmpDir, "main.tbl")
	mainContent := `
#include "included.tbl";
let A1 = B1 + B2;
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Parse
	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	// Should have 3 statements: B1, B2, A1
	if len(program) != 3 {
		t.Errorf("Expected 3 statements, got %d", len(program))
	}
}

func TestIncludeDuplicate(t *testing.T) {
	tmpDir := t.TempDir()

	// Create included file
	includedFile := filepath.Join(tmpDir, "included.tbl")
	err := os.WriteFile(includedFile, []byte("let B1 = 42;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create main file that includes same file twice
	mainFile := filepath.Join(tmpDir, "main.tbl")
	mainContent := `
#include "included.tbl";
#include "included.tbl";
let A1 = B1;
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	// Should have 2 statements (B1 once, A1 once)
	if len(program) != 2 {
		t.Errorf("Expected 2 statements (duplicate should be ignored), got %d", len(program))
	}
}

func TestIncludeCircular(t *testing.T) {
	tmpDir := t.TempDir()

	fileA := filepath.Join(tmpDir, "a.tbl")
	fileB := filepath.Join(tmpDir, "b.tbl")

	// A includes B
	err := os.WriteFile(fileA, []byte(`
#include "b.tbl";
let A1 = 1;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// B includes A (circular!)
	err = os.WriteFile(fileB, []byte(`
#include "a.tbl";
let B1 = 2;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(fileA)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, fileA)
	parser := New(lex)
	_, _, err = parser.Parse()

	if err == nil {
		t.Fatal("Expected circular dependency error, got nil")
	}

	if !strings.Contains(err.Error(), "circular include") {
		t.Errorf("Expected 'circular include' error, got: %v", err)
	}
}

func TestIncludeNotFound(t *testing.T) {
	tmpDir := t.TempDir()

	mainFile := filepath.Join(tmpDir, "main.tbl")
	mainContent := `
#include "nonexistent.tbl";
let A1 = 1;
`
	err := os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	_, _, err = parser.Parse()

	if err == nil {
		t.Fatal("Expected file not found error, got nil")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Expected 'not found' error, got: %v", err)
	}
}

func TestIncludeNested(t *testing.T) {
	tmpDir := t.TempDir()

	// Create nested includes: main -> a -> b -> c
	fileC := filepath.Join(tmpDir, "c.tbl")
	err := os.WriteFile(fileC, []byte("let C1 = 3;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fileB := filepath.Join(tmpDir, "b.tbl")
	err = os.WriteFile(fileB, []byte(`
#include "c.tbl";
let B1 = 2;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fileA := filepath.Join(tmpDir, "a.tbl")
	err = os.WriteFile(fileA, []byte(`
#include "b.tbl";
let A1 = 1;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	mainFile := filepath.Join(tmpDir, "main.tbl")
	err = os.WriteFile(mainFile, []byte(`
#include "a.tbl";
let M1 = 0;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	// Should have 4 statements: C1, B1, A1, M1
	if len(program) != 4 {
		t.Errorf("Expected 4 statements, got %d", len(program))
	}
}

func TestIncludeDiamond(t *testing.T) {
	tmpDir := t.TempDir()

	// Diamond dependency: main includes A and B, both include C
	fileC := filepath.Join(tmpDir, "c.tbl")
	err := os.WriteFile(fileC, []byte("let C1 = 3;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fileB := filepath.Join(tmpDir, "b.tbl")
	err = os.WriteFile(fileB, []byte(`
#include "c.tbl";
let B1 = 2;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	fileA := filepath.Join(tmpDir, "a.tbl")
	err = os.WriteFile(fileA, []byte(`
#include "c.tbl";
let A1 = 1;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	mainFile := filepath.Join(tmpDir, "main.tbl")
	err = os.WriteFile(mainFile, []byte(`#include "a.tbl";
#include "b.tbl";
let M1 = 0;
`), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	// Should have 4 statements: C1 (once), A1, B1, M1
	if len(program) != 4 {
		t.Errorf("Expected 4 statements (C1 included once only), got %d", len(program))
	}
}

func TestIncludeWithoutSemicolon(t *testing.T) {
	tmpDir := t.TempDir()

	includedFile := filepath.Join(tmpDir, "included.tbl")
	err := os.WriteFile(includedFile, []byte("let B1 = 42;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	mainFile := filepath.Join(tmpDir, "main.tbl")
	// Note: no semicolon after #include
	mainContent := `
#include "included.tbl"
let A1 = B1;
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error (semicolon should be optional): %v", err)
	}

	if len(program) != 2 {
		t.Errorf("Expected 2 statements, got %d", len(program))
	}
}

func TestIncludeRelativePath(t *testing.T) {
	tmpDir := t.TempDir()

	// Create subdirectory
	subDir := filepath.Join(tmpDir, "lib")
	err := os.Mkdir(subDir, 0o755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create included file in subdirectory
	includedFile := filepath.Join(subDir, "utils.tbl")
	err = os.WriteFile(includedFile, []byte("let B1 = 42;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create main file
	mainFile := filepath.Join(tmpDir, "main.tbl")
	mainContent := `
#include "lib/utils.tbl";
let A1 = B1;
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if len(program) != 2 {
		t.Errorf("Expected 2 statements, got %d", len(program))
	}
}

func TestIncludeMultipleFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create multiple included files
	file1 := filepath.Join(tmpDir, "file1.tbl")
	err := os.WriteFile(file1, []byte("let B1 = 1;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file2 := filepath.Join(tmpDir, "file2.tbl")
	err = os.WriteFile(file2, []byte("let C1 = 2;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file3 := filepath.Join(tmpDir, "file3.tbl")
	err = os.WriteFile(file3, []byte("let D1 = 3;\n"), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create main file that includes all three
	mainFile := filepath.Join(tmpDir, "main.tbl")
	mainContent := `
#include "file1.tbl";
#include "file2.tbl";
#include "file3.tbl";
let A1 = B1 + C1 + D1;
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0o644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	file, err := os.Open(mainFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer func() { _ = file.Close() }()

	lex := lexer.New(file, mainFile)
	parser := New(lex)
	program, _, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	// Should have 4 statements: B1, C1, D1, A1
	if len(program) != 4 {
		t.Errorf("Expected 4 statements, got %d", len(program))
	}
}
