package validpalindrome

import (
	"gopherland/pkg/utils"
	"regexp"
	"strings"
)

// isPalindrome checks if a phrase/strings is a valid palindrome. Returns true if it is, false otherwise
// after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	regexExpression := `^[A-Za-z0-9]`
	regex := regexp.MustCompile(regexExpression)
	word := ""
	for _, char := range s {
		isAlphanumeric := regex.MatchString(string(char))
		if isAlphanumeric {
			word += string(char)
		}
	}

	return strings.ToLower(word) == strings.ToLower(reverse(word))
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

// isPalindromeTwoPointers uses 2 pointers to find out if a string is a palindrome. This only considers alphanumeric characters & ignores casing.
// Complexity Analysis:
// Time: O(n) where n is the length of the string
// Space: O(1) no extra space is used in determining if this string is a palindrome
func isPalindromeTwoPointers(s string) bool {
	// initialize the left & the right pointers
	left, right := 0, len(s)-1

	// as long as the left pointer is less than the right pointer
	for left < right {

		// and as long as the left pointer is less than the right pointer AND the character is not alphanumeric
		// increase the left pointer by 1 to move it to the next character that is alphanumeric
		for left < right && !utils.IsAlphanumeric(s[left]) {
			left++
		}

		// and as long as the right pointer is greater than the left pointer AND the character at the right pointer is not
		// alphanumeric, decrease the right pointer until we get to a valid alphanumeric character
		for right > left && !utils.IsAlphanumeric(s[right]) {
			right--
		}

		// get the lower case representation of the alphanumeric characters where the pointers now are at
		leftLower := strings.ToLower(string(s[left]))
		rightLower := strings.ToLower(string(s[right]))

		// if they are not equal, then return false
		if leftLower != rightLower {
			return false
		}

		// increase the left & decrease the right pointer to move to the next characters
		left++
		right--
	}
	return true
}
