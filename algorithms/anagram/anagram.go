package anagram

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func isAnagram(subject, candidate string) bool {
	return subject != candidate && sortString(candidate) == sortString(subject)
}

// Detect checks if the candidates list contains words that can be formed by re-arranging the subject
func Detect(subject string, candidates []string) (sublist []string) {

	if len(candidates) == 0 {
		return sublist
	}

	subject = strings.ToLower(subject)

	for _, candidate := range candidates {

		if tc := strings.ToLower(candidate); isAnagram(subject, tc) {
			sublist = append(sublist, candidate)
		}
	}

	return
}

// GroupAnagrams groups a slice of strings into a group of anagrams
func GroupAnagrams(strs []string) [][]string {
	length := len(strs)
	if length == 0 {
		return [][]string{}
	}

	if length == 1 {
		return [][]string{strs}
	}

	m := map[string][]string{}

	for _, word := range strs {
		sortedWord := sortString(word)

		if _, ok := m[sortedWord]; ok {
			m[sortedWord] = append(m[sortedWord], word)
		} else {
			m[sortedWord] = []string{word}
		}
	}

	words := [][]string{}
	for _, v := range m {
		words = append(words, v)
	}

	return words
}
