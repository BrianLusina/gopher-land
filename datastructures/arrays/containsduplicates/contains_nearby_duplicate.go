package containsduplicates

// ContainsNearbyDuplicate checks if there are any duplicates in the array
// if there are 2 distinct indices i & j in the array such that nums[i] == nums[j] & abs(i - j) <= k, true is returned
func ContainsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || k < 1 {
		return false
	}

	// create a map to store the values
	m := make(map[int]int)

	// iterate through the array
	for i, num := range nums {
		// check if the value is already in the map
		if _, ok := m[num]; ok {
			// if it is, check if the difference between the current index and the previous index is less than or equal to k
			if i-m[num] <= k {
				return true
			}
		}

		// add the value to the map
		m[num] = i
	}

	return false
}
