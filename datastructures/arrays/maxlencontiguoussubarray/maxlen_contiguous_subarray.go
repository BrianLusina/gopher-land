package maxlencontiguoussubarray

import "math"

// findMaxLength finds the length of the maximum length contiguous subarray of a binary array
// Makes use of a hashmap to store the entries in the form of (count, index). We make an entry for a
// count in the map whenever the count is encountered first and later on use the corresponding index
// to find the length of the largest sub array with equal number of zeros and ones
// when the same count is encountered again.

// Complexity Analysis:

// - Time complexity : O(n). The entire array is traversed only once.
// - Space complexity : O(n). Maximum size of the HashMap  will be n, if all the elements are either 1 or 0.
func findMaxLength(nums []int) int {
	hashmap := make(map[int]int)

	// initializing map to (0, -1) is to avoid the edge cases like [0, 1] when you only have one zero and one.
	hashmap[0] = -1
	max := 0
	count := 0

	for idx, num := range nums {
		if num == 1 {
			count = count + 1
		} else {
			count = count - 1
		}

		if _, ok := hashmap[count]; ok {
			max = int(math.Max(float64(max), float64(idx-hashmap[count])))
		} else {
			hashmap[count] = idx
		}
	}
	return max
}
