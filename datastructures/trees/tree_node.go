package trees

// TreeNode represent a Node in a Tree
type TreeNode struct {
	Data     any
	Children []*TreeNode
}

// NewTreeNode returns a new TreeNode
func NewTreeNode(data interface{}) TreeNode {
	return TreeNode{
		Data: data,
	}
}
