package rotatedsortedarray

// search searches through a rotated sorted slice of ints to find the given target using Binary Search. if the target
// does not exist, -1 is returned
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if target == nums[middle] {
			return middle
		}

		// left sorted portion
		if nums[left] <= nums[middle] {
			// if our target is greater than our middle value or less than the left most value, then we need to move our
			// left pointer to the middle plus 1, this means our target is in the right sorted portion
			if target > nums[middle] || target < nums[left] {
				left = middle + 1
			} else {
				// if not, we move our right pointer to the middle. Meaning our target is in the left sorted portion
				right = middle - 1
			}
			// right sorted portion
		} else {
			// else our target is in the right sorted portion. We check if the target is less than our middle or our
			// target is greater than our right most value. Then we move to the left portion by moving our right pointer
			// to the middle value
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				// Move our left to the middle position if our target is greater than middle value or less than right
				// most value
				left = middle + 1
			}
		}
	}

	return -1
}
