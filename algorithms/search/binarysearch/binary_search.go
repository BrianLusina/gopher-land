// Package binarysearch contains functions with binary search algorithms
package binarysearch

import "gopherland/pkg/types"

// Search uses binary search algorithm to find the index of a given k in a collection. -1 is returned if the key can not be found
// The assumption made is that the collection is already sorted giving this algorithm a complexity of O(log(n)) performance with O(1)
// space complexity as no extra space is used other than variable assignment
func Search[T types.Comparable](collection []T, key T) int {
	if len(collection) == 0 {
		return -1
	}

	low, high := 0, len(collection)-1

	for low <= high {
		mid := (low + high) / 2
		if collection[mid] == key {
			return mid
		} else if collection[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
