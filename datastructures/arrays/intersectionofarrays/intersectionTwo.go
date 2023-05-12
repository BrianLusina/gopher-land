package intersectionofarrays

import "gopherland/pkg/types"

func intersectTwo[T types.Comparable](a []T, b []T) []T {
	result := []T{}

	for x, val1 := range a {
		for y, _ := range b {
			if a[x] == b[y] {
				result = append(result, val1)
				break
			}
		}
	}

	return result
}
