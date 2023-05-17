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

// GenerateAnagrams returns a slice of all anagrams of a given string
func GenerateAnagrams(word string) []string {
	// if the input string is 0, return an empty slice
	if len(word) == 0 {
		return []string{}
	}

	// if the input string only has 1 character, return it in a slice
	if len(word) == 1 {
		return []string{word}
	}

	// collection to hold all our anagrams, This results in a space complexity of O(n!), factorial complexity, where n is the
	// length of ths string. This is because a string of n length could potentially have n! possible anagrams
	anagrams := []string{}

	// find all the anagrams of the substring from the second character until the end.
	// for example, if the string is abcd, the substring is bcd, so, we find the anagrams
	// bcd

	substringAnagrams := GenerateAnagrams(word[1:])

	// iterate over each substring
	for _, substringAnagram := range substringAnagrams {

		// iterate over each index of the substring, from 0 until 1 index past the end of the string
		for idx := range substringAnagram {
			// create a copy of the substring anagram
			substringAnagramCopy := strings.Clone(substringAnagram)

			// insert the first character of ours string into the substring anagram copy. Where it will go
			// depends on the index we're up to within this loop. Then, take this new string and add it to
			// our collection of anagrams
			char := word[0]
			a := strings.Replace(substringAnagramCopy, string(substringAnagramCopy[idx]), string(char), idx)
			anagrams = append(anagrams, a)
		}
	}

	return anagrams
}
