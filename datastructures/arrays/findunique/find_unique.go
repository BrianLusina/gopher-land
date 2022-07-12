package findunique

import "sort"

type sortBy []float32

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i] < a[j] }

// FindUniq returns the unique number in the array
func FindUniq(arr []float32) float32 {
	sort.Sort(sortBy(arr))

	first, second, last := arr[0], arr[1], arr[len(arr)-1]

	if first == second {
		return last
	} else {
		return first
	}
}
