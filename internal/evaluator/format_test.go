package evaluator

import (
	"testing"
)

func TestCleanFormat(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "string with left alignment and width",
			input:    "%-15s",
			expected: "%s",
		},
		{
			name:     "string with width",
			input:    "%9s",
			expected: "%s",
		},
		{
			name:     "float with width",
			input:    "%9f",
			expected: "%f",
		},
		{
			name:     "float with width and precision",
			input:    "%6.2f",
			expected: "%f",
		},
		{
			name:     "integer with width",
			input:    "%5d",
			expected: "%d",
		},
		{
			name:     "integer with zero padding and width",
			input:    "%05d",
			expected: "%d",
		},
		{
			name:     "float with plus sign, width and precision",
			input:    "%+8.3f",
			expected: "%f",
		},
		{
			name:     "format without width",
			input:    "%s",
			expected: "%s",
		},
		{
			name:     "format with only precision",
			input:    "%.2f",
			expected: "%f",
		},
		{
			name:     "complex format with prefix and suffix",
			input:    "Value: %10.2f units",
			expected: "Value: %f units",
		},
		{
			name:     "multiple formats in string",
			input:    "%5d %8.2f %-10s",
			expected: "%d %f %s",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := cleanFormat(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}

