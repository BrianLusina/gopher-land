package atbash

import (
	"regexp"
	"strings"
)

const (
	ALPHABET_LOWER = "abcdefghijklmnopqrstuvwxyz"
)

func indexOf(slice []string, s string) int {
	for p, v := range slice {
		if v == s {
			return p
		}
	}
	return -1
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func encode(plain string) string {
	key := reverse(ALPHABET_LOWER)
	inputSlice := strings.Split(plain, "")
	originalSlice := strings.Split(ALPHABET_LOWER, "")
	reversedSlice := strings.Split(key, "")
	output := ""
	for i := 0; i < len(plain); i++ {
		char := inputSlice[i]
		index := indexOf(originalSlice, char)
		if index > -1 {
			output += reversedSlice[index]
		} else {
			output += char
		}
	}
	return output
}

func chunk(s string) []string {
	reChunk := regexp.MustCompile(".{1,5}")
	return reChunk.FindAllString(s, -1)
}

func normalize(s string) string {
	reNormalize := regexp.MustCompile("[^a-z0-9]")
	return reNormalize.ReplaceAllString(s, "")
}

func Atbash(s string) string {
	plain := strings.ToLower(s)
	normalized := normalize(plain)
	cipher := encode(normalized)

	chunks := chunk(cipher)
	return strings.Join(chunks, " ")
}
