package intbst

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left != nil {
			bst.left.Insert(i)
		} else {
			bst.left = NewBst(i)
		}
	} else {
		if bst.right != nil {
			bst.right.Insert(i)
		} else {
			bst.right = NewBst(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	var result []int
	if bst.left != nil {
		result = append(bst.left.SortedData(), result...)
	}
	result = append(result, bst.data)
	if bst.right != nil {
		result = append(result, bst.right.SortedData()...)
	}
	return result
}
