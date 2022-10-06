package floodfill

// floodFill performs the algorithm by painting the starting pixels, plus adjacent pixels of the same color, and so on.
// Say 'color' is the color of the starting pixel. Let's floodfill the starting pixel: we change the color of that
// pixel to the new color, then check the 4 neighboring pixels to make sure they are valid pixels of the same color,
// and of the valid ones, we floodfill those, and so on. We can use a function dfs to perform a floodfill on a target pixel.
// Complexity Analysis:
// Time Complexity: O(N), where NN is the number of pixels in the image. We might process every pixel.
// Space Complexity: O(N), the size of the implicit call stack when calling dfs.
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
