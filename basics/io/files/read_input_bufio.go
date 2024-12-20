package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_input_bufio() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}
	fmt.Printf("The input was: %s\n", input)
}
