package markdown

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

func readTable(reader io.Reader) ([][]string, error) {
	input, err := io.ReadAll(reader)
	if err != nil {
		return nil, ErrReadMD(err)
	}

	lines := strings.Split(string(input), "\n")

	output := make([][]string, len(lines))

	for i := range lines {
		rawFields := strings.Split(string(lines[i]), "|")
		output[i] = make([]string, len(rawFields)-2)
		for j := 1; j < len(rawFields)-1; j++ {
			output[i][j-1] = strings.TrimSpace(rawFields[j])
		}
	}

	columns := len(output[0])
	for i := range len(output) - 1 {
		line := output[i+1]
		if columns != len(line) && i == 0 {
			return nil, fmt.Errorf("malformed table header")
		}
		if columns != len(line) {
			return nil, fmt.Errorf("malformed table body line %d", i)
		}
	}

	output = slices.Delete(output, 1, 2)
	return output, nil
}
