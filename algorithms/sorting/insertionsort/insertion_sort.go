package insertionsort

import "gopherland/pkg/types"

// InsertionSort uses the insertion sort algorithm on a collection of comparable item types
func InsertionSort[T types.Comparable](collection []T) []T {
	for index := 1; index < len(collection); index++ {
		temp := collection[index]
		previousIdx := index - 1

		for previousIdx >= 0 {
			if collection[previousIdx] > temp {
				collection[previousIdx+1] = collection[previousIdx]
				previousIdx--
			} else {
				break
			}
		}
		collection[previousIdx+1] = temp
	}

	return collection
}
