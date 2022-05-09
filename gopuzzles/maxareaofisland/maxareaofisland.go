package maxareaofisland

import "math"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	nrows, ncols := len(grid), len(grid[0])
	maxarea := 0

	for r := 0; r < nrows; r++ {
		for c := 0; c < ncols; c++ {
			if grid[r][c] == 1 {
				maxarea = int(math.Max(float64(maxarea), float64(dfs(&grid, r, c, nrows, ncols))))
			}
		}
	}

	return maxarea
}

func dfs(grid *[][]int, r, c, nrows, ncols int) int {
	if r < 0 || r >= nrows || c < 0 || c >= ncols || (*grid)[r][c] == 0 {
		return 0
	}

	(*grid)[r][c] = 0

	return 1 + dfs(grid, r+1, c, nrows, ncols) +
		dfs(grid, r-1, c, nrows, ncols) +
		dfs(grid, r, c+1, nrows, ncols) +
		dfs(grid, r, c-1, nrows, ncols)
}
