package main

import (
	"io"
	"log"
	"os"
)

func closeOrFatal(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func removeOrFatal(f string) {
	if err := os.Remove(f); err != nil {
		log.Fatal(err)
	}
}
