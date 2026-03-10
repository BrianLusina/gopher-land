package asteroidcollision

import (
	"gopherland/math/utils"
)

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	// time complexity incurred here is O(N) as we iterate over all the asteroids
	for _, asteroid := range asteroids {
		flag := true

		for len(stack) > 0 && (0 < stack[len(stack)-1] && asteroid < 0) {
			// If the top asteroid in the stack is smaller, then it will explode.
			// Hence pop it from the stack, also continue with the next asteroid in the stack.
			if utils.Abs(stack[len(stack)-1]) < utils.Abs(asteroid) {
				_ = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				continue
			} else if utils.Abs(stack[len(stack)-1]) == utils.Abs(asteroid) {
				_ = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

			}

			// If we reach here, the current asteroid will be destroyed
			// Hence, we should not add it to the stack
			flag = false
			break
		}

		if flag {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
