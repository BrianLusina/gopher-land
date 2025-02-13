package findduplicate

func findDuplicateSet(nums []int) int {
	visited := map[int]bool{}
	var duplicate int

	for _, num := range nums {
		_, ok := visited[num]
		if ok {
			duplicate = num
			break
		}
		visited[num] = true
	}
	return duplicate
}

// findDuplicate Finds the duplicate number in a list of integers
// This code implements Floyd's Tortoise and Hare algorithm to find the duplicate number in a list of integers. It
// uses two pointers, slow and fast, that move at different speeds through the list, eventually meeting at the
// duplicate number. The algorithm assumes that the list contains a duplicate number and that the numbers in the list
// are indices that point to other numbers in the list.
// Time Complexity: O(n)
// Space Complexity: O(1)
// Args:
// nums (list): list of integers
// Returns:
// int: duplicate number
func findDuplicate(nums []int) int {
	length := len(nums)

	if length == 0 || length <= 1 {
		return -1
	}

	for _, num := range nums {
		if num < 1 || num >= length {
			return -1 // number is not a valid index
		}
	}

	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		if slow < 0 || slow >= length {
			return -1
		}
		if fast < 0 || fast >= length {
			return -1
		}
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		if slow < 0 || slow >= length {
			return -1
		}
		if fast < 0 || fast >= length {
			return -1
		}
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
