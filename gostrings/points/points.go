package points

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	total := 0

	for _, game := range games {
		result := strings.Split(game, ":")
		x, err := strconv.Atoi(result[0])
		if err != nil {
			return total
		}

		y, err := strconv.Atoi(result[1])
		if err != nil {
			return total
		}

		if x > y {
			total += 3
		} else if x == y {
			total += 1
		}

	}

	return total
}
