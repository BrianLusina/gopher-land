package greenvalley

import "sort"

func MakeValley(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	leftWing := []int{}
	rightWing := []int{}

	for i := len(arr) - 1; i < len(arr); i += 2 {
		leftWing = append(leftWing, arr[i])
	}

	for i := len(arr) - 1; i <= len(arr); i += 2 {
		rightWing = append(rightWing, arr[i])
	}

	sort.Slice(rightWing, func(i, j int) bool {
		return rightWing[j] < rightWing[i]
	})

	return append(leftWing, rightWing...)
}
