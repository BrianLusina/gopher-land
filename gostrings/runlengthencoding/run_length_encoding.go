package runlengthencoding

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// rle is RunLengthEncode
func rle(input string) string {
	var encoded string

	for len(input) > 0 {
		char := input[0]
		slen := len(input)
		input = strings.TrimLeft(input, string(char))
		if n := slen - len(input); n > 1 {
			encoded += strconv.Itoa(n)
		}
		encoded += string(char)
	}

	return encoded
}

// rleDecode decodes a string
func rleDecode(input string) string {
	var decoded string

	for len(input) > 0 {
		i := strings.IndexFunc(input, func(r rune) bool {
			return !unicode.IsDigit(r)
		})

		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(input[:i])
		}
		decoded += strings.Repeat(string(input[i]), n)
		input = input[i+1:]
	}

	return decoded
}

// RunLengthEncode encodes a string
func RunLengthEncode(input string) string {
	result := ""
	count := 1

	encode := func(count, idx int, s string) string {
		if count > 1 {
			return fmt.Sprintf("%d%c", count, s[idx-1])
		}
		return fmt.Sprintf("%c", s[idx-1])
	}

	for idx, char := range input {
		if idx != 0 {
			if char == rune(input[idx-1]) {
				count++
			} else {
				result += encode(count, idx, input)
				count = 1
			}
		}
	}

	if input != "" {
		result += encode(count, len(input), input)
	}

	return result
}

// RunlengthDecode decodes a string
func RunLengthDecode(input string) string {
	count := 1
	var stringCount, result string
	for _, c := range input {
		if _, err := strconv.Atoi(string(c)); err == nil {
			stringCount += string(c)
		} else {
			if stringCount != "" {
				count, _ = strconv.Atoi(stringCount)
			}
			for j := 0; j < count; j++ {
				result += string(c)
			}
			count = 1
			stringCount = ""
		}
	}
	return result
}
