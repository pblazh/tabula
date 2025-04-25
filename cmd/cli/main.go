package main

import (
	"os"
)

var usageMessage = "Usage: 'csvss -s ./script.file ./table.csv'"

func main() {
	_, _ = os.Stdout.WriteString(usageMessage + "\n")
	os.Exit(0)
}
