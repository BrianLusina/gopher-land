package highlow

import (
	"fmt"
	"strconv"
	"strings"
)

func findMinAndMax(nums []int) (min, max int) {
	min = nums[0]
	max = nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max
}

func HighAndLow(in string) string {
	numsStrArr := strings.Split(in, " ")
	nums := []int{}

	for _, num := range numsStrArr {
		n, _ := strconv.ParseInt(num, 10, 0)

		nums = append(nums, int(n))
	}

	min, max := findMinAndMax(nums)

	return fmt.Sprintf("%d %d", max, min)
}
