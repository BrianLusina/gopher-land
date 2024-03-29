package longestincreasingsubsequence

func longestIncreasingSubsequence(nums []int) int {
	piles := []int{}

	for _, num := range nums {
		// binary search over the piles
		left, right := 0, len(piles)

		for left < right {
			middle := left + (right-left)/2
			if num > piles[middle] {
				left = middle + 1
			} else {
				right = middle
			}
		}

		// if our left pointer isn't larger than piles array
		// it means we found a place to put our number
		if left < len(piles) {
			// overwrite piles[left] with the current num
			piles[left] = num
		} else {
			// left > length of piles array, it means there is no pile with a large
			// enough number to put this number on top of
			// create a new pile by appending current number to end
			piles = append(piles, num)
		}
	}

	return len(piles)
}
