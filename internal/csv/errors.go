package csv

import (
	"fmt"
)

func ErrParseScript(err error) error {
	return fmt.Errorf("error parsing script, %v", err)
}

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("error sorting script statements, %v", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	return fmt.Errorf("error evaluating script %s, %v", scriptName, err)
}

func ErrWriteCSVOutput(err error) error {
	return fmt.Errorf("error writing CSV output, %v", err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("failed to write csv, %s", err)
}

func ErrWriteComments(err error) error {
	return fmt.Errorf("failed to write comments to csv, %s", err)
}

func ErrWriteDataOutput(err error) error {
	return fmt.Errorf("error writing data output, %v", err)
}

func ErrReadCSV(err error) error {
	return fmt.Errorf("error reading CSV, %v", err)
}
