// Package csv is meant to read and process table data from markdown
package csv

import (
	"encoding/json"
)

type Config struct {
	Align   bool   // format output with alignment
	Execute string // script content to execute
	Input   string
	Name    string
	Output  string
	Script  string
	Sort    bool
}

func (c Config) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
