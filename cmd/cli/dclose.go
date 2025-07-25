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

func dremove(f string) {
	if err := os.Remove(f); err != nil {
		log.Fatal(err)
	}
}
