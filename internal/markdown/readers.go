package markdown

import (
	"encoding/csv"
	"io"
)

func readMarkdown(input io.Reader) ([][]string, error) {
	reader := csv.NewReader(input)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.Comment = '#'

	return reader.ReadAll()
}
