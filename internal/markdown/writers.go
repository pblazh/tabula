package markdown

import (
	"fmt"
	"io"
	"slices"
	"strings"
)

func WriteAligned(writer io.Writer, result [][]string, comments map[int]string) error {
	out, widths := alignResult(result)
	_, err := writeHeader(writer, out[0])
	if err != nil {
		return err
	}

	_, err = writeHeaderSeparator(writer, widths)
	if err != nil {
		return err
	}

	for i := range len(out) - 1 {
		_, err = writeRow(writer, out[i+1])
		if err != nil {
			return err
		}
	}

	return nil
}

func writeHeader(writer io.Writer, result []string) (int, error) {
	header := "| " + strings.Join(result, " | ") + " |\n"
	return writer.Write([]byte(header)) //nolint:wrapcheck
}

func writeHeaderSeparator(writer io.Writer, widths []int) (int, error) {
	var headerSeparator strings.Builder
	headerSeparator.Write([]byte("|"))
	for _, w := range widths {
		headerSeparator.Write([]byte(" "))
		headerSeparator.Write([]byte(strings.Repeat("-", w)))
		headerSeparator.Write([]byte(" |"))
	}
	headerSeparator.Write([]byte("\n"))
	return writer.Write([]byte(headerSeparator.String())) //nolint:wrapcheck
}

func writeRow(writer io.Writer, row []string) (int, error) {
	line := "| " + strings.Join(row, " | ") + " |\n"
	return writer.Write([]byte(line)) //nolint:wrapcheck
}

func alignResult(result [][]string) ([][]string, []int) {
	out := slices.Clone(result)
	var widths []int

	for i := range len(out[0]) {
		var width int
		for j := range len(out) {
			width = max(width, len(out[j][i]))
		}
		widths = append(widths, width)
		for j := range len(out) {
			template := fmt.Sprintf("%%-%ds", width)
			result[j][i] = fmt.Sprintf(template, out[j][i])
		}
	}
	return out, widths
}
