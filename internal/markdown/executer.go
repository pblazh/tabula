package markdown

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/csv"
	"github.com/pblazh/tabula/internal/evaluator"
)

func execute(config *Config, data *chunk, script *chunk) ([][]string, map[int]string, error) {
	comments := make(map[int]string)
	scriptReader := strings.NewReader("")
	if script != nil && len(script.text) > 0 {
		scriptLines := script.text[1 : len(script.text)-1]
		scriptData := strings.Join(scriptLines, "\n")
		scriptReader = strings.NewReader(scriptData)
	}

	if data.kind == csvKind {
		dataLines := data.text[1 : len(data.text)-1]
		dataString := strings.Join(dataLines, "\n")
		dataReader := strings.NewReader(dataString)
		records, comments, embedded, err := csv.Read(dataReader)
		if err != nil {
			return nil, nil, fmt.Errorf("%s", err)
		}

		codeReader := io.MultiReader(strings.NewReader(embedded), scriptReader)

		result, err := executeChunk(config, codeReader, records)
		if err != nil {
			return nil, nil, err
		}
		return result, comments, nil
	}

	if data.kind == tableKind {
		dataLines := data.text
		dataString := strings.Join(dataLines, "\n")
		dataReader := strings.NewReader(dataString)
		records, err := readTable(dataReader)
		if err != nil {
			return nil, nil, err
		}

		result, err := executeChunk(config, scriptReader, records)
		if err != nil {
			return nil, nil, err
		}

		return result, comments, nil
	}

	return nil, nil, errors.New("unsupported chunk")
}

func executeChunk(
	config *Config,
	scriptReader io.Reader,
	records [][]string,
) ([][]string, error) {
	program, _, err := evaluator.ParseProgram(scriptReader, config.Input)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	// Sort program topologically if Sort flag is set
	if config.Sort {
		program, err = ast.SortProgram(program)
		if err != nil {
			return nil, ErrSortScriptStatements(err)
		}
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return result, nil
}
