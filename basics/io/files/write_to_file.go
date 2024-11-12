package main

import (
	"bufio"
	"fmt"
	"os"
)

// write_to_file writes content of a slice of bytes to a given file name
// this demonstrates that the closure of the file is deferred and the error ignored this is because
// this ensure the file buffer is written to disk with f.Sync() which could potentially return an error.
// NOTE: closing the file after a successful write does not guarantee that the file is persisted to disk, hence the use of f.Sync()
func write_to_file(filename string, content []byte) error {
	f, err := os.Open(filename)

	defer func() {
		_ = f.Close()
	}()
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return f.Sync()
}

func run_write_to_file() {
	outputFile, outputError := os.OpenFile("./files/output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Printf("An error occurred with file closing\n")
			return
		}
	}(outputFile)

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			return
		}
	}
	err := outputWriter.Flush()
	if err != nil {
		fmt.Printf("An error occurred flushing file\n")
		return
	}
}
