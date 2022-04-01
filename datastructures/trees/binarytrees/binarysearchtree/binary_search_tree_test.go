package binarysearchtree

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trie Suite")
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
})
