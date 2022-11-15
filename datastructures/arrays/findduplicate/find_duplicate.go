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
