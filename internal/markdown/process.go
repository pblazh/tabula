package markdown

import (
	"fmt"
	"io"
	"strings"
	"sync"

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

	result := make([][]*chunk, len(chunks))
	var wg sync.WaitGroup

	for i := range chunks {
		ch := chunks[i]
		result[i] = make([]*chunk, 0)
		scriptChunk := getScriptChunk(chunks, i)

		if chunkNeedsProcessing(&ch, scriptChunk) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				output, err := processChunk(config, &ch, scriptChunk)
				if err != nil {
					fmt.Println(err)
				}
				result[i] = output
			}()
			continue
		}
		result[i] = append(result[i], &ch)
	}

	wg.Wait()

	for _, chs := range result {
		for _, ch := range chs {
			_, err = fmt.Fprintf(writer, "%s\n", ch)
			if err != nil {
				return ErrWriteMD(err)
			}
		}
	}

	return nil
}

func processChunk(config *Config, data, script *chunk) ([]*chunk, error) {
	var result []*chunk
	lines, comments, err := execute(config, data, script)
	if err != nil {
		result = append(result, data)
		result = append(
			result,
			&chunk{
				kind: messageKind,
				text: []string{toMessage(err.Error())},
			},
		)
		return result, nil
	}

	write := csv.Write
	wrap := wrapCodeBlock

	if data.kind == csvKind && config.Align {
		write = csv.WriteAligned
	}

	if data.kind == tableKind {
		write = WriteAligned
		wrap = wrapTable
	}
	// create a formatter depending on if align was requested
	var sb strings.Builder
	err = write(&sb, lines, comments)
	if err != nil {
		return nil, ErrWriteCSV(err)
	}

	output := wrap(sb.String())

	result = append(result, &chunk{kind: csvKind, text: output})
	return result, nil
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

func chunkNeedsProcessing(ch *chunk, script *chunk) bool {
	return (ch.kind == tableKind && script != nil) ||
		(ch.kind == csvKind && (script != nil || csv.HasEmbeddedScript(strings.Join(ch.text, "\n"))))
}
