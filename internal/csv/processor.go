package csv

import (
	"fmt"
	"io"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/evaluator"
)

func Process(
	config *Config,
	csvReader io.Reader,
) ([][]string, map[int]string, error) {
	records, comments, embedded, err := Read(csvReader)
	if err != nil {
		return nil, nil, ErrReadCSV(err)
	}

	scriptReader, err := setupScriptReader(config, embedded)
	if err != nil {
		return nil, nil, ErrReadCSV(err)
	}

	program, identifiers, err := evaluator.ParseProgram(scriptReader, config.Input)
	if err != nil {
		return nil, nil, fmt.Errorf("%s", err)
	}
	records = evaluator.EnsureProgramDimensions(identifiers, records)

	// Sort program topologically if Sort flag is set
	if config.Sort {
		program, err = ast.SortProgram(program)
		if err != nil {
			return nil, nil, ErrSortScriptStatements(err)
		}
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return nil, nil, ErrEvaluateScript(config.Name, err)
	}

	return result, comments, nil
}
