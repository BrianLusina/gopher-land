package highestaltitude

import "gopherland/math/utils"

func largestAltitude(gain []int) int {
	currentAltitude, highest := 0, 0
	for altitude := range gain {
		currentAltitude += gain[altitude]
		highest = utils.Max(highest, currentAltitude)
	}

	return highest
}
