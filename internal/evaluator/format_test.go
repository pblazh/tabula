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

func TestDetectPlaceholderType(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected int
	}{
		// Integer placeholders
		{
			name:     "integer decimal",
			input:    "%d",
			expected: intPlacehoder,
		},
		{
			name:     "integer octal",
			input:    "%o",
			expected: intPlacehoder,
		},
		{
			name:     "integer hexadecimal lowercase",
			input:    "%x",
			expected: intPlacehoder,
		},
		{
			name:     "integer hexadecimal uppercase",
			input:    "%X",
			expected: intPlacehoder,
		},
		{
			name:     "unsigned integer",
			input:    "%u",
			expected: intPlacehoder,
		},
		{
			name:     "integer with width",
			input:    "%5d",
			expected: intPlacehoder,
		},
		{
			name:     "integer with flags and width",
			input:    "%+10d",
			expected: intPlacehoder,
		},

		// Float placeholders
		{
			name:     "float f format",
			input:    "%f",
			expected: floatPlacehoder,
		},
		{
			name:     "float F format",
			input:    "%F",
			expected: floatPlacehoder,
		},
		{
			name:     "float e format",
			input:    "%e",
			expected: floatPlacehoder,
		},
		{
			name:     "float E format",
			input:    "%E",
			expected: floatPlacehoder,
		},
		{
			name:     "float g format",
			input:    "%g",
			expected: floatPlacehoder,
		},
		{
			name:     "float G format",
			input:    "%G",
			expected: floatPlacehoder,
		},
		{
			name:     "float with precision",
			input:    "%.2f",
			expected: floatPlacehoder,
		},
		{
			name:     "float with width and precision",
			input:    "%8.3f",
			expected: floatPlacehoder,
		},

		// Boolean placeholders
		{
			name:     "boolean",
			input:    "%t",
			expected: boolPlacehoder,
		},
		{
			name:     "boolean with width",
			input:    "%5t",
			expected: boolPlacehoder,
		},
		// String placeholders
		{
			name:     "string",
			input:    "%s",
			expected: stringPlacehoder,
		},
		{
			name:     "character",
			input:    "%c",
			expected: stringPlacehoder,
		},
		{
			name:     "string with width",
			input:    "%10s",
			expected: stringPlacehoder,
		},
		{
			name:     "string with left alignment",
			input:    "%-15s",
			expected: stringPlacehoder,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := detectPlaceholderType(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
