package multiplicationtable

// MultiplicationTable returns an n*n multiplication table matrix
func MultiplicationTable(size int) [][]int {
	matrix := [][]int{}

	for x := 1; x <= size; x++ {
		row := []int{}
		for y := 1; y <= size; y++ {
			row = append(row, x*y)
		}
		matrix = append(matrix, row)
	}

	return matrix
}
