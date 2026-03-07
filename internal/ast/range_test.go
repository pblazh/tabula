package ast

import (
	"reflect"
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
			input:  `A1`,
			column: 0,
			row:    0,
		},
		{
			name:   "valid uppercase cell",
			input:  `Z1`,
			column: 25,
			row:    0,
		},
		{
			name:   "valid lowercase cell",
			input:  `a1`,
			column: 0,
			row:    0,
		},
		{
			name:   "valid mixed case cell",
			input:  `aB10`,
			column: 27,
			row:    9,
		},
		{
			name:   "valid multi-letter column",
			input:  `AA123`,
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

func TestExpandRange(t *testing.T) {
	testcases := []struct {
		name     string
		start    string
		end      string
		expected []string
		error    string
	}{
		{
			name:     "valid same",
			start:    "A1",
			end:      "A1",
			expected: []string{"A1"},
		},
		{
			name:     "valid column",
			start:    "A1",
			end:      "A3",
			expected: []string{"A1", "A2", "A3"},
		},
		{
			name:     "valid row",
			start:    "A1",
			end:      "C1",
			expected: []string{"A1", "B1", "C1"},
		},
		{
			name:     "valid rect",
			start:    "A1",
			end:      "C3",
			expected: []string{"A1", "B1", "C1", "A2", "B2", "C2", "A3", "B3", "C3"},
		},
		{
			name:     "valid flipped rect",
			start:    "C3",
			end:      "A1",
			expected: []string{"C3", "B3", "A3", "C2", "B2", "A2", "C1", "B1", "A1"},
		},
		{
			name:  "valid flipped rect",
			start: "x",
			end:   "y",
			error: "invalid range x:y",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ExpandRange(tc.start, tc.end)

			if tc.error == "" && err != nil {
				t.Errorf("Unexpected error %s", err)
			}
			if tc.error != "" && err == nil {
				t.Errorf("Expected error")
			}
			if tc.error != "" && err.Error() != tc.error {
				t.Errorf("Expected '%s' error, but got '%s'", tc.error, err.Error())
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected row=%v, got row=%v", tc.expected, result)
			}
		})
	}
}
