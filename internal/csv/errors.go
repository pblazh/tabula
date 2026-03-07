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

func ErrWriteCSV(err error) error {
	return fmt.Errorf("cannot write CSV: %w", err)
}

func ErrWriteComments(err error) error {
	return fmt.Errorf("cannot write comments: %w", err)
}
