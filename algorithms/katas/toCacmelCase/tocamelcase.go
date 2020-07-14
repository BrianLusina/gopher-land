package toCacmelCase

import (
	"regexp"
	"strings"
)

// ToCamelCase takes in a string and returns camel case
func ToCamelCase(s string) string {
	words := regexp.MustCompile("[-_]").Split(s, -1)

	for index, word := range words[1:] {
		words[index+1] += strings.Title(word)
	}

	return strings.Join(words, "")
}
