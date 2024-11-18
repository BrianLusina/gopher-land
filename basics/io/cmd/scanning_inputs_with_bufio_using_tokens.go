package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func bufio_with_tokens() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some input: ")

	// scanner.Split(bufio.ScanWords)
	scanner.Split(AlphanumericWord)
	for scanner.Scan() {
		fmt.Printf("Token ->%s<-\n", scanner.Text())
	}
}

func AlphanumericWord(data []byte, atEOF bool) (int, []byte, error) {
	// Skip leading spaces
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:]) // Read a rune
		if !unicode.IsSpace(r) {
			break // Break if a non-space character has been found (beginning of a word)
		}
	}
	// Start reading a word
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:]) // Read a rune
		// Stop when you read something that is not alphanumeric
		if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
			if i == 0 {
				// if the first character is not a space (already skipped)
				// but it is not an alphanumeric character, print that character alone
				return i + width, data[start : i+width], nil
			} else {
				// Otherwise you just completed a word
				return i, data[start:i], nil
			}
		}
	}
	if atEOF && len(data) > 0 {
		return len(data), data[0:], nil
	}
	return 0, nil, nil
}
