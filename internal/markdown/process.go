package markdown

import (
	"fmt"
	"io"
	"strings"

	"github.com/pblazh/tabula/internal/csv"
)

func Process(
	config *Config,
	reader io.Reader,
	writer io.Writer,
) error {
	chunks, err := parse(reader)
	if err != nil {
		return err
	}

	var result []chunk

	for i := range chunks {
		ch := chunks[i]
		write := csv.Write
		wrap := wrapCodeBlock

		if ch.kind == csvKind && config.Align {
			write = csv.WriteAligned
		}

		if ch.kind == tableKind {
			write = WriteAligned
			wrap = wrapTable
		}

		if ch.kind == csvKind || ch.kind == tableKind {
			data, comments, err := execute(config, &ch, getScriptChunk(chunks, i))
			if err != nil {
				result = append(result, ch)
				result = append(
					result,
					chunk{
						kind: messageKind,
						text: []string{fmt.Sprintf("<!-- tabula: %s -->\n", err)},
					},
				)
				continue
			}

			// create a formatter depending on if align was requested
			var sb strings.Builder
			err = write(&sb, data, comments)
			if err != nil {
				return ErrWriteCSV(err)
			}

			lines := wrap(sb.String())

			result = append(result, chunk{kind: csvKind, text: lines})
			continue
		}
		result = append(result, ch)
	}

	for _, ch := range result {
		_, err = fmt.Fprintf(writer, "%s\n", ch)
		if err != nil {
			return err
		}
	}

	return nil
}

func wrapCodeBlock(code string) []string {
	lines := []string{"```csv"}
	lines = append(lines, strings.Split(strings.TrimSpace(code), "\n")...)
	lines = append(lines, "```")
	return lines
}

func wrapTable(code string) []string {
	return strings.Split(strings.TrimSpace(code), "\n")
}
