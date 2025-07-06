package ast

import "testing"

func TestParseCell(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expCol   string
		expRow   int
		expValid bool
	}{
		{
			name:     "valid uppercase cell",
			input:    "A1",
			expCol:   "A",
			expRow:   1,
			expValid: true,
		},
		{
			name:     "valid lowercase cell",
			input:    "a1",
			expCol:   "A", // should be converted to uppercase
			expRow:   1,
			expValid: true,
		},
		{
			name:     "valid mixed case cell",
			input:    "aB10",
			expCol:   "AB", // should be converted to uppercase
			expRow:   10,
			expValid: true,
		},
		{
			name:     "valid multi-letter column",
			input:    "AA123",
			expCol:   "AA",
			expRow:   123,
			expValid: true,
		},
		{
			name:     "invalid - letters only",
			input:    "ABC",
			expValid: false,
		},
		{
			name:     "invalid - numbers only",
			input:    "123",
			expValid: false,
		},
		{
			name:     "invalid - mixed format",
			input:    "A1B2",
			expValid: false,
		},
		{
			name:     "invalid - empty string",
			input:    "",
			expValid: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			col, row, valid := parseCell(tc.input)

			if valid != tc.expValid {
				t.Errorf("Expected valid=%v, got valid=%v", tc.expValid, valid)
				return
			}

			if tc.expValid {
				if col != tc.expCol {
					t.Errorf("Expected column='%s', got column='%s'", tc.expCol, col)
				}
				if row != tc.expRow {
					t.Errorf("Expected row=%d, got row=%d", tc.expRow, row)
				}
			}
		})
	}
}

