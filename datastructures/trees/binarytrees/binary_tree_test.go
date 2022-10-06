package binarytrees

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBinaryTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binary Tree Test Suite")
}

var _ = Describe("Binary Tree", func() {
	Context("Deserialize & Serialize", func() {
		Describe("Serialize", func() {
			It("should return [1,2,3,nil,nil,4,5] for binary tree", func() {
				binaryTree := NewBinaryTree()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)

				root.Right.Left = NewBinaryTreeNode(4)
				root.Right.Right = NewBinaryTreeNode(5)

				binaryTree.root = root

				actual := root.Serialize()
				expected := "1,2,3,nil,nil,4,5"

				Expect(actual).To(Equal(expected))
			})
		})
	})
})
