package sumofthree

import "sort"

func findSomeOfThree(numbers []int, target int) bool {
	sort.IntSlice(numbers).Sort()

	for i := 0; i < len(numbers)-2; i++ {
		left, right := i+1, len(numbers)-1
		for left < right {
			sum := numbers[i] + numbers[left] + numbers[right]
			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}
