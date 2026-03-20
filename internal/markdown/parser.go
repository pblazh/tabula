package markdown

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

const (
	textKind = iota
	tableKind
	csvKind
	scriptKind
	messageKind
)

type chunk struct {
	kind int
	text []string
}

func (ch chunk) String() string {
	return strings.Join(ch.text, "\n")
}

func (ch chunk) isEmpty() bool {
	return len(strings.TrimSpace(strings.Join(ch.text, ""))) == 0
}

var (
	tableRowRg  = regexp.MustCompile(`^\|*([^|]+\|)+([^|]+\|)*$`)
	scriptBegin = regexp.MustCompile("^```tabula\\s*$")
	csvBegin    = regexp.MustCompile("^```csv\\s*$")
	scriptEnd   = regexp.MustCompile("^```\\s*$")
	message     = regexp.MustCompile(`^<!-- Tabula: .* -->$`)
)

func toMessage(msg string) string {
	return fmt.Sprintf("<!-- Tabula: %s -->", msg)
}

func parse(reader io.Reader) ([]chunk, error) {
	var chunks []chunk
	var buffer []string
	kind := textKind

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		isTable := tableRowRg.MatchString(line)
		isScriptStart := scriptBegin.MatchString(line)
		isCsvStart := csvBegin.MatchString(line)
		isEnd := scriptEnd.MatchString(line)
		isMessage := message.MatchString(line)

		switch {
		// --- Table start ---
		case kind != tableKind && isTable:
			appendChunk(&chunks, kind, buffer)
			buffer = []string{line}
			kind = tableKind

		// --- Table end ---
		case kind == tableKind && !isTable:
			appendChunk(&chunks, tableKind, buffer)
			buffer = []string{line}
			kind = detectKind(isScriptStart, isCsvStart, isMessage)

		// --- Script/CSV start ---
		case isScriptStart || isCsvStart:
			appendChunk(&chunks, kind, buffer)
			buffer = []string{line}
			kind = detectKind(isScriptStart, isCsvStart, isMessage)

		// --- Script/CSV end ---
		case (kind == scriptKind || kind == csvKind) && isEnd:
			buffer = append(buffer, line)
			appendChunk(&chunks, kind, buffer)
			buffer = nil
			kind = textKind

		// --- Error message ---
		case isMessage:
			appendChunk(&chunks, kind, buffer)
			buffer = []string{}
			kind = textKind

		// --- Default ---
		default:
			buffer = append(buffer, line)
		}
	}

	appendChunk(&chunks, kind, buffer)

	err := scanner.Err()
	if err != nil {
		return nil, ErrReadMD(err)
	}
	return chunks, nil
}

func getScriptChunk(chunks []chunk, i int) *chunk {
	ch := chunks[i]
	if ch.kind != csvKind && ch.kind != tableKind {
		return nil
	}

	for j := i + 1; j < len(chunks); j++ {
		if chunks[j].isEmpty() {
			continue
		}
		if chunks[j].kind == scriptKind {
			return &chunks[j]
		}
		return nil
	}
	return nil
}

func detectKind(isScript, isCSV, isMessage bool) int {
	switch {
	case isScript:
		return scriptKind
	case isCSV:
		return csvKind
	case isMessage:
		return messageKind
	default:
		return textKind
	}
}

func appendChunk(chunks *[]chunk, kind int, buffer []string) {
	if len(buffer) == 0 {
		return
	}
	*chunks = append(*chunks, chunk{
		kind: kind,
		text: buffer,
	})
}
