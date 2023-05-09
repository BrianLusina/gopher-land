package selectionsort

import "gopherland/pkg/types"

// SelectionSort uses selection sort algorithm to sort a collection of comparable types
// This works on strings, floats and integers and has a runtime complexity of O(N^2) and space complexity of O(1) as we modify the collection in place
func SelectionSort[T types.Comparable](collection []T) []T {
	for i := 0; i < len(collection); i++ {
		// lowestIdx is the index of the lowest value seen so far,at the start this is the first element
		lowestIdx := i

		// we iterate over the rest of the elements starting from the next element from the first one(assuming this is the first iteration)
		for y := i + 1; y < len(collection); y++ {
			// we check if the next element is less than the previous element
			if collection[y] < collection[i] {
				// if it is, we keep track of the new lowest index seen so far
				lowestIdx = y
			}
		}

		// if the lowest index seen so far is not equal to the current index, then we have found a lower value element
		if lowestIdx != i {
			// we swap the 2 elements
			collection[i], collection[lowestIdx] = collection[lowestIdx], collection[i]
		}
	}

	// return the modified collection
	return collection
}
