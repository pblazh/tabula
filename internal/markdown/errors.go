package markdown

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

func ErrReadCSV(err error) error {
	return fmt.Errorf("error reading CSV, %v", err)
}

func ErrReadMD(err error) error {
	return fmt.Errorf("error reading markdown, %v", err)
}

func ErrWriteCSV(err error) error {
	return fmt.Errorf("error writing CSV, %v", err)
}

func ErrWriteMD(err error) error {
	return fmt.Errorf("error writing markdown, %v", err)
}

func ErrProcessing(err error) error {
	return fmt.Errorf("failed to process markdown, %v", err)
}
