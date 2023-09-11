package mininrotatedsortedarray

// findMin returns the minimum value in a rotated sorted array
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		if left == right {
			return nums[left]
		} else if left+1 == right {
			return min(nums[left], nums[right])
		}

		middle := left + (right-left)/2

		if nums[middle] < nums[right] {
			right = middle
		} else if nums[left] < nums[middle] {
			left = middle
		}
	}

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
