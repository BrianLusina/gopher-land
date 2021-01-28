package isbn

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

// IsValidISBN checks if an ISBN is a valid ISBN-10
func IsValidISBN(isbn string) bool {
	isbn = removeHyphen(isbn)

	characterArray, err := toSlice(isbn)

	if len(characterArray) != 10 || err != nil {
		return false
	}

	return calculate(characterArray)
}

func removeHyphen(isbn string) (result string) {
	for _, char := range isbn {
		if char == '-' {
			continue
		}
		result += string(char)
	}
	return
}

func toSlice(isbn string) (result []int, err error) {
	for position, character := range isbn {
		switch {
		case unicode.IsLetter(character) && (character != 'X' || position != 9):
			err = errors.New("Invalid character")
			return
		case character == 'X':
			result = append(result, 10)
		default:
			i, _ := strconv.Atoi(string(character))
			result = append(result, i)
		}
	}
	return
}

func calculate(isbn []int) bool {
	var pool int
	for idx, value := range isbn {
		pool += int(math.Abs(float64(idx)-10)) * value
	}
	result := pool % 11

	return result == 0
}
