package intersection

import "gopherland/pkg/types"

// intersectTwo finds the intersection between 2 slices that have comparable types.
// Complexity Analysis:
//
//	Time Complexity:
//		This will depend on the sizes of each list. If both lists are identical in size, the Worst case scenario is
//		O(N^2) where N is the size of each list, this is because we are comparing each value in the first list to
//		each value in the second list. If they are not identical in size, then the complexity becomes O(N * M) where
//		N is the size of the first list and M the size of the second list.
//	Space Complexity:
//		Again, this will depend on the sizes of each list, but it averages to O(N+M) where N is the size of the first
//		list and M is the size of the second list
func intersectTwo[T types.Comparable](a []T, b []T) []T {
	result := []T{}

	for x, val := range a {
		for y := range b {
			if a[x] == b[y] {
				result = append(result, val)
				// adding break here ensures that we save on time and steps. if we have found a value that is identical
				// then there is no need to perform another iteration to check another value, We should proceed to the
				// next value
				break
			}
		}
	}

	return result
}
