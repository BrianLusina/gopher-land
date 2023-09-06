// Package shellsort contains algorithms that uses the shell sort algorithm to perform sorting on comparable collection of items
package shellsort

import "gopherland/pkg/types"

// ShellSort sorts a collection of comparable items in place
func ShellSort[T types.Comparable](collection []T) []T {
	if len(collection) <= 1 {
		return collection
	}

	distance := len(collection)

	for distance > 0 {
		for i := distance; i < len(collection); i++ {
			temp := collection[i]
			j := i

			// sort the sub list for this distance
			for j >= distance && collection[j-distance] > temp {
				collection[j] = collection[j-distance]
				j = j - distance
			}
			collection[j] = temp
		}
		// reduce the distance for the next element
		distance = distance / 2
	}
	return collection
}
