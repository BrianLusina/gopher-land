package twosum

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
