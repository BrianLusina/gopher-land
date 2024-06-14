package optimaltaskassignment

import "sort"

type Pair struct {
	first  int
	second int
}

/*
*
After sorting the array, the for loop iterates for half the length of the array adds the pairs using indexing to a
list. So ~x is the bitwise complement operator which puts a negative sign in front of x and subtracts 1 from it.
Thus, the negative numbers as indexes mean that you count from the right of the array instead of the left. So,
numbers[-1] refers to the last element, numbers[-2] is the second-last, and so on. In this way, we are able to pair
the numbers from the beginning of the array to the end of the array.

	*
*/
func optimalTaskAssignment(tasks []int) []Pair {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})

	result := []Pair{}

	for index := 0; index < len(tasks)/2; index++ {
		first := tasks[index]
		second := tasks[len(tasks)+^index]
		result = append(result, Pair{first, second})
	}
	return result
}
