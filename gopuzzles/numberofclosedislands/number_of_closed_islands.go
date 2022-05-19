package numberofclosedislands

const NUMBER_OF_DIRECTIONS = 4

var rowOffset = []int{1, -1, 0, 0}
var colOffset = []int{0, 0, 1, -1}

func numberOfClosedIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	nrows, ncols := len(grid), len(grid[0])
	visited := make([][]bool, nrows)

	for i := range visited {
		visited[i] = make([]bool, ncols)
	}

	for row := range grid {
		for col := range grid[row] {
			if !visited[row][col] && grid[row][col] == 0 && dfs(grid, row, col, visited) {
				islands++

			}
		}
	}

	return islands
}

func dfs(grid [][]int, row, col int, visited [][]bool) bool {
	visited[row][col] = true

	if grid[row][col] == 0 {
		isClosed := true

		for i := 0; i < NUMBER_OF_DIRECTIONS; i++ {
			newRow, newCol := row+rowOffset[i], col+colOffset[i]
			if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
				isClosed = false
			} else if !visited[newRow][newCol] {
				isClosed = dfs(grid, newRow, newCol, visited) && isClosed
			}
		}
		return isClosed
	}

	return true
}
