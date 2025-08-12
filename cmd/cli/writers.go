package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
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

func writeCompact(csvWriter io.Writer, result [][]string, comments map[int]string) error {
	writer := csv.NewWriter(csvWriter)
	defer writer.Flush()

	lineNum := 0
	for _, row := range result {
		if comment, ok := comments[lineNum]; ok {
			writer.Flush()
			if err := writer.Error(); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(csvWriter, comment); err != nil {
				return err
			}
			lineNum++
		}
		if err := writer.Write(row); err != nil {
			return ErrWriteCSVOutput(err)
		}
		lineNum++
	}

	writer.Flush()

	return dumpComments(comments, lineNum, csvWriter)
}

func writeAligned(csvWriter io.Writer, result [][]string, comments map[int]string) error {
	var buf bytes.Buffer
	tb := new(tabwriter.Writer)
	tb.Init(&buf, 0, 0, 1, ' ', 0)

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
		if _, err := tb.Write([]byte(sb.String())); err != nil {
			return err
		}
	}
	if err := tb.Flush(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(&buf)
	var lineNum int
	for scanner.Scan() {
		if comment, ok := comments[lineNum]; ok {
			if _, err := fmt.Fprintln(csvWriter, comment); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(csvWriter, scanner.Text()); err != nil {
			return err
		}
		lineNum++
	}

	return dumpComments(comments, lineNum, csvWriter)
}

func dumpComments(comments map[int]string, lineNum int, w io.Writer) error {
	var remainingLines []int
	for lineNo := range comments {
		if lineNo >= lineNum {
			remainingLines = append(remainingLines, lineNo)
		}
	}

	// Sort the line numbers to maintain order
	sort.Ints(remainingLines)

	// Write comments in order
	for _, lineNo := range remainingLines {
		if _, err := fmt.Fprintln(w, comments[lineNo]); err != nil {
			return err
		}
	}
	return nil
}
