package finddiffoftwoarrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	result1 := []int{}
	result2 := []int{}
	onlyInNums1, onlyInNums2 := getElementsInFirstList(nums1, nums2), getElementsInFirstList(nums2, nums1)

	for n1 := range onlyInNums1 {
		result1 = append(result1, n1)
	}

	for n2 := range onlyInNums2 {
		result2 = append(result2, n2)
	}

	return [][]int{result1, result2}
}

func getElementsInFirstList(firstList []int, secondList []int) map[int]bool {
	firstSet := map[int]bool{}
	secondSet := map[int]bool{}

	for _, n2 := range secondList {
		secondSet[n2] = true
	}

	for _, n1 := range firstList {
		if !secondSet[n1] {
			firstSet[n1] = true
		}
	}

	return firstSet
}
