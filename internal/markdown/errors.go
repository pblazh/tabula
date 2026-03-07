package markdown

import (
	"fmt"
)

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("cannot sort statements: %w", err)
}

func ErrReadMD(err error) error {
	return fmt.Errorf("cannot read markdown: %w", err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("cannot write CSV: %w", err)
}

func ErrWriteMD(err error) error {
	return fmt.Errorf("cannot write markdown: %w", err)
}

func ErrProcessing(err error) error {
	return fmt.Errorf("cannot process markdown: %w", err)
}

func ErrProcessingTableLine(n int) error {
	return fmt.Errorf("malformed table at line %d", n)
}

var ErrProcessingTableHeader = fmt.Errorf("malformed table header")
