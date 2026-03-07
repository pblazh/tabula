package csv

import (
	"fmt"
)

func ErrReadCSV(err error) error {
	return fmt.Errorf("cannot read CSV: %w", err)
}

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("cannot sort statements: %w", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	if scriptName == "" {
		return err
	}
	return fmt.Errorf("%s: %w", scriptName, err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("cannot write CSV: %w", err)
}

func ErrWriteComments(err error) error {
	return fmt.Errorf("cannot write comments: %w", err)
}
