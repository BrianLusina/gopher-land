package gcdofstrings

import (
	"gopherland/math/utils"
	"strings"
)

func gcdOfStringsBruteForce(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)

	n := utils.Min(len1, len2)

	valid := func(k int) bool {
		if utils.Modulo(len1, k) > 0 || utils.Modulo(len2, k) > 0 {
			return false
		}

		n1, n2 := len1/k, len2/k
		base := str1[:k]

		return str1 == strings.Repeat(base, n1) && str2 == strings.Repeat(base, n2)
	}

	for i := n; i > 0; i-- {
		if valid(i) {
			return str1[:i]
		}
	}

	return ""
}
