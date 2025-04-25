package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	usageMessage      = "Usage: 'csvss -s ./script.file ./table.csv'"
	csvPathMessage    = "Path to a csv file is required"
	scriptPathMessage = "Path to a script file is required"
)

func main() {
	var script string
	var inPlace bool
	var help bool

	flag.StringVar(&script, "s", "", "path to a script file")
	flag.BoolVar(&inPlace, "i", false, "update CSV file in place")
	flag.BoolVar(&help, "h", false, "usage")
	flag.Parse()

	if help {
		_, _ = os.Stdout.WriteString(usageMessage + "\n")
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) == 0 {
		_, _ = os.Stderr.WriteString(
			strings.Join([]string{csvPathMessage, usageMessage, ""}, "\n"),
		)
		os.Exit(1)
	}

	if script == "" {
		_, _ = os.Stderr.WriteString(
			strings.Join([]string{scriptPathMessage, usageMessage, ""}, "\n"),
		)
		os.Exit(1)
	}

	csvFile, err := os.Open(args[0])
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	reader := csv.NewReader(csvFile)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	fmt.Println(records)
}
