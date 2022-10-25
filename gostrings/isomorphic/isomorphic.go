/**
Intuition

This approach is based on the idea that the two given strings, if isomorphic, will in some way be exactly the same. If we have two isomorphic strings, we can replace the characters in the first string with the corresponding mapped characters to get the second string. The idea we explore here is the following:

Is there any string transformation we can apply to both the strings such that to check for isomorphism, we simply check if their modified versions are exactly the same?

One can come up with various such transformations giving us different variations of this solution. We will stick with one such transformations for the official solution.

For each character in the given string, we replace it with the index of that character's first occurence in the string.

For a string like paper, the transformed string will be 01034. The character p occurs first at the index 0; so we replace future occurrences of p with the index 0. Similar modifications are made for the other characters. Now let's look at title. The transformed string would be 01034 which is the same as that for paper. This confirms the isomorphic nature of both the strings.

Algorithm

Define a function called transform that takes a string as an input and returns a new string with modifications as explained in the intuition section.
We maintain a dictionary to store the character to index mapping for the given string.
For each character, we look up the mapping in the dictionary. If there is a mapping, that means this character already has its first occurrence recorded and we simply use the first occurrence's index in the new string. Otherwise, we use the current index for the first occurrence.
We find the transformed strings for both of our input strings. Let's say the transformed strings are s1 and s2 respectively.
If s1 == s2, that implies the two input strings are isomorphic. Otherwise, they're not.

Complexity Analysis

Here N is the length of each string (if the strings are not the same length, they cannot be isomorphic).

Time Complexity: O(N). We process each character in both the strings exactly once to determine if they are isomorphic.
Space Complexity: O(N). We form two new strings returned by our transformation function. The size of ASCII character set is fixed and the keys in our dictionary are valid ASCII characters only. So the size of the map in the transform function doesn't contribute to the space complexity.
*/
package isomorphic

import "fmt"

// transformString transforms a string and returns a new string where the first occurence of the string index
// is used as a mapping for all other occurences of the string
func transformString(s string) (newString string) {
	mapping := map[rune]int{}

	for i, v := range s {
		_, ok := mapping[v]
		if !ok {
			mapping[v] = i
		}
		newString += fmt.Sprintf("%d", mapping[v])
	}

	return
}

// IsIsomorphic returns true if 2 strings are isomorphic
func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return transformString(s) == transformString(t)
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[byte]byte)
	mapT := make(map[byte]byte)

	for i := range s {

		vS, vT := s[i], t[i]

		if m, ok := mapS[vS]; ok {
			if m != vT {
				return false
			}
		} else {
			mapS[vS] = vT
		}

		if m, ok := mapT[vT]; ok {
			if m != vS {
				return false
			}
		} else {
			mapT[vT] = vS
		}

	}

	return true
}
