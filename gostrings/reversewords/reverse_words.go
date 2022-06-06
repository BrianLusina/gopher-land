package reversewords

import (
	"strings"
)

func reverseWords(phrase string) string {
	words := strings.Split(phrase, ".")
	output := make([]string, len(words))

	for i := range words {
		word := words[len(words)-1-i]
		output[i] = word
	}

	return strings.Join(output, ".")
}
