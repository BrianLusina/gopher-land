package removeduplicates

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slowPointer := 0
	fastPointer := 1

	for fastPointer < len(nums) {
		if nums[slowPointer] != nums[fastPointer] {
			slowPointer += 1
			nums[slowPointer] = nums[fastPointer]
		}
		fastPointer++
	}

	return slowPointer + 1
}
