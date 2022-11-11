package threesum

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for idx := 0; idx < len(nums)-2; idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		target, left, right := -nums[idx], idx+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]

			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				result = append(result, []int{nums[idx], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}
