package indexedcapitalization

import (
	// "strings"
	"unicode"
)

func Capitalize(st string, arr []int) string {
	r := []rune(st)
	for _, idx := range arr {
		if idx < len(st) {
			r[idx] = unicode.ToUpper(r[idx])
			// another implementation of this, this modifies the string
			// st = st[:idx] + strings.ToUpper(string(st[idx])) + st[idx+1:]
		}
	}
	return string(r)
	// return st
}
