package main

import (
	"fmt"
)

func ErrCreateOutputFile(err error) error {
	return fmt.Errorf("cannot create output file: %w", err)
}

func ErrOpenDataFile(err error) error {
	return fmt.Errorf("cannot open data file: %w", err)
}

func ErrSeekDataFile(err error) error {
	return fmt.Errorf("cannot seek data file: %w", err)
}

func ErrReopenDataFile(err error) error {
	return fmt.Errorf("cannot reopen data file: %w", err)
}

func ErrReadStdin(err error) error {
	return fmt.Errorf("cannot read stdin: %w", err)
}

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

func ErrWriteDataOutput(err error) error {
	return fmt.Errorf("cannot write output: %w", err)
}
