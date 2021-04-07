package kthlargestelement

import "sort"

type sortBy []int

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[j] < a[i] }

func sortNumbers(nums []int) []int {
	sort.Sort(sortBy(nums))
	return nums
}

func FindKthLargest(nums []int, k int) int {
	sorted := sortNumbers(nums)
	return sorted[k-1]
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
