package numberofislands

func NumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	nrows, ncols := len(grid), len(grid[0])
	islands := 0

	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(&grid, r, c, nrows, ncols)
			}
		}
	}

	return islands
}

func dfs(grid *[][]byte, r, c, nrows, ncols int) {
	if r < 0 || r >= nrows || c < 0 || c >= ncols || (*grid)[r][c] == '0' {
		return
	}
	(*grid)[r][c] = '0'
	dfs(grid, r+1, c, nrows, ncols)
	dfs(grid, r-1, c, nrows, ncols)
	dfs(grid, r, c+1, nrows, ncols)
	dfs(grid, r, c-1, nrows, ncols)
}
