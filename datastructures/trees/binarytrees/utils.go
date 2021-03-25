package binarytrees

func generator(start, end int) []*BinaryTreeNode {
	bsts := []*BinaryTreeNode{}

	if start > end {
		bsts = append(bsts, nil)
		return nil
	}

	for i := start; i < end+1; i++ {
		leftSubtree := generator(start, i-1)
		rightSubtree := generator(i+1, end)

		for j := range leftSubtree {
			left := leftSubtree[j]

			for k := range rightSubtree {
				right := rightSubtree[k]
				node := &BinaryTreeNode{
					Data: i,
				}
				node.Left = left
				node.Right = right
				bsts = append(bsts, node)
			}
		}
	}
	return bsts
}

// Generates n Binary Trees given a number n
func GenerateTrees(n int) []*BinaryTreeNode {
	return generator(1, n)
}
