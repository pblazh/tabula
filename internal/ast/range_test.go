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
			name:   "valid uppercase cell",
			input:  "A1",
			expCol: "A",
			expRow: 1,
		},
		{
			name:   "valid lowercase cell",
			input:  "a1",
			expCol: "A", // should be converted to uppercase
			expRow: 1,
		},
		{
			name:   "valid mixed case cell",
			input:  "aB10",
			expCol: "AB", // should be converted to uppercase
			expRow: 10,
		},
		{
			name:   "valid multi-letter column",
			input:  "AA123",
			expCol: "AA",
			expRow: 123,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			col, row := parseCell(tc.input)

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

