package main

import (
	"fmt"
)

func ErrCreateTempFile(err error) error {
	return fmt.Errorf("error creating temp file: %v", err)
}

func ErrCreateOutputFile(err error) error {
	return fmt.Errorf("error creating output file: %v", err)
}

func ErrOpenCSVFile(err error) error {
	return fmt.Errorf("error opening CSV file: %v", err)
}

func ErrSeekCSVFile(err error) error {
	return fmt.Errorf("error seeking CSV file: %v", err)
}

func ErrReopenCSVFile(err error) error {
	return fmt.Errorf("error reopening CSV file: %v", err)
}

func ErrReadStdin(err error) error {
	return fmt.Errorf("error reading stdin: %v", err)
}

func ErrOpenScriptFile(err error) error {
	return fmt.Errorf("error opening script file: %v", err)
}

func ErrReadCSV(err error) error {
	return fmt.Errorf("error reading CSV: %v", err)
}

func ErrParseScript(err error) error {
	return fmt.Errorf("error parsing script: %v", err)
}

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("error sorting script statements: %v", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	return fmt.Errorf("error evaluating script %s: %v", scriptName, err)
}

func ErrWriteCSVOutput(err error) error {
	return fmt.Errorf("error writing CSV output: %v", err)
}

