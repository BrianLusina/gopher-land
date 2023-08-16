package utils

import (
	"gopherland/datastructures/trees"
	"testing"

	"gopherland/datastructures/trees/binary"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenerateTrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate Trees Test Suite")
}

var _ = Describe("Generate Trees Utility function", func() {
	It("should return tree with root node only for input of 1", func() {
		expected := []*binary.BinaryTreeNode[int]{{TreeNode: trees.TreeNode[int]{Data: 1}}}
		actual := GenerateTrees(1)
		Expect(actual).To(Equal(expected))
	})

	It("should return 5 trees with input of 3", func() {
		expected := []*binary.BinaryTreeNode[int]{
			{
				TreeNode: trees.TreeNode[int]{
					Data: 1,
				},
				Left: nil,
				Right: &binary.NewBinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 2,
					},
					Left: nil,
					Right: &binary.BinaryTreeNode[int]{
						TreeNode: trees.TreeNode[int]{
							Data: 3,
						},
						Left:  nil,
						Right: nil,
					},
				},
			},
			{
				TreeNode: trees.TreeNode[int]{
					Data: 1,
				},
				Left: nil,
				Right: &binary.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 3,
					},
					Left: &binary.BinaryTreeNode[int]{
						TreeNode: trees.TreeNode[int]{
							Data: 2,
						},
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
			{
				TreeNode: trees.TreeNode[int]{
					Data: 2,
				},
				Left: &binary.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 1,
					},
					Left:  nil,
					Right: nil,
				},
				Right: &binary.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 3,
					},
					Left:  nil,
					Right: nil,
				},
			},
			{
				TreeNode: trees.TreeNode[int]{
					Data: 3,
				},
				Left: &binary.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 1,
					},
					Left: nil,
					Right: &binary.BinaryTreeNode[int]{
						TreeNode: trees.TreeNode[int]{
							Data: 2,
						},
						Left:  nil,
						Right: nil,
					},
				},
				Right: nil,
			},
			{
				TreeNode: trees.TreeNode[int]{
					Data: 3,
				},
				Left: &binary.BinaryTreeNode[int]{
					TreeNode: trees.TreeNode[int]{
						Data: 2,
					},
					Left: &binary.BinaryTreeNode[int]{
						TreeNode: trees.TreeNode[int]{
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
