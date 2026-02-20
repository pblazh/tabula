package csv

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/pblazh/tabula/internal/ast"
)

func ensureProgramDimensions(identifiers []string, records [][]string) [][]string {
	for i, row := range records {
		for j, cel := range row {
			records[i][j] = strings.TrimSpace(cel)
		}
	}

	requiredWidth, requiredHeight := getProgramDimensions(identifiers)

	for i, row := range records {
		diff := requiredWidth - len(row)
		if diff > 0 {
			records[i] = append(records[i], make([]string, requiredWidth-len(row))...)
		}
	}
	for range requiredHeight - len(records) {
		records = append(records, make([]string, requiredWidth))
	}
	return records
}

func getProgramDimensions(identifiers []string) (int, int) {
	requiredWidth := 0
	requiredHeight := 0

	for _, id := range identifiers {
		if ast.IsCellIdentifier(id) {
			col, row := ast.ParseCell(id)

			if col > requiredWidth {
				requiredWidth = col
			}
			if row > requiredHeight {
				requiredHeight = row
			}
		}
	}

	return requiredWidth + 1, requiredHeight + 1
}

func escapeCSVField(field string) string {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	_ = writer.Write([]string{field})
	writer.Flush()
	escaped := buf.String()
	return escaped[:len(escaped)-1] // remove trailing newline
}
