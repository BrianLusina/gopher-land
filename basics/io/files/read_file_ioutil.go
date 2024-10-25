package main

import (
	"fmt"
	"os"
)

func read_file_ioutil() {
	buf, err := os.ReadFile("./files/products.txt")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		if err != nil {
			return
		}
	}
	fmt.Printf("%s\n", string(buf))
}
