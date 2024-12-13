package lookandsaysequence

import (
	"strconv"
	"strings"
)

// lookAndSaySequence
func lookAndSaySequence(sequence string) string {
	result := []string{}
	i := 0

	for i < len(sequence) {
		count := 0
		for i+1 < len(sequence) && sequence[i] == sequence[i+1] {
			i++
			count++
		}

		c := strconv.Itoa(count)
		s := sequence[i]

		char := strconv.Itoa(int(s))

		result = append(result, c+char)

		i++
	}

	return strings.Join(result, "")
}
