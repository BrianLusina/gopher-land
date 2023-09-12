package numberofprovinces

type unionFind struct {
	root []int
	rank []int
	size int
}

func newUnionFind(size int) *unionFind {
	rank := make([]int, size)
	root := make([]int, size)

	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}

	return &unionFind{
		root: root,
		size: size,
		rank: rank,
	}
}

func (uf *unionFind) find(x int) int {
	if x == uf.root[x] {
		return x
	}
	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *unionFind) union(x, y int) {
	xRoot := uf.find(x)
	yRoot := uf.find(y)

	if xRoot == yRoot {
		return
	}

	if uf.rank[xRoot] < uf.rank[yRoot] {
		uf.root[xRoot] = yRoot
	} else if uf.rank[xRoot] > uf.rank[yRoot] {
		uf.root[yRoot] = xRoot
	} else {
		uf.root[yRoot] = xRoot
		uf.rank[xRoot]++
	}
	uf.size--
}

func numberOfProvinces(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	nrows := len(grid)
	uf := newUnionFind(nrows)

	for row := 0; row < nrows; row++ {
		for col := row; col < nrows; col++ {
			if grid[row][col] == 1 {
				uf.union(row, col)
			}
		}
	}

	return uf.size
}
