package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

func escapeCSVField(field string) string {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	_ = writer.Write([]string{field})
	writer.Flush()
	escaped := buf.String()
	return escaped[:len(escaped)-1] // remove trailing newline
}

func writeCompact(csvWriter io.Writer, result [][]string) error {
	writer := csv.NewWriter(csvWriter)
	defer writer.Flush()

	for _, row := range result {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing CSV output: %v", err)
		}
	}

	return nil
}

func writeAligned(csvWriter io.Writer, result [][]string) error {
	tb := new(tabwriter.Writer)
	tb.Init(csvWriter, 0, 0, 1, ' ', 0)

	var sb strings.Builder
	for _, row := range result {
		sb.Reset()
		for c, col := range row {
			sb.Write([]byte(escapeCSVField(strings.TrimSpace(col))))
			if c < len(row)-1 {
				sb.Write([]byte("\t, "))
			}
		}
		sb.Write([]byte("\n"))
		_, err := tb.Write([]byte(sb.String()))
		if err != nil {
			return err
		}
	}
	return tb.Flush()
}
