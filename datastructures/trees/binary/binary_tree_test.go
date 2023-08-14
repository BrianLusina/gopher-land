package binary

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

func TestBinaryTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binary Tree Test Suite")
}

var _ = Describe("Binary Tree", func() {
	Context("Deserialize & Serialize", func() {
		Describe("Serialize", func() {
			// FIXME: skipped test
			XIt("should return [1,2,3,nil,nil,4,5] for binary tree", func() {
				binaryTree := NewBinaryTree[int]()

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

	Context("IsFull", func() {
		When("Binary Tree contains nodes with type int", func() {
			It("should return false for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int]()

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(false))
			})

			It("should return false for binary tree with no root efficiently", func() {
				experiment := gmeasure.NewExperiment("IsFull binary check on no root")

				//Register the experiment as a ReportEntry - this will cause Ginkgo's reporter infrastructure
				//to print out the experiment's report and to include the experiment in any generated reports
				// AddReportEntry(experiment.Name, experiment)

				experiment.SampleDuration("IsFull", func(idx int) {
					binaryTree := NewBinaryTree[int]()
					actual := binaryTree.IsFull()

					Expect(actual).To(Equal(false))

				}, gmeasure.SamplingConfig{N: 20})
			})

			It("should return true for binary tree with root with no children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)

				binaryTree.root = root

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, 1 on the left and 1 on the right", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)

				binaryTree.root = root

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with root with 1 child, 1 on the left and 0 on the right", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)

				binaryTree.root = root

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(false))
			})
		})
	})

	Context("IsComplete", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return true for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int]()

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with no root efficiently", func() {
				experiment := gmeasure.NewExperiment("IsFull binary check on no root")

				//Register the experiment as a ReportEntry - this will cause Ginkgo's reporter infrastructure
				//to print out the experiment's report and to include the experiment in any generated reports
				// AddReportEntry(experiment.Name, experiment)

				experiment.SampleDuration("IsFull", func(idx int) {
					binaryTree := NewBinaryTree[int]()
					actual := binaryTree.IsFull()

					Expect(actual).To(Equal(false))

				}, gmeasure.SamplingConfig{N: 20})
			})

			It("should return true for binary tree with root with no children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)

				binaryTree.root = root

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)

				binaryTree.root = root

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 3 grandchildren", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)
				root.Left.Left = NewBinaryTreeNode(4)
				root.Left.Right = NewBinaryTreeNode(5)
				root.Right.Left = NewBinaryTreeNode(6)

				binaryTree.root = root

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})
		})
	})

	Context("IsPerfect", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return false for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int]()

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(false))
			})

			It("should return true for binary tree with root with no children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)

				binaryTree.root = root

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)

				binaryTree.root = root

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 4 grandchildren", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)
				root.Left.Left = NewBinaryTreeNode(4)
				root.Left.Right = NewBinaryTreeNode(5)
				root.Right.Left = NewBinaryTreeNode(6)
				root.Right.Right = NewBinaryTreeNode(7)

				binaryTree.root = root

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with root with 2 children, & 3 grandchildren", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)
				root.Left.Left = NewBinaryTreeNode(4)
				root.Left.Right = NewBinaryTreeNode(5)
				root.Right.Left = NewBinaryTreeNode(6)

				binaryTree.root = root

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(false))
			})
		})
	})

	Context("IsBalanced", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return true for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int]()

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with no children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)

				binaryTree.root = root

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)

				binaryTree.root = root

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 4 grandchildren", func() {
				binaryTree := NewBinaryTree[int]()

				root := NewBinaryTreeNode(1)
				root.Left = NewBinaryTreeNode(2)
				root.Right = NewBinaryTreeNode(3)
				root.Left.Left = NewBinaryTreeNode(4)
				root.Left.Right = NewBinaryTreeNode(5)

				binaryTree.root = root

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})
		})
	})

	Context("Height", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return 0 for no root", func() {
				tree := NewBinaryTree[int]()
				actual := tree.Height()
				Expect(actual).To(Equal(0))
			})

			It("should return 1 if the binary tree has a root, but no left nor right subtrees", func() {
				root := NewBinaryTreeNode(1)
				tree := NewBinaryTree[int]()
				tree.root = root
				actual := tree.Height()
				Expect(actual).To(Equal(1))
			})

			It("should return 3 for tree of 3,9,20,null,null,15,7", func() {
				left := NewBinaryTreeNode(9)
				rightLeft := NewBinaryTreeNode(15)
				rightRight := NewBinaryTreeNode(7)

				right := NewBinaryTreeNode(20)
				right.Left = rightLeft
				right.Right = rightRight

				root := NewBinaryTreeNode(3)
				root.Left = left
				root.Right = right

				tree := NewBinaryTree[int]()
				tree.root = root

				actual := tree.Height()
				Expect(actual).To(Equal(3))
			})

			It("should return 1 for tree of 1,null,2", func() {
				right := NewBinaryTreeNode(2)

				root := NewBinaryTreeNode(1)
				root.Right = right

				tree := NewBinaryTree[int]()
				tree.root = root

				actual := tree.Height()

				Expect(actual).To(Equal(2))
			})
		})
	})
})
