package wordcount

import (
	re "regexp"
	"strings"
)

type Frequency map[string]int

func normalize(words string) string {
	r := re.MustCompile(`[^\w|']`)
	wordsLower := strings.ToLower(words)
	return r.ReplaceAllLiteralString(wordsLower, " ")
}

func WordCount(phrase string) Frequency {
	frequency := Frequency{}
	n := normalize(phrase)
	words := strings.Fields(n)

	for _, word := range words {
		w := strings.Trim(word, "'")
		frequency[w]++
	}
	return frequency
}
