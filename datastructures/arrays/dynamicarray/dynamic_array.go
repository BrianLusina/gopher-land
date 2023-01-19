package dynamicarray

func dynamicArray(n int, queries [][]int) []int {
	lastAnswer := 0
	arr := make([][]int, n)
	answers := []int{}

	for _, query := range queries {
		queryType, x, y := query[0], query[1], query[2]

		if queryType == 1 {
			idx := (x ^ lastAnswer) % n
			arr[idx] = append(arr[idx], y)
		} else if queryType == 2 {
			idx := (x ^ lastAnswer) % n
			lastAnswer = arr[idx][y%len(arr[idx])]
			answers = append(answers, lastAnswer)
		}
	}

	return answers
}
