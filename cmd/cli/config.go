package main

import "encoding/json"

type Config struct {
	Align    bool   // format output with alignment
	Execute  string // script content to execute
	Input    string
	Markdown bool
	Output   string
	Script   string
	Sort     bool
}

func (c Config) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
