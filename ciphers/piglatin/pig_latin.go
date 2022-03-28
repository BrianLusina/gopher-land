package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	constPattern      = `^(?P<head>[^aeiou]?qu|[^aeiou]+)(?P<tail>[a-z]*)`
	constPatternWithy = `^(?P<head>[^aeiou]+)y(?P<tail>[a-z]*)`
	vowelPattern      = `^([aeiou]|y[^aeiou]|xr)[a-z]*`
)

var (
	constRegex     = regexp.MustCompile(constPattern)
	vowelRegex     = regexp.MustCompile(vowelPattern)
	containsyRegex = regexp.MustCompile(constPatternWithy)
)

func splitInitialConsonantWord(word string) (string, string) {
	containsY := containsyRegex.MatchString(word)
	match := constRegex.MatchString(word)
	if containsY {
		matches := containsyRegex.FindStringSubmatch(word)
		headIndex := containsyRegex.SubexpIndex("head")
		tailIndex := containsyRegex.SubexpIndex("tail")
		return matches[headIndex], fmt.Sprintf("y%s", matches[tailIndex])
	} else if match {
		matches := constRegex.FindStringSubmatch(word)
		headIndex := constRegex.SubexpIndex("head")
		tailIndex := constRegex.SubexpIndex("tail")
		return matches[headIndex], matches[tailIndex]
	}
	return "", ""
}

func startsWithVowelSound(word string) bool {
	return vowelRegex.Match([]byte(word))
}

func Sentence(sentence string) string {
	pigLatin := []string{}
	words := strings.Fields(sentence)

	for _, word := range words {
		if startsWithVowelSound(word) {
			pigLatin = append(pigLatin, fmt.Sprintf("%say", word))
		} else {
			head, tail := splitInitialConsonantWord(word)
			pigLatin = append(pigLatin, fmt.Sprintf("%s%say", tail, head))
		}
	}
	return strings.Join(pigLatin, " ")
}
