package searchrange

// searchRange Searches for the left and right positions of a target in nums
// For the first_pos, we want to find the leftmost target. So we will do binary search on nums, and whenever we find
// a target, we search for its left side to see whether there is another target on the left while keeping track of the
// leftmost seen first_pos.
// Similarly, we want to find the rightmost target in nums to be last_pos. So when we see a target, we search for its
// right side and record the last_pos = current index. If the current element is not target, then ordinary binary
// search will lead us to the correct searching side.
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	firstPos, lastPos := -1, -1

	// find first position
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			firstPos = mid
			right = mid - 1
		} else if nums[mid] < target {
			// move the left pointer to the middle + 1, because the search range is on the right partition
			left = mid + 1
		} else {
			// move the right pointer to the middle - 1, because the search range is on the left partition
			right = mid - 1
		}
	}

	// reset the pointers
	left, right = 0, len(nums)-1

	// find last position
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			lastPos = mid
			left = mid + 1
		} else if nums[mid] < target {
			// move the left pointer to the middle + 1, because the search range is on the right partition
			left = mid + 1
		} else {
			// move the right pointer to the middle - 1, because the search range is on the left partition
			right = mid - 1
		}
	}

	return []int{firstPos, lastPos}
}
