// Package interpolation contains functions that use interpolation search to find a key in a given collection
package interpolation

import "gopherland/pkg/types"

// Search searches through a collection using interpolation search algorithm. Assumption made is that the collection is uniformly distributed and sorted
// Further T must implement the operators -, !=, ==, >=, <= and < such that >=, <=, !=, == and < define a total order on T and such that
// (tm - tl) * k / (th - tl) is an int between 0 and k (inclusive) for any tl, tm, th in T with tl <= tm <= th, tl != th.
// arr must be sorted according to this ordering.
func Search[T types.Comparable](collection []T, key T) int {
	if len(collection) == 0 {
		// nothing to do here as collection is already empty, exit early
		return -1
	}

	low := 0
	high := len(collection) - 1

	for low <= high && collection[low] <= key && key <= collection[high] {
		midpoint := low + ((key - collection[low]) * (high - low) / (collection[high] - collection[low]))

		if collection[midpoint] == key {
			return midpoint
		} else if key > collection[midpoint] {
			low = midpoint + 1
		} else {
			high = midpoint - 1
		}
	}

	if key == collection[low] {
		return low
	}

	return -1
}
