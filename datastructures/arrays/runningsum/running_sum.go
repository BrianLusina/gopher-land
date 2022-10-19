package runningsum

func runningSum(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		next := sum(nums[0 : i+1])
		result[i] = next
	}

	return result

	// OR
	// for i:=1; i<len(nums); i++ {
	//     nums[i] += nums[i-1]
	// }
	// return nums
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
