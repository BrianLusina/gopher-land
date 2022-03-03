package fixstringcase

import (
	"strings"
	"unicode"
)

func fixstringcase(str string) string {
	uppercaseCount := 0
	lowercaseCount := 0

	for _, char := range str {
		if unicode.IsUpper(char) {
			uppercaseCount++
		} else if unicode.IsLower(char) {
			lowercaseCount++
		}
	}
	if uppercaseCount > lowercaseCount {
		return strings.ToUpper(str)
	} else if lowercaseCount > uppercaseCount || lowercaseCount == uppercaseCount {
		return strings.ToLower(str)
	} else {
		return strings.ToLower(str)
	}
}
