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
		if ch.kind == csvKind {
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
			if config.Align {
				err = csv.WriteAligned(&sb, data, comments)
			} else {
				err = csv.Write(&sb, data, comments)
			}
			if err != nil {
				return ErrWriteCSV(err)
			}

			lines := []string{ch.text[0]}
			lines = append(lines, strings.Split(strings.TrimSpace(sb.String()), "\n")...)
			lines = append(lines, ch.text[len(ch.text)-1])

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
