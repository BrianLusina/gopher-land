package rectangles

func Count(diagram []string) int {

	count := 0
	rows := len(diagram)
	cols := 0
	if rows > 0 {
		cols = len(diagram[0])
	}

	for r := 0; r < rows-1; r++ {
		for c := 0; c < cols-1; c++ {
			if diagram[r][c] == '+' {
				for r1 := r + 1; r1 < rows; r1++ {
					for c1 := c + 1; c1 < cols; c1++ {
						if diagram[r1][c1] == '+' && diagram[r][c1] == '+' && diagram[r1][c] == '+' {
							if checkHorizontal(diagram, c+1, c1, r) &&
								checkHorizontal(diagram, c+1, c1, r1) &&
								checkVertical(diagram, r+1, r1, c) &&
								checkVertical(diagram, r+1, r1, c1) {
								count++
							}
						}
					}
				}
			}
		}
	}

	return count
}

func checkHorizontal(diagram []string, start, end, rowIdx int) bool {
	for idx := start; idx < end; idx++ {
		point := diagram[rowIdx][idx]
		if point != '+' && point != '-' {
			return false
		}
	}
	return true
}

func checkVertical(diagram []string, start, end, colIdx int) bool {
	for idx := start; idx < end; idx++ {
		point := diagram[idx][colIdx]
		if point != '+' && point != '|' {
			return false
		}
	}
	return true
}
