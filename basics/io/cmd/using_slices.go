package main

import (
	"flag"
	"fmt"
	"os"
)

// cat2 is a variation of cat, which takes in a file and prints its contents
func cat2(file *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch linesRead, err := file.Read(buf[:]); true {
		case linesRead < 0:
			_, err := fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			if err != nil {
				return
			}
			os.Exit(1)

		case linesRead == 0: // end of file (EOF) encountered
			return
		case linesRead > 0:
			if linesWritten, err2 := os.Stdout.Write(buf[0:linesRead]); linesWritten != linesRead {
				_, err := fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", err2.Error())
				if err != nil {
					return
				}
				os.Exit(1)
			}
		}
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat2(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))

		if file == nil {
			_, err := fmt.Fprintf(os.Stderr, "cat: can't open %s: error: %s\n", flag.Arg(i), err.Error())
			if err != nil {
				return
			}
			os.Exit(1)
		}

		cat2(file)

		err = file.Close()
		if err != nil {
			return
		}
	}
}
