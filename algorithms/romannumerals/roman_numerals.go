package romannumerals

import (
	"errors"
	"strings"
)

var arabicNumbers = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}
var romanNumbers = []string{
	"M",
	"CM",
	"D",
	"CD",
	"C",
	"XC",
	"L",
	"XL",
	"X",
	"IX",
	"V",
	"IV",
	"I",
}

// ToRomanNumeral converts an arabic number to a roman numeral
func ToRomanNumeral(arabic int) (romanNumber string, err error) {
	if arabic <= 0 || arabic > 3000 {
		err = errors.New("cannot convert given Arabic number to Roman Numeral")
		return "", err
	}

	for index := range arabicNumbers {
		count := arabic / arabicNumbers[index]
		romanNumber += strings.Repeat(romanNumbers[index], count)
		arabic -= arabicNumbers[index] * count
	}
	return romanNumber, nil
}
