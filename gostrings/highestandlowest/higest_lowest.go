package highestandlowest

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Nums []int

func (a Nums) Len() int           { return len(a) }
func (a Nums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nums) Less(i, j int) bool { return a[i] < a[j] }

func HighAndLow(in string) string {
	s := strings.Split(in, " ")
	nums := make(Nums, 0, len(s))

	for _, v := range s {
		i, _ := strconv.Atoi(v)
		nums = append(nums, i)
	}
	sort.Sort(Nums(nums))
	return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}
