package highestrank

// HighestRank returns the highest rank number in the array. If there are multiple highest rank numbers, return the largest of them.
func HighestRank(nums []int) int {
	counter, maxKey, maxValue := make(map[int]int), 0, 0

	for _, num := range nums {
		counter[num]++
		if counter[num] > maxValue || (counter[num] == maxValue && num > maxKey) {
			maxValue, maxKey = counter[num], num
		}
	}

	return maxKey
}
