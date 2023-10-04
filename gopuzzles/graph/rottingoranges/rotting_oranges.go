package rottingoranges

// pair is the coordinate of a rotting orange
type pair struct {
	x, y int
}

// orangesRotting returns the minimum time it will take to make all oranges rotten in a grid. If this is not possible, -1 is returned
// Since we must track time, a breadth-first search approach makes the most sense, as we can track each iteration of
// our bfs loop, and update the time before moving on to the next.
//
// First, we need to iterate our grid looking for rotten oranges to add to our queue. Note, if we also take the time
// during this iteration to count the number of fresh oranges, we can save a little time in the end.
//
// Once we have a queue of oranges, we just need to run our BFS and update our time after we exhaust all current cells
// in each iteration of it, and finally return the time as long as we have successfully converted all fresh oranges to
// rotten ones.
//
// Time Complexity: O(m*n)
// where m is the number of rows, and n is the number of columns. We must traverse the graph once to find all our
// fresh/rotten oranges, and then potentially again during our BFS.
//
// Space Complexity: O(m*n). In the worst case, all oranges will be rotten, and our queue will be of size m*n.
func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// keep track of rotting oranges
	queue := []pair{}

	// number of fresh oranges
	fresh := 0

	// keep track of the time
	time := 0

	// directions in the form of x,y
	directions := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

	// iterate the grid to find the rotting oranges
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, pair{x: r, y: c})
			}

			if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	// while we have a queue of rotten oranges, and there are still fresh oranges.
	for len(queue) > 0 && fresh > 0 {
		for range queue {
			rotting_orange_coordinates := queue[0]
			queue = queue[1:]

			row, col := rotting_orange_coordinates.x, rotting_orange_coordinates.y

			// check the 4 adjacent oranges of this rotten orange
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]

				// position of adjacent orange
				r, c := row+dr, col+dc

				// check the adjacent orange is in bounds and fresh
				if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != 1 {
					continue
				}

				// they are fresh, therefore, it becomes rotten
				grid[r][c] = 2
				fresh--

				// add new rotten orange to the queue, we won't reach it on this iteration, as we are only iterating
				// through the rotten from the last iteration
				queue = append(queue, pair{x: r, y: c})
			}
		}
		// since we only looped through the rotten oranges inside the queue, and not the adjacent ones, we can increment
		// the time now, and on the next iteration we will check those adjacent ones
		time++
	}

	// if we turned all oranges rotten, return the time, else -1
	if fresh == 0 {
		return time
	}
	return -1
}
