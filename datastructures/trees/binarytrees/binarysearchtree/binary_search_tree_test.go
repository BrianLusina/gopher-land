package binarysearchtree

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinarySearchTree Suite")
}

var _ = Describe("BinarySearchTree", func() {
	It("should correctly insert new node with no root", func() {
		bst := NewBinarySearchTree[int]()
		bst.Insert(1)

		// root should not be nil
		Expect(bst.root).To(Not(BeNil()))
	})

	It("should correctly insert new nodes to left and right or root", func() {
		bst := NewBinarySearchTree[int]()
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
		bst := NewBinarySearchTree[int]()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(6)

		// root should not be nil
		Expect(bst.root).To(Not(BeNil()))
		Expect(bst.root.Data).To(Equal(4))
		Expect(bst.root.Left.Data).To(Equal(2))
		Expect(bst.root.Right.Data).To(Equal(6))

		// should return sorted values
		Expect(bst.Values()).To(Equal([]int{2, 4, 6}))
	})

	It("should correctly return depth of tree", func() {
		bst := NewBinarySearchTree[int]()
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
		bst := NewBinarySearchTree[int]()

		actual := bst.Serialize()
		Expect(actual).To(BeNil())
	})

	It("should serialize and return slice of values ordered as inserted", func() {
		bst := NewBinarySearchTree[int]()
		bst.Insert(4)
		bst.Insert(2)
		bst.Insert(6)

		actual := bst.Serialize()
		expected := []string{"4", "2", "6"}
		Expect(actual).To(Equal(expected))
	})

	Context("Delete", func() {
		It("should correctly delete values from tree", func() {
			bst := NewBinarySearchTree[int]()
			bst.Insert(50)
			bst.Insert(25)
			bst.Insert(75)
			bst.Insert(10)
			bst.Insert(33)
			bst.Insert(56)
			bst.Insert(89)
			bst.Insert(4)
			bst.Insert(11)
			bst.Insert(30)
			bst.Insert(40)
			bst.Insert(52)
			bst.Insert(61)
			bst.Insert(82)
			bst.Insert(95)

			actualErr := bst.Delete(4)
			assert.NoError(GinkgoT(), actualErr)

			// start at root
			Expect(bst.root).To(Not(BeNil()))
			Expect(bst.root.Data).To(Equal(50))

			// 2nd level check, left should be 25 and right should be 75
			Expect(bst.root.Left.Data).To(Equal(25))
			Expect(bst.root.Right.Data).To(Equal(75))

			// 3rd level checks

			// check left subtree with root 25
			Expect(bst.root.Left.Left.Data).To(Equal(10))
			Expect(bst.root.Left.Right.Data).To(Equal(33))

			// check left node 10
			// what we actually deleted is a left child of node 10
			Expect(bst.root.Left.Left.Left).To(BeNil())
			Expect(bst.root.Left.Left.Right.Data).To(Equal(11))

			// check right node 33
			Expect(bst.root.Left.Right.Left.Data).To(Equal(30))
			Expect(bst.root.Left.Right.Right.Data).To(Equal(40))

			// check node 75 from 2nd level
			Expect(bst.root.Right.Left.Data).To(Equal(56))
			Expect(bst.root.Right.Right.Data).To(Equal(89))

			// check the node 56 (left child of node 75)
			Expect(bst.root.Right.Left.Left.Data).To(Equal(52))
			Expect(bst.root.Right.Left.Right.Data).To(Equal(61))

			// check the node 89, right child of 75
			Expect(bst.root.Right.Right.Left.Data).To(Equal(82))
			Expect(bst.root.Right.Right.Right.Data).To(Equal(95))
		})
	})

	// XIt("should deserialize and construct the tree with values", func() {
	// 	data := "4 2 6"
	// 	bst := NewBinarySearchTree[int]()

	// 	bst.Deserialize(data)

	// 	expected := &BinarySearchTree[int]{
	// 		root: &binarytrees.BinaryTreeNode[int]{
	// 			TreeNode: trees.TreeNode[int]{
	// 				Data: 4,
	// 			},
	// 			Left: &binarytrees.BinaryTreeNode[int]{
	// 				TreeNode: trees.TreeNode[int]{
	// 					Data: 2,
	// 				},
	// 			},
	// 			Right: &binarytrees.BinaryTreeNode[int]{
	// 				TreeNode: trees.TreeNode[int]{
	// 					Data: 6,
	// 				},
	// 			},
	// 		},
	// 	}

	// 	Expect(bst.root).To(Equal(expected.root))
	// })

	// XIt("should return nil root for empty data input when deserializing", func() {
	// 	data := ""
	// 	bst := NewBinarySearchTree[int]()

	// 	bst.Deserialize(data)

	// 	Expect(bst.root).To(BeNil())
	// })
})
