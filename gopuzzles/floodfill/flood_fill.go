package floodfill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return image
	}

	nrows, ncols := len(image), len(image[0])

	startPixelColor := image[sr][sc]

	if startPixelColor == newColor {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if image[r][c] == startPixelColor {
			image[r][c] = newColor

			if r >= 1 {
				dfs(r-1, c)
			}

			if r+1 < nrows {
				dfs(r+1, c)
			}

			if c >= 1 {
				dfs(r, c-1)
			}

			if c+1 < ncols {
				dfs(r, c+1)
			}
		}
	}

	dfs(sr, sc)

	// dfs(&image, sr, sc, nrows, ncols, startPixelColor, newColor)

	return image
}

func dfs(grid *[][]int, r, c, nrows, ncols, startPixelColor, newColor int) {
	if (*grid)[r][c] == startPixelColor {
		(*grid)[r][c] = newColor

		if r >= 1 {
			dfs(grid, r-1, c, nrows, ncols, startPixelColor, newColor)
		}

		if r+1 < nrows {
			dfs(grid, r+1, c, nrows, ncols, startPixelColor, newColor)
		}

		if c >= 1 {
			dfs(grid, r, c-1, nrows, ncols, startPixelColor, newColor)
		}

		if c+1 < ncols {
			dfs(grid, r, c+1, nrows, ncols, startPixelColor, newColor)
		}
	}
}
