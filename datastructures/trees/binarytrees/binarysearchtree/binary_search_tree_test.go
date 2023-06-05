package binarysearchtree

import (
	"gopherland/datastructures/trees"
	"gopherland/datastructures/trees/binarytrees"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinarySearchTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TrieNode Suite")
}

var _ = Describe("BinarySearchTree", func() {
	It("should correctly insert new node with no root", func() {
		bst := NewBinarySearchTree()
		bst.Insert(1)

		// root should not be nil
		Expect(bst.root).To(Not(BeNil()))
	})

	It("should correctly insert new nodes to left and right or root", func() {
		bst := NewBinarySearchTree()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(6)

		// root should not be nil
		Expect(bst.root).To(Not(BeNil()))
		Expect(bst.root.Data).To(Equal(4))
		Expect(bst.root.Left.Data).To(Equal(2))
		Expect(bst.root.Right.Data).To(Equal(6))
	})

	It("should correctly return values in sorted order", func() {
		bst := NewBinarySearchTree()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(6)

		// root should not be nil
		Expect(bst.root).To(Not(BeNil()))
		Expect(bst.root.Data).To(Equal(4))
		Expect(bst.root.Left.Data).To(Equal(2))
		Expect(bst.root.Right.Data).To(Equal(6))

		// should return sorted values
		Expect(bst.Values()).To(Equal([]interface{}{2, 4, 6}))
	})

	It("should correctly return depth of tree", func() {
		bst := NewBinarySearchTree()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(1)
		bst.Insert(6)
		bst.Insert(3)
		bst.Insert(5)
		bst.Insert(7)

		expected := 3
		actual := bst.Depth()

		// should return sorted values
		Expect(actual).To(Equal(expected))
	})

	It("should serialize and return nil if there are no nodes", func() {
		bst := NewBinarySearchTree()

		actual := bst.Serialize()
		Expect(actual).To(BeNil())
	})

	It("should serialize and return slice of values ordered as inserted", func() {
		bst := NewBinarySearchTree()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(6)

		actual := bst.Serialize()
		expected := []string{"4", "2", "6"}
		Expect(actual).To(Equal(expected))
	})

	It("should deserialize and construct the tree with values", func() {
		data := "4 2 6"
		bst := NewBinarySearchTree()

		bst.Deserialize(data)

		expected := &BinarySearchTree{
			root: &binarytrees.BinaryTreeNode[int]{
				TreeNode: trees.TreeNode[int]{
					Data: 4,
				},
				Left: &binarytrees.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 2,
					},
				},
				Right: &binarytrees.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 6,
					},
				},
			},
		}

		Expect(bst.root).To(Equal(expected.root))
	})

	It("should return nil root for empty data input when deserializing", func() {
		data := ""
		bst := NewBinarySearchTree()

		bst.Deserialize(data)

		Expect(bst.root).To(BeNil())
	})
})
