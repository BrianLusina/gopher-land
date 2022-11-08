package productexceptself

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	output := make([]int, len(nums))

	prefix := 1

	for idx := range nums {
		output[idx] = prefix
		prefix *= nums[idx]
	}

	postfix := 1

	for i := len(nums) - 1; i > -1; i-- {
		output[i] *= postfix
		postfix *= nums[i]
	}

	return output
}
