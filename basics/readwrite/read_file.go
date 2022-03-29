package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./files/input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("An error occurred on opening the input file\n")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("An error occurred on closing the input file\n")
			return
		}
	}(file)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("%s", line)
	}
}
