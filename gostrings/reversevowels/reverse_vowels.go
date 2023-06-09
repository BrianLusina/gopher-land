package reversevowels

import "strings"

// reverseVowels reverses all the vowels in the string
func reverseVowels(word string) string {
	leftPointer := 0
	rightPointer := len(word) - 1

	letters := strings.Split(word, "")

	for leftPointer < rightPointer {
		for leftPointer < rightPointer && !isVowel(word[leftPointer]) {
			leftPointer++
		}

		for leftPointer < rightPointer && !isVowel(word[rightPointer]) {
			rightPointer--
		}

		letters[leftPointer], letters[rightPointer] = letters[rightPointer], letters[leftPointer]

		leftPointer++
		rightPointer--
	}

	return strings.Join(letters, "")
}

func isVowel(character byte) bool {
	return character == 'a' || character == 'i' || character == 'e' || character == 'o' || character == 'u' || character == 'A' || character == 'I' || character == 'E' || character == 'O' || character == 'U'
}
