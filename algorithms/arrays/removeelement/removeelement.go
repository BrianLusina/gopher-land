package removeelement

// RemoveElement removes all instances of an element from a slice in place
func RemoveElement(nums []int, val int) int {
	pointer := 0
	size := len(nums)

	for pointer < size {
		if nums[pointer] == val {
			nums[pointer] = nums[size-1]
			size--
		} else {
			pointer++
		}
	}

	return size
}
