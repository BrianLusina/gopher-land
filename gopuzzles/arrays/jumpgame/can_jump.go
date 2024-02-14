package jumpgame

import "gopherland/math/utils"

func canJump(nums []int) bool {
	currentPosition := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentPosition == 0 {
			return false
		}

		currentPosition -= 1

		currentPosition = utils.Max(currentPosition, nums[i])
	}

	return true
}
