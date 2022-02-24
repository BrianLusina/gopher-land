package containsduplicates

import "math"

// ContainsNearbyDuplicate checks if there are any duplicates in the array
// Given an integer array nums and two integers k and t, return true if there
// are two distinct indices i and j in the array such that
// abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.
func ContainsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	d := make(map[int]int)
	w := t + 1

	for idx, num := range nums {
		m := num / w

		if _, ok := d[m]; ok {
			return true
		}

		if v, ok := d[m-1]; ok {
			if int(math.Abs(float64(num)-float64(v))) < w {
				return true
			}
		}

		if v, ok := d[m+1]; ok {
			if int(math.Abs(float64(num)-float64(v))) < w {
				return true
			}
		}

		d[m] = num
		if idx >= k {
			delete(d, nums[idx-k]/w)
		}
	}
	return false
}
