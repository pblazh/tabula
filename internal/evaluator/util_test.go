package evaluator

import (
	"strings"
	"testing"
)

func TestGetDimensions(t *testing.T) {
	cases := []struct {
		identifiers []string
		width       int
		height      int
	}{
		{
			identifiers: []string{},
			width:       0,
			height:      0,
		},
		{
			identifiers: []string{"A1"},
			width:       1,
			height:      1,
		},
		{
			identifiers: []string{"A1", "A1"},
			width:       1,
			height:      1,
		},
		{
			identifiers: []string{"A1", "A2"},
			width:       1,
			height:      2,
		},
		{
			identifiers: []string{"A1", "C4"},
			width:       3,
			height:      4,
		},
	}

	for _, tc := range cases {
		t.Run("getProgramDimensions "+strings.Join(tc.identifiers, ", "), func(t *testing.T) {
			width, height := getProgramDimensions(tc.identifiers)

			if width != tc.width || height != tc.height {
				t.Errorf("expected %dx%d got %dx%d", tc.width, tc.height, width, height)
			}
		})
	}
}

// func ensureProgramDimensions(identifiers []string, records [][]string) [][]string {
func TestEnsureProgramDimensions(t *testing.T) {
	cases := []struct {
		identifiers []string
		records     [][]string
		width       int
		height      int
	}{
		{
			identifiers: []string{},
			records:     [][]string{},
			width:       0,
			height:      0,
		},
		{
			identifiers: []string{"A1", "C4"},
			records:     [][]string{},
			width:       3,
			height:      4,
		},
		{
			identifiers: []string{"A1"},
			records:     [][]string{{"0", "1", "2"}, {"a", "b", "c"}},
			width:       3,
			height:      2,
		},
	}

	for _, tc := range cases {
		t.Run("ensureProgramDimensions "+strings.Join(tc.identifiers, ", "), func(t *testing.T) {
			expanded := EnsureProgramDimensions(tc.identifiers, tc.records)

			height := len(expanded)
			width := 0
			if height > 0 {
				width = len(expanded[0])
			}

			if width != tc.width || height != tc.height {
				t.Errorf("expected %dx%d got %dx%d", tc.width, tc.height, width, height)
			}
		})
	}
}
