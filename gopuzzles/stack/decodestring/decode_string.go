package decodestring

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	countStack := []int{}
	stack := []string{}
	result := ""
	counts := ""

	for _, char := range s {
		if unicode.IsDigit(char) {
			counts += string(char)
		} else if char == '[' {
			c, err := strconv.Atoi(counts)
			if err != nil {
				panic(fmt.Errorf("failed to convert %s to int", counts))
			}
			countStack = append(countStack, c)
			stack = append(stack, result)
			counts = ""
			result = ""
		} else if char == ']' {
			str := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			result = str + strings.Repeat(result, count)
		} else {
			result += string(char)
		}
	}

	return result
}
