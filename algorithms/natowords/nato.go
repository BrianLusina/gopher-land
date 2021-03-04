package natowords

import (
	"strings"
	"unicode"
)

var nato = map[rune]string{
	'a': "Alfa ",
	'b': "Bravo ",
	'c': "Charlie ",
	'd': "Delta ",
	'e': "Echo ",
	'f': "Foxtrot ",
	'g': "Golf ",
	'h': "Hotel ",
	'i': "India ",
	'j': "Juliett ",
	'k': "Kilo ",
	'l': "Lima ",
	'm': "Mike ",
	'n': "November ",
	'o': "Oscar ",
	'p': "Papa ",
	'q': "Quebec ",
	'r': "Romeo ",
	's': "Sierra ",
	't': "Tango ",
	'u': "Uniform ",
	'v': "Victor ",
	'w': "Whiskey ",
	'x': "X-ray ",
	'y': "Yankee ",
	'z': "Zulu ",
}

// ToNato converts a string into Nato code
func ToNato(words string) (result string) {
	wordsLower := strings.ToLower(words)
	for _, letter := range wordsLower {

		natoWord := nato[letter]

		result += string(natoWord)

		if natoWord == "" && letter != ' ' {
			result += string(letter)
		}
	}

	return strings.Trim(result, " ")
}

func toNato(words string) (result string) {
	natoList := []string{
		"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
		"Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
		"Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
		"Whiskey", "Xray", "Yankee", "Zulu",
	}

	toCharlieMap := map[rune]string{}

	for _, value := range natoList {
		toCharlieMap[rune(value[0])] = value
	}

	for _, letter := range words {
		if unicode.IsLetter(letter) {
			result += toCharlieMap[unicode.ToUpper(letter)] + " "
		} else if unicode.IsPunct(letter) {
			result += string(letter)
		}
	}
	return strings.TrimSpace(result)
}
