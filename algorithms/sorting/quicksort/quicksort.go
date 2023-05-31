package quicksort

import "gopherland/pkg/types"

// partitions a slice into 2 and returns the pivot index. Assumes the pivot is at the end of the slice
func partition[T types.Number](collection []T, startIndex, endIndex int) int {

	pivot := collection[endIndex]

	leftPointer := startIndex
	rightPointer := endIndex - 1

	for leftPointer <= rightPointer {
		// walk until we find something on the left side that belongs on the right (less than the pivot)
		for leftPointer <= endIndex && collection[leftPointer] < pivot {
			leftPointer++
		}

		// walk until we find something on the right side that belongs on the left(greater than or equal to the pivot)
		for rightPointer >= startIndex && collection[rightPointer] >= pivot {
			rightPointer--
		}

		if leftPointer < rightPointer {
			// swap the items at the left_pinter and right_pointer, moving the element that's smaller than the pivot to the left
			// half and the element that's larger than the pivot to the right half
			collection[rightPointer], collection[leftPointer] = collection[leftPointer], collection[rightPointer]
		} else {
			// unless we have looked at all the elements in the list and are done partitioning. In that case, move the pivot element
			// into it's final position
			collection[endIndex], collection[leftPointer] = collection[leftPointer], collection[endIndex]
		}
	}

	return leftPointer
}

// quicksortSubSlice uses recursion to sort each partition of the slice
func quicksortSubSlice[T types.Number](collection []T, startIndex, endIndex int) {

	// base case, list with 0 or 1 element
	if startIndex >= endIndex {
		return
	}

	// divide the list into 2 smaller sublists
	pivotIndex := partition(collection, startIndex, endIndex)

	// Recursively sort each sublist
	quicksortSubSlice(collection, startIndex, pivotIndex-1)
	quicksortSubSlice(collection, pivotIndex+1, endIndex)
}

// Quicksort sorts a slice of integers using quicksort algorithm
func Quicksort[T types.Number](collection []T) []T {
	length := len(collection)

	// Nothing to sort here
	if length <= 1 {
		return collection
	}

	quicksortSubSlice(collection, 0, length-1)
	return collection
}
