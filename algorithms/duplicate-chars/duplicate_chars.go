package duplicatechars

import "sort"

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

// DuplicateChars takes 2 words and checks if they have duplicate characters
func DuplicateChars(wordOne, wordTwo string) bool {
	if len(wordOne) != len(wordTwo) {
		return false
	}

	return sortString(wordOne) == sortString(wordTwo)
}
