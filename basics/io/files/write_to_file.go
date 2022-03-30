package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
