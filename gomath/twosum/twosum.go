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

// twoSumII finds the indices of two integers from a 1-index based numbers slice sorted in increasing order that add up to a target
// This uses a binary search to find the two numbers that add up to the target with the left and right pointer at both ends of thes lice
// Space complexity is O(2) where we store 2 numbers in the return result, and Time complexity is O(n log(n)) as a binary search algorithm
// is used to find the two numbers
func twoSumII(numbers []int, target int) [2]int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return [2]int{left + 1, right + 1}
		}
	}

	return [2]int{0, 0}
}
