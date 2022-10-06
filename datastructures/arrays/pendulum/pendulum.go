package pendulum

import "sort"

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func Pendulum(values []int) []int {
	output := make([]int, len(values))
	sort.IntSlice(values).Sort()
	n := len(values)

	for x, y, z := 0, (n-1)/2, n; x < len(output); x++ {
		if x%2 == 0 {
			output[y] = values[x]
			y--
		} else {
			output[z] = values[x]
			z++
		}
	}

	return output
}
