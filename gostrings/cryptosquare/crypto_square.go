package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// norm is a function that normalizes a rune, replaces all runes that are not alphanumeric
// and all runes that are uppercase with lowercase
func norm(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
		return r
	case r >= 'A' && r <= 'Z':
		return r + 'a' - 'A'
	}
	return -1
}

// normalize a string, replacing all non-alphanumeric runes with spaces
func normalize(s string) string {
	re := regexp.MustCompile(`[^\w]+`)
	return re.ReplaceAllString(strings.ToLower(s), "")
}

func Encode(pt string) string {
	// either normalize or nor pass the tests and implement the solution correctly,
	// however norm uses less memory than normalize and allocates less bytes per operaation
	plain := strings.Map(norm, pt)
	// plain := normalize(pt)
	numCols := int(math.Ceil(math.Sqrt(float64(len(plain)))))
	padding := numCols*(numCols-1) - len(plain)
	if padding < 0 {
		padding = numCols*numCols - len(plain)
	}

	cols := make([]string, numCols)
	for i, r := range plain {
		cols[i%numCols] += string(r)
	}
	for i := 0; i < padding; i++ {
		cols[numCols-i-1] += " "
	}
	return strings.Join(cols, " ")
}
