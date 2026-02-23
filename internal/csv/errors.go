package csv

import (
	"fmt"
)

func ErrReadCSV(err error) error {
	return fmt.Errorf("failed to read CSV, %v", err)
}

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("failed to sort statements, %v", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	return fmt.Errorf("failed to evaluate %s, %v", scriptName, err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("failed to write CSV, %s", err)
}

func ErrWriteComments(err error) error {
	return fmt.Errorf("failed to write comments to, %s", err)
}
