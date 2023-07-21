package removingstars

import "strings"

func removeStars(s string) string {
	stack := []string{}

	for _, char := range s {
		if char == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(char))
		}
	}

	return strings.Join(stack, "")
}

func removeStars2(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
