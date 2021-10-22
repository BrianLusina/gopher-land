package luhn

import "strings"

func sum(digits []int) int {
	sum := 0
	for idx, digit := range digits {
		j := len(digits) - idx
		if j%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum
}

func Valid(input string) bool {
	if len(strings.TrimSpace(input)) <= 1 {
		return false
	}

	digits := make([]int, 0, len(input))

	numbers := strings.ReplaceAll(input, " ", "")

	for _, number := range numbers {
		if number < '0' || number > '9' {
			return false
		}

		digit := int(number - '0')
		digits = append(digits, digit)
	}

	return sum(digits)%10 == 0
}
