package kokoeatingbananas

import "math"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1000000000
	result := -1

	for left <= right {
		mid := (left + right) / 2

		if canFinishEating(piles, h, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func canFinishEating(piles []int, h, k int) bool {
	hoursUsed := 0

	for _, pile := range piles {
		// Since Koko eats at only one pile during each hour, ceil(float(p)/k) is the time Koko takes to finish 1 pile
		// Note that p/k does not work here because we want a whole number of hours so we needed to round up p/k.
		// Therefore, the feasiblity is determined by whether Koko's hours_used is within h hours, where hours_used is the total hours to finish all piles.
		hoursUsed += int(math.Ceil(float64(pile) / float64(k)))
	}

	return hoursUsed <= h
}
