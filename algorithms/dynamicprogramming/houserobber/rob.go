package houserobber

import "gopherland/pkg/utils"

func rob(nums []int) int {
	current, previous := 0, 0

	for _, house := range nums {
		current, previous = utils.Max(previous+house, current), current
	}

	return current
}
