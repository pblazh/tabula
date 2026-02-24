package main

import (
	"fmt"
)

func ErrCreateOutputFile(err error) error {
	return fmt.Errorf("error creating output file, %v", err)
}

func ErrOpenDataFile(err error) error {
	return fmt.Errorf("error opening data file, %v", err)
}

func ErrSeekDataFile(err error) error {
	return fmt.Errorf("error seeking data file, %v", err)
}

func ErrReopenDataFile(err error) error {
	return fmt.Errorf("error reopening data file, %v", err)
}

func ErrReadStdin(err error) error {
	return fmt.Errorf("error reading stdin, %v", err)
}

func ErrReadCSV(err error) error {
	return fmt.Errorf("error reading CSV, %v", err)
}

func ErrSortScriptStatements(err error) error {
	return fmt.Errorf("error sorting script statements, %v", err)
}

func ErrEvaluateScript(scriptName string, err error) error {
	return fmt.Errorf("error evaluating script %s, %v", scriptName, err)
}

func ErrWriteDataOutput(err error) error {
	return fmt.Errorf("error writing data output, %v", err)
}
