package intersectionofarrays

func toSet(nums []int) map[int]bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if _, ok := set[num]; !ok {
			set[num] = true
		}
	}
	return set
}

func setIntersection(setOne, setTwo map[int]bool) []int {
	intersection := []int{}
	for key := range setOne {
		if _, ok := setTwo[key]; ok {
			intersection = append(intersection, key)
		}
	}
	return intersection
}

func intersection(nums1 []int, nums2 []int) []int {
	setOne := toSet(nums1)
	setTwo := toSet(nums2)

	if len(setOne) < len(setTwo) {
		return setIntersection(setOne, setTwo)
	} else {
		return setIntersection(setTwo, setOne)
	}
}
