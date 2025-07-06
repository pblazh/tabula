package parser

import (
	"strings"
	"testing"

	"github.com/pblazh/csvss/internal/lexer"
)

func TestParserIdentifiers(t *testing.T) {
	testcases := []struct {
		name        string
		input       string
		expectedIds []string
	}{
		{
			name:        "single identifier",
			input:       "something;",
			expectedIds: []string{"something"},
		},
		{
			name:        "multiple identifiers",
			input:       "a + b;",
			expectedIds: []string{"a", "b"},
		},
		{
			name:        "let statement identifier",
			input:       "let x = y;",
			expectedIds: []string{"x", "y"},
		},
		{
			name:        "call expression",
			input:       "SUM(a, b);",
			expectedIds: []string{"SUM", "a", "b"},
		},
		{
			name:        "range expression",
			input:       "a1:B2;",
			expectedIds: []string{"a1", "B2"},
		},
		{
			name:        "no identifiers",
			input:       "5 + 10;",
			expectedIds: []string{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			lex := lexer.New(strings.NewReader(tc.input), tc.name)
			parser := New(lex)
			_, identifiers, err := parser.Parse()
			if err != nil {
				t.Errorf("Unexpected error '%v'", err)
			}

			if len(identifiers) != len(tc.expectedIds) {
				t.Errorf("Expected %d identifiers, got %d: %v", len(tc.expectedIds), len(identifiers), identifiers)
				return
			}

			for i, expected := range tc.expectedIds {
				if identifiers[i] != expected {
					t.Errorf("Expected identifier %d to be '%s', got '%s'", i, expected, identifiers[i])
				}
			}
		})
	}
}
