package longestconsecutiveseq

import "math"

// longestConsecutive returns the length longest consecutive sequence in a slice of integers
// Time Complexity: O(n)
// Although the time complexity appears to be quadratic due to the while loop nested within the for loop, closer inspection reveals it to be linear.
// Because the while loop is reached only when currentNum marks the beginning of a sequence (i.e. currentNum-1 is not present in nums),
// the while loop can only run for nnn iterations throughout the entire runtime of the algorithm.
// This means that despite looking like O(n⋅n) complexity, the nested loops actually run in O(n+n)=O(n) time.
// All other computations occur in constant time, so the overall runtime is linear.
// Space Complexity: O(n)
// In order to set up O(1) containment lookups, we allocate linear space for a hash table to store the O(n) numbers in nums.
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longest := 0
	// O(n) space
	numSet := map[int]bool{}

	for _, n := range nums {
		if !numSet[n] {
			numSet[n] = true
		}
	}

	for _, number := range nums {
		if !numSet[number-1] {
			currentStreak := 0
			for numSet[number+currentStreak] {
				currentStreak++
			}
			longest = int(math.Max(float64(currentStreak), float64(longest)))
		}
	}

	return longest
}
