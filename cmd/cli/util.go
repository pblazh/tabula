package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening file %s for copying from, %s", src, err)
	}
	defer closeOrFatal(source)

	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error opening file %s for copying to, %s", dst, err)
	}
	defer closeOrFatal(destination)

	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("error copying a file %s to %s, %s", src, dst, err)
	}

	return nil
}
