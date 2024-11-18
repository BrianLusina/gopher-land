package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// cat takes in a bufio Reader and prints out the contents of the file
func cat(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		_, err = fmt.Fprintf(os.Stdout, "%s", line)
		if err != nil {
			return
		}
	}
}

func using_buffer() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))

		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			if err != nil {
				continue
			}
			continue
		}
		cat(bufio.NewReader(file))
	}
}
