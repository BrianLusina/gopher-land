package sqrtestimate

func sqrtEstimate(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := x
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
