package charcodecalc

import (
	"strconv"
	"strings"
)

func calculateTotalFromDigits(digits string) int {
	sum := 0
	for _, digit := range digits {
		if i, err := strconv.Atoi(string(digit)); err == nil {
			sum += i
		}
	}

	return sum
}

func toAsciiCodes(s string) string {
	asciiCodes := []byte(s)

	ascii := ""
	for _, code := range asciiCodes {
		ascii += strconv.Itoa(int(code))
	}

	return ascii
}

// Calc calculates the difference between 2 ascii character codes
func Calc(s string) int {
	asciiCodesOne := toAsciiCodes(s)
	asciiCodesTwo := strings.ReplaceAll(asciiCodesOne, "7", "1")

	sumOne := calculateTotalFromDigits(asciiCodesOne)
	sumTwo := calculateTotalFromDigits(asciiCodesTwo)

	return sumOne - sumTwo
}
