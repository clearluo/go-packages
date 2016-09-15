package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and open 2 files, and store a file handle
	// for each in fw1 and fw2 variables respectively
	fw1, _ := os.Create("fileone.txt")
	fw2, _ := os.Create("filetwo.txt")

	// Defer func to close both file handles after
	// main function finishes execution
	defer func() {
		fw1.Close()
		fw2.Close()
	}()

	// Create an io.Writer slice that contains
	// fw1, fw2 as well as os.Stdout
	writers := []io.Writer{fw1, fw2, os.Stdout}

	// Create a new multiwriter that takes writers slice
	// as its parameter
	//
	// MultiWriter creates a writer that duplicates its writes to
	// all the provided writers, similar to the Unix tee(1) command.
	mw := io.MultiWriter(writers...)

	// Create a new byte slice of data to write
	// to multiwriter mw
	data := []byte("Hello, World!\n")

	// Write byte slice `data` to mw (both files and os.Stdout)
	n, err := mw.Write(data)

	// Check if there were any errors writing with multiwriter mw
	if err != nil {
		// If there was, log any errors
		log.Fatalln(err)
	}

	// Print out how many bytes were written using mw
	fmt.Println("MultiWriter mw printed", n,
		"bytes to 3 writers at the same time")
}
