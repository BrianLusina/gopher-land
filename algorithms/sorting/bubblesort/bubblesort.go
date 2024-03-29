package bubblesort

import "gopherland/pkg/types"

// BubbleSort sorts a slice of numbers in ascending order using bubble sort algorithm
// Approach
//
//	select the first element of the array
//	compare it with its next element
//	if it is larger than the next element then swap them
//	else do nothing
//	keep doing this for every index of the array
//	repeat the above process n times.
//
// Time Complexity
//
//	O(n^2) Worst case performance
//	O(n) Best-case performance
//	O(n^2) Average performance
//	Space Complexity
//	O(1) Worst case as no extra memory is used & the original slice is sorted.
func BubbleSort[T types.Number](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	unsortedUntilIdx := len(arr) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsortedUntilIdx; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		unsortedUntilIdx--
	}
	return arr
}
