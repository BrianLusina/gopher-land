package binarytrees

func mergeTrees(root1 *BinaryTreeNode, root2 *BinaryTreeNode) *BinaryTreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	r2, _ := root2.Data.(int)

	r1, _ := root1.Data.(int)

	r1 += r2

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
