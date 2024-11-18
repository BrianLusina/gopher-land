package main

import (
	"bufio"
	"fmt"
	"os"
)

func bufio_scan() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some input: ")
	if scanner.Scan() {
		fmt.Println("The input was", scanner.Text())
	}
}
