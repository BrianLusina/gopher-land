package leastnumberofuniqueints

import (
	"sort"
)

// findLeastNumOfUniqueInts takes in an array of integers and an integer k, and returns the least number of unique integers after removing exactly k elements from the array.
//
// The function first counts the frequency of each integer in the array, then sorts these frequencies in ascending order. It iteratively removes the least frequent integers
// until k elements have been removed, and returns the count of remaining unique integers.
//
// Complexity:
//
// Time Complexity: O(n log n) due to sorting the frequencies
//
// Space Complexity: O(n) for storing the frequency map and frequencies slice
func findLeastNumOfUniqueInts(arr []int, k int) int {
	// Count frequency of each element in the array
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	// Create a slice to store the frequencies of the integers
	frequencies := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		frequencies = append(frequencies, freq)
	}

	// Sort frequencies in ascending order to remove least frequent elements first
	sort.Ints(frequencies)

	// Iterate through sorted frequencies
	for index, freq := range frequencies {
		// Subtract current frequency from k (simulate removing all occurrences)
		k -= freq
		if k < 0 {
			// Return remaining unique integers (total unique - number of completely removed elements)
			return len(frequencies) - index
		}
	}

	// If we removed all elements (k >= sum of all frequencies)
	return 0
}
