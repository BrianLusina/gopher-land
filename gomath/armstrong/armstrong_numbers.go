package armstrong

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {
	nums := strconv.Itoa(n)
	l := len(nums)
	numbers := strings.Split(nums, "")
	sum := 0
	for _, num := range numbers {
		digit, _ := strconv.Atoi(num)
		sum += int(math.Pow(float64(digit), float64(l)))
	}
	return n == sum
}
