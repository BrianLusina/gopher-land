package twosum

//TwoSum returns the indices of 2 numbers in a slice that add up to the target integer
// Uses O(n) Space complexity as a map is used to store the seen integers
func TwoSum(numbers []int, target int) [2]int {
	seen := make(map[int]int)

	for index, number := range numbers {
		diff := target - number

		res, ok := seen[diff]

		if ok {
			return [2]int{res, index}
		}

		seen[number] = index
	}
	return [2]int{0, 0}
}
