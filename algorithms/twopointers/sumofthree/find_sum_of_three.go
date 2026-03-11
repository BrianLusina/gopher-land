package sumofthree

import (
	"sort"
)

func findSomeOfThree(numbers []int, target int) bool {
	sort.IntSlice(numbers).Sort()

	for i := 0; i < len(numbers)-2; i++ {
		left, right := i+1, len(numbers)-1
		for left < right {
			sum := numbers[i] + numbers[left] + numbers[right]
			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}

func threeNumberSum(nums []int, targetSum int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	n := len(nums)
	numbers := make([]int, n)
	copy(numbers, nums)
	sort.IntSlice(numbers).Sort()
	result := [][]int{}

	for idx := range n {
		num := numbers[idx]

		// Skip duplicate values to avoid duplicate triplets
		if idx > 0 && num == numbers[idx-1] {
			continue
		}

		left, right := idx+1, n-1

		for left < right {
			totalSum := num + numbers[left] + numbers[right]

			if totalSum == targetSum {
				result = append(result, []int{num, numbers[left], numbers[right]})
				// move the left pointer to avoid duplicates while it is still less than the right
				for left < right && numbers[left] == numbers[left+1] {
					left += 1
				}

				for left < right && numbers[right] == numbers[right-1] {
					right -= 1
				}

				left += 1
				right -= 1
			} else if totalSum > targetSum {
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return result
}
