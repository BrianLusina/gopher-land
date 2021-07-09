package binarytrees

// Walk walks the tree and sends all values to the channel intChannel
func Walk(tree *BinaryTreeNode, intChannel chan int) {
	recursivelyWalk(tree, intChannel)
	close(intChannel)
}

// recursivelyWalk recursively walks down a tree and pushes values to the channel
func recursivelyWalk(tree *BinaryTreeNode, intChannel chan int) {
	if tree != nil {
		// walk to the left first and send left part to the channel
		recursivelyWalk(tree.Left, intChannel)
		// send value to channel
		intChannel <- tree.Data.(int)
		// walk to the right
		recursivelyWalk(tree.Right, intChannel)
	}
}

// AreTreesSame determines whether 2 trees are the same
func AreTreesSame(tree1, tree2 *BinaryTreeNode) bool {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go Walk(tree1, channel1)
	go Walk(tree2, channel2)

	for {
		x1, ok1 := <-channel1
		x2, ok2 := <-channel2

		switch {
		case ok1 != ok2:
			// not the same size, as one channel has closed before the other channel
			return false

		case !ok1:
			// both channels are empty
			return true

		case x1 != x2:
			// elements don't match
			return false

		default:
			// keep going
		}
	}
}

// func main() {
// 	ch := make(chan int)
// 	go Walk(NewBinaryTree(1), ch)
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// 	fmt.Println(AreTreesSame(NewBinaryTree(1), NewBinaryTree(1)))
// 	fmt.Println(AreTreesSame(NewBinaryTree(1), NewBinaryTree(2)))
// }
