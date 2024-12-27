package permutation

import (
	"sort"
	"strings"
)

// CheckPermutationWithSorting checks if two strings are permutations of each other
// It uses sorting to achieve this
//
// Complexity:
// Time: O(nlog(n)) as sorting is used and sorting here is an O(nlog(n)) time complexity
// Space: O(1) as no extra space is used to achieve checking for permutations
//
// If the length of the strings are different, it returns false
// Otherwise, it returns true
func CheckPermutationWithSorting(str1, str2 string) bool {
	// If the length of the strings are different, they can't be permutations
	if len(str1) != len(str2) {
		return false
	}

	s1 := strings.Split(strings.ToLower(str1), "")
	s2 := strings.Split(strings.ToLower(str2), "")

	sort.Strings(s1)
	sort.Strings(s2)

	s1Sorted := strings.Join(s1, "")
	str2Sorted := strings.Join(s2, "")

	n := len(s1Sorted)

	for i := 0; i < n; i++ {
		if s1Sorted[i] != str2Sorted[i] {
			return false
		}
	}

	return true
}

// CheckPermutationWithMap checks if two strings are permutations of each other
// It uses a map to achieve this
//
// Complexity:
// Time: O(n) as it iterates through the strings once
// Space: O(n) as it uses a map to store the frequency of each character in the strings
//
// If the length of the strings are different, it returns false
// Otherwise, it returns true

func CheckPermutationWithMap(str1, str2 string) bool {
	// If the length of the strings are different, they can't be permutations
	if len(str1) != len(str2) {
		return false
	}

	s1 := strings.Split(strings.ToLower(str1), "")
	s2 := strings.Split(strings.ToLower(str2), "")

	charCount := make(map[string]int)

	for _, s := range s1 {
		charCount[s]++
	}

	for _, s := range s2 {
		_, ok := charCount[s]
		if ok {
			charCount[s]--
		} else {
			return false
		}
	}

	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}

	return true
}
