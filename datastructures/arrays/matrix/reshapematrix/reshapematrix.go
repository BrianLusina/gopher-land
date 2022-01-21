package reshapematrix

func flatten(matrix [][]int) []int {
	values := []int{}

	for _, row := range matrix {
		for c, _ := range row {
			values = append(values, row[c])
		}
	}
	return values
}

func createEmptyMatrix(rows, cols int) [][]int {
	matrix := [][]int{}

	for r := 0; r < rows; r++ {
		columns := make([]int, cols)
		matrix = append(matrix, columns)
	}
	return matrix
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	flattenedMatrix := flatten(mat)
	result := createEmptyMatrix(r, c)

	i := 0

	for row := 0; row < r; row++ {
		for col := 0; col < c; col++ {
			result[row][col] = flattenedMatrix[i]
			i++
		}
	}

	return result
}
