package markdown

import (
	"fmt"
)

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("failed to sort statements, %v", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	return fmt.Errorf("failed to evaluate %s, %v", scriptName, err)
}

func ErrReadMD(err error) error {
	return fmt.Errorf("failed to read markdown, %v", err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("failed to write CSV, %v", err)
}

func ErrWriteMD(err error) error {
	return fmt.Errorf("failed to write Markdown, %v", err)
}

func ErrProcessing(err error) error {
	return fmt.Errorf("failed to process markdown, %v", err)
}

func ErrProcessingTableLine(n int) error {
	return fmt.Errorf("failed to process table line %d", n)
}

var ErrProcessingTableHeader = fmt.Errorf("malformed table header")
