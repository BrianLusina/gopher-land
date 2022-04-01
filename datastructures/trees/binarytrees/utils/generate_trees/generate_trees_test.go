package utils

import (
	"gopherland/datastructures/trees"
	"testing"

	. "gopherland/datastructures/trees/binarytrees"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenerateTrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate Trees Test Suite")
}

var _ = Describe("Generate Trees Utility function", func() {
	It("should return tree with root node only for input of 1", func() {
		expected := []*BinaryTreeNode{{TreeNode: trees.TreeNode{Data: 1}}}
		actual := GenerateTrees(1)
		Expect(actual).To(Equal(expected))
	})

	It("should return 5 trees with input of 3", func() {
		expected := []*BinaryTreeNode{
			{
				TreeNode: trees.TreeNode{
					Data: 1,
				},
				Left: nil,
				Right: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 2,
					},
					Left: nil,
					Right: &BinaryTreeNode{
						TreeNode: trees.TreeNode{
							Data: 3,
						},
						Left:  nil,
						Right: nil,
					},
				},
			},
			{
				TreeNode: trees.TreeNode{
					Data: 1,
				},
				Left: nil,
				Right: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 3,
					},
					Left: &BinaryTreeNode{
						TreeNode: trees.TreeNode{
							Data: 2,
						},
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			{
				TreeNode: trees.TreeNode{
					Data: 2,
				},
				Left: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 1,
					},
					Left:  nil,
					Right: nil,
				},
				Right: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 3,
					},
					Left:  nil,
					Right: nil,
				},
			},
			{
				TreeNode: trees.TreeNode{
					Data: 3,
				},
				Left: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 1,
					},
					Left: nil,
					Right: &BinaryTreeNode{
						TreeNode: trees.TreeNode{
							Data: 2,
						},
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
			{
				TreeNode: trees.TreeNode{
					Data: 3,
				},
				Left: &BinaryTreeNode{
					TreeNode: trees.TreeNode{
						Data: 2,
					},
					Left: &BinaryTreeNode{
						TreeNode: trees.TreeNode{
							Data: 1,
						},
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
		}
		actual := GenerateTrees(3)
		Expect(actual).To(Equal(expected))
	})
})
