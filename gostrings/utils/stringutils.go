package utils

import "strings"

// ReverseString reverses a string with 1st character become last and so on
func ReverseString(s string) string {
	var rns = make([]byte, len(s))
	copy(rns, s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters/runes of the string
		// the first with the last, etc
		rns[i], rns[j] = rns[j], rns[i]
	}

	// convert rune slice back to string
	return string(rns)

	// another mode of reversing a string
	// var result string
	// for _, v := range s {
	// 	result = string(v) + result
	// }
	// return result
}

func RemoveNewLineSuffixes(s string) string {
	if s == "" {
		return s
	}

	if strings.HasSuffix(s, "\r\n") {
		return RemoveNewLineSuffixes(s[:len(s)-2])
	}

	if strings.HasSuffix(s, "\n") {
		return RemoveNewLineSuffixes(s[:len(s)-1])
	}
	return s
}
