package ast

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    []string{"A1"},
			expected: []string{"A1"},
		},
		{
			name:     "no duplicates",
			input:    []string{"A1", "B1", "C1"},
			expected: []string{"A1", "B1", "C1"},
		},
		{
			name:     "all duplicates",
			input:    []string{"A1", "A1", "A1"},
			expected: []string{"A1"},
		},
		{
			name:     "some duplicates",
			input:    []string{"B1", "A1", "C1", "A1", "B1"},
			expected: []string{"A1", "B1", "C1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input since removeDuplicates modifies it in-place
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			result := removeDependenciesDuplicates(inputCopy)

			if slices.Compare(result, tt.expected) != 0 {
				t.Errorf("removeDuplicates() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

