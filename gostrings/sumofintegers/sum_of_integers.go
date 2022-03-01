package sumofintegers

import (
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	reNum := regexp.MustCompile(`[0-9]+`)
	nums := reNum.FindAllString(strng, -1)
	sum := 0

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic("SumOfIntegersInString: " + err.Error())
		}
		sum += n
	}
	return sum
}
