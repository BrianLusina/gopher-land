package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf, err := ioutil.ReadFile("./files/products.txt")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		if err != nil {
			return
		}
	}
	fmt.Printf("%s\n", string(buf))
}
