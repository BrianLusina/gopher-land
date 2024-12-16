package spreadsheetencoding

import "math"

/*
spreadsheetEncodeColumn encodes a Spreadsheet column ID as an integer and returns it.
Here the encoding uses the alphabets A, B, C, etc. and
further uses the indexing of the alphabet starting from 1 like: A=1, B=2, C=3,..., Z=26
Complexity
Where n represents the number of characters in the column_id

- Time O(n) as we iterate over each character in the column ID to calculate the encoding
- Space O(1) no extra space is required to perform the encoding
*/
func spreadsheetEncodeColumn(columnID string) int {
	num := 0
	count := len(columnID) - 1

	for _, char := range columnID {
		// we use the base 26 system here as there are 26 letters in the alphabet
		// char is a rune which is the unicode code point for a character in a string.
		// In order to make sure that A=1, we need to determine the relative difference from this rune and 'A' & from the representation
		// we require for base 26 system
		// Now we know that the rune 'A' equals 65. So if we find the Unicode code point of a character,
		// subtract 'A' from it, and then add 1 to it, we’ll get the representation we want for the base 26 system.

		base := int(math.Pow(26.0, float64(count)))
		diff := int(char - 'A' + 1)

		num += base * diff
		count--
	}
	return num
}
