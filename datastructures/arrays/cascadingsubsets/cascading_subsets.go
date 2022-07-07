package cascadingsubsets

func EachCons(arr []int, n int) [][]int {
	result := [][]int{}

	if len(arr) == 0 {
		return result
	}

	for x := 0; x <= len(arr)-n; x++ {
		chunk := []int{}
		for y := x; y < x+n; y++ {
			chunk = append(chunk, arr[y])
		}
		result = append(result, chunk)
	}

	return result
}
