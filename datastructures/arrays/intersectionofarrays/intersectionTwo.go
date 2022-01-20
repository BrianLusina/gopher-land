package intersectionofarrays

func intersectTwo(nums1 []int, nums2 []int) []int {
	counts := make(map[int]int)
	result := []int{}

	for _, num := range nums1 {
		if _, ok := counts[num]; !ok {
			counts[num] = 1
		} else {
			counts[num] += 1
		}
	}

	for _, num := range nums2 {
		if _, ok := counts[num]; ok {
			if counts[num] > 0 {
				result = append(result, num)
				counts[num] -= 1
			}
		}
	}

	return result
}
