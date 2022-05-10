package pov

type Tree struct {
	value    string
	children []*Tree
	parent   *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	if value == "" {
		return nil
	}

	tree := &Tree{value: value, children: children, parent: nil}
	for _, child := range children {
		if child != nil {
			child.parent = tree
		}
	}
	return tree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	if tr == nil {
		return ""
	}
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	if tr == nil {
		return nil
	}
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

func (tr *Tree) findNode(nodeValue string) *Tree {
	if tr == nil || tr.Value() == nodeValue {
		return tr
	}
	for _, child := range tr.Children() {
		result := child.findNode(nodeValue)
		if result != nil {
			return result
		}
	}
	return nil
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	if tr.Value() == from {
		return tr
	}

	node := tr.findNode(from)
	visited := make(map[string]bool)

	var traverseTree func(*Tree) *Tree
	traverseTree = func(tree *Tree) *Tree {
		if tree == nil || visited[tree.value] {
			return nil
		}
		visited[tree.value] = true
		children := make([]*Tree, 0, len(tr.children)+1)

		for _, child := range tree.Children() {
			fchild := traverseTree(child)
			if fchild != nil {
				children = append(children, fchild)
			}
		}

		if tree.parent != nil {
			parent := traverseTree(tree.parent)
			if parent != nil {
				children = append(children, parent)
			}
		}

		return New(tree.value, children...)
	}

	return traverseTree(node)
}

func (tr *Tree) pathFromRoot(value string) []string {
	node := tr.findNode(value)

	if node == nil {
		return nil
	}

	result := make([]string, 0)

	for node != tr {
		result = append(result, node.value)
		node = node.parent
	}

	result = append(result, tr.value)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	if tr == nil {
		return nil
	}

	if from == to {
		return []string{from}
	}

	fromPov := tr.FromPov(from)
	if fromPov == nil {
		return nil
	}

	return fromPov.pathFromRoot(to)
}
