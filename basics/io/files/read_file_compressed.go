package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func read_compressed_file() {
	fName := "./files/example.json.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
