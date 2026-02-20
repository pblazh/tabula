package csv

import (
	"io"

	"github.com/pblazh/tabula/internal/ast"
	"github.com/pblazh/tabula/internal/evaluator"
)

func Process(
	config *Config,
	csvReader io.Reader,
	csvWriter io.Writer,
) error {
	records, comments, embedded, err := Read(csvReader)
	if err != nil {
		return ErrReadCSV(err)
	}

	scriptReader, err := setupScriptReader(config, embedded)
	if err != nil {
		return ErrReadCSV(err)
	}

	program, identifiers, err := evaluator.ParseProgram(scriptReader, config.Input)
	records = ensureProgramDimensions(identifiers, records)

	if err != nil {
		return ErrParseScript(err)
	}

	// Sort program topologically if Sort flag is set
	if config.Sort {
		program, err = ast.SortProgram(program)
		if err != nil {
			return ErrSortScriptStatements(err)
		}
	}

	// Evaluate the program with CSV data
	result, err := evaluator.Evaluate(program, records)
	if err != nil {
		return ErrEvaluateScript(config.Name, err)
	}

	if config.Align {
		err = WriteAligned(csvWriter, result, comments)
		if err != nil {
			return ErrWriteDataOutput(err)
		}
		return nil
	}

	err = Write(csvWriter, result, comments)
	if err != nil {
		return ErrWriteDataOutput(err)
	}

	return nil
}
