package pivotindex

func pivotIndex(nums []int) int {
	sumLeft := 0
	sumRight := 0

	for _, num := range nums {
		sumRight += num
	}

	for i, num := range nums {
		sumRight -= num
		if sumLeft == sumRight {
			return i
		}
		sumLeft += num
	}

	return -1
}

// pivotIndex1 is the first iteration of the solution to this problem. Even though this works, it's quite slow
// Time: O(n^2)
// Space: O(n)
func pivotIndexOne(nums []int) int {
	sumLeft := []int{}
	sumRight := []int{}

	for i := 0; i < len(nums); i++ {
		left := sum(nums[0:i])
		right := sum(nums[i+1:])

		sumLeft = append(sumLeft, left)
		sumRight = append(sumRight, right)
	}

	for i := range nums {
		if sumLeft[i] == sumRight[i] {
			return i
		}
	}

	return -1
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
