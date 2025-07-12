package ast

import (
	"strings"
	"testing"
)

func TestParseCell(t *testing.T) {
	testcases := []struct {
		name   string
		input  string
		column int
		row    int
	}{
		{
			name:   "valid uppercase cell",
			input:  "A1",
			column: 0,
			row:    0,
		},
		{
			name:   "valid uppercase cell",
			input:  "Z1",
			column: 25,
			row:    0,
		},
		{
			name:   "valid lowercase cell",
			input:  "a1",
			column: 0,
			row:    0,
		},
		{
			name:   "valid mixed case cell",
			input:  "aB10",
			column: 27,
			row:    9,
		},
		{
			name:   "valid multi-letter column",
			input:  "AA123",
			column: 26,
			row:    122,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			col, row := ParseCell(tc.input)

			if col != tc.column {
				t.Errorf("Expected column='%d', got column='%d'", tc.column, col)
			}
			if row != tc.row {
				t.Errorf("Expected row=%d, got row=%d", tc.row, row)
			}

			cell := ToCell(col, row)
			if cell != strings.ToUpper(tc.input) {
				t.Errorf("Expected cell=%s, got row=%s", tc.input, cell)
			}
		})
	}
}
