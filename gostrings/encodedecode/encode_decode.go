package encodedecode

import (
	"fmt"
	"strconv"
)

func encode(words []string) string {
	encodedStr := ""
	for _, word := range words {
		encodedStr += fmt.Sprintf("%d#%s", len(word), word)
	}
	return encodedStr
}

func decode(encodedStr string) []string {
	result, i := []string{}, 0

	for i < len(encodedStr) {
		j := i
		for encodedStr[j] != '#' {
			j++
		}

		wordLength := encodedStr[i:j]
		length, err := strconv.Atoi(wordLength)
		if err != nil {
			panic(fmt.Sprintf("Cannot convert %s to integer", wordLength))
		}
		word := encodedStr[j+1 : j+1+length]
		result = append(result, word)
		i = j + 1 + length
	}

	return result
}
