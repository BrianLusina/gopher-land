package minstrlen

import "strings"

func MinStrLenStack(s string) int {
	if len(s) == 0 {
		return 0
	}

	stack := []rune{}

	for _, char := range s {
		// If the stack is empty
		if len(stack) == 0 {
			// Add the current character to the stack and continue to the next character
			stack = append(stack, char)
			continue
		}

		// If the current character is B and the top of the stack is A
		if char == 'B' && stack[len(stack)-1] == 'A' {
			// Pop off the character from the stack
			stack = stack[:len(stack)-1]
			// If the current character is D and the top of the stack is C
		} else if char == 'D' && stack[len(stack)-1] == 'C' {
			// Pop off the character from the top of the stack
			stack = stack[:len(stack)-1]
		} else {
			// otherwise, push the current character to the stack
			stack = append(stack, char)
		}
	}

	return len(stack)
}

func MinStrLenReplace(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		if strings.Contains(s, "AB") {
			s = strings.Replace(s, "AB", "", -1)
		} else if strings.Contains(s, "CD") {
			s = strings.Replace(s, "CD", "", -1)
		}
	}

	return len(s)
}

func MinStrLenInPlaceManipulation(s string) int {
	runes := []rune(s)
	writeIndex := 0

	for _, char := range runes {
		if writeIndex > 0 && ((char == 'B' && runes[writeIndex-1] == 'A') || (char == 'D' && runes[writeIndex-1] == 'C')) {
			writeIndex--
		} else {
			runes[writeIndex] = char
			writeIndex++
		}
	}

	return writeIndex
}
