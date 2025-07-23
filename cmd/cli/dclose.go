package main

import (
	"io"
	"log"
	"os"
)

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func dremove(f *os.File) {
	dclose(f)

	if err := os.Remove(f.Name()); err != nil {
		log.Fatal(err)
	}
}
