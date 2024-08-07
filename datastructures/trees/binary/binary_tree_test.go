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
			It("should return [nil] for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int](nil)

				actual := binaryTree.Serialize()
				expected := "nil"

				Expect(actual).To(Equal(expected))
			})

			It("should return 10 86 nil nil 100 nil nil for root(10, left(86), right(100)", func() {
				root := NewBinaryTreeNode(10, Left(NewBinaryTreeNode(86)), Right(NewBinaryTreeNode(100)))
				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.Serialize()
				expected := "10,86,nil,nil,100,nil,nil"

				Expect(actual).To(Equal(expected))
			})

			It("should return 1 2 nil nil 3 x x for root(1, left(2), right(3) for binary tree", func() {
				root := NewBinaryTreeNode(1, Left(NewBinaryTreeNode(2)), Right(NewBinaryTreeNode(3)))

				binaryTree := NewBinaryTree(root)

				actual := binaryTree.Serialize()
				expected := "1,2,nil,nil,3,nil,nil"

				Expect(actual).To(Equal(expected))
			})

			It("should return 6 4 3 nil nil 5 nil nil 8 nil nil for root(6, left(4, left(3), right(5)), right(8)", func() {
				root := NewBinaryTreeNode(6, Left(NewBinaryTreeNode(4, Left(NewBinaryTreeNode(3)), Right(NewBinaryTreeNode(5)))), Right(NewBinaryTreeNode(8)))

				binaryTree := NewBinaryTree(root)

				actual := binaryTree.Serialize()
				expected := "6,4,3,nil,nil,5,nil,nil,8,nil,nil"

				Expect(actual).To(Equal(expected))
			})
		})

		Describe("Deserialize", func() {
			It("should return nil for a tree [nil]", func() {
				binaryTree := "nil"

				actual := Deserialize(binaryTree)

				Expect(actual).To(BeNil())
			})

			It("should return root(10, left(86), right(100) for [10,86,nil,nil,100,nil,nil]", func() {
				binaryTree := "10,86,nil,nil,100,nil,nil"
				actual := Deserialize(binaryTree)

				expected := NewBinaryTreeNode("10", Left(NewBinaryTreeNode("86")), Right(NewBinaryTreeNode("100")))

				Expect(actual).To(Equal(expected))
			})

			It("should return root(1, left(2), right(3) for [1,2,nil,nil,3,nil,nil]", func() {
				binaryTree := "1,2,nil,nil,3,nil,nil"

				actual := Deserialize(binaryTree)
				expected := NewBinaryTreeNode("1", Left(NewBinaryTreeNode("2")), Right(NewBinaryTreeNode("3")))

				Expect(actual).To(Equal(expected))
			})

			It("should return root(6, left(4, left(3), right(5)), right(8) for [6,4,3,nil,nil,5,nil,nil,8,nil,nil]", func() {
				binaryTree := "6,4,3,nil,nil,5,nil,nil,8,nil,nil"

				actual := Deserialize(binaryTree)
				expected := NewBinaryTreeNode("6", Left(NewBinaryTreeNode("4", Left(NewBinaryTreeNode("3")), Right(NewBinaryTreeNode("5")))), Right(NewBinaryTreeNode("8")))

				Expect(actual).To(Equal(expected))
			})
		})

	})

	Context("IsFull", func() {
		When("Binary Tree contains nodes with type int", func() {
			It("should return false for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int](nil)

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(false))
			})

			It("should return false for binary tree with no root efficiently", func() {
				experiment := gmeasure.NewExperiment("IsFull binary check on no root")

				//Register the experiment as a ReportEntry - this will cause Ginkgo's reporter infrastructure
				//to print out the experiment's report and to include the experiment in any generated reports
				// AddReportEntry(experiment.Name, experiment)

				experiment.SampleDuration("IsFull", func(idx int) {
					binaryTree := NewBinaryTree[int](nil)
					actual := binaryTree.IsFull()

					Expect(actual).To(Equal(false))

				}, gmeasure.SamplingConfig{N: 20})
			})

			It("should return true for binary tree with root with no children", func() {
				root := NewBinaryTreeNode(1)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, 1 on the left and 1 on the right", func() {

				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with root with 1 child, 1 on the left and 0 on the right", func() {
				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsFull()

				Expect(actual).To(Equal(false))
			})
		})
	})

	Context("IsComplete", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return true for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int](nil)

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with no root efficiently", func() {
				experiment := gmeasure.NewExperiment("IsFull binary check on no root")

				//Register the experiment as a ReportEntry - this will cause Ginkgo's reporter infrastructure
				//to print out the experiment's report and to include the experiment in any generated reports
				// AddReportEntry(experiment.Name, experiment)

				experiment.SampleDuration("IsFull", func(idx int) {
					binaryTree := NewBinaryTree[int](nil)
					actual := binaryTree.IsFull()

					Expect(actual).To(Equal(false))

				}, gmeasure.SamplingConfig{N: 20})
			})

			It("should return true for binary tree with root with no children", func() {
				root := NewBinaryTreeNode(1)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 3 grandchildren", func() {

				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)
				root.left.left = NewBinaryTreeNode(4)
				root.left.right = NewBinaryTreeNode(5)
				root.right.left = NewBinaryTreeNode(6)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsComplete()

				Expect(actual).To(Equal(true))
			})
		})
	})

	Context("IsPerfect", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return false for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int](nil)

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(false))
			})

			It("should return true for binary tree with root with no children", func() {
				root := NewBinaryTreeNode(1)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 4 grandchildren", func() {

				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)
				root.left.left = NewBinaryTreeNode(4)
				root.left.right = NewBinaryTreeNode(5)
				root.right.left = NewBinaryTreeNode(6)
				root.right.right = NewBinaryTreeNode(7)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(true))
			})

			It("should return false for binary tree with root with 2 children, & 3 grandchildren", func() {

				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)
				root.left.left = NewBinaryTreeNode(4)
				root.left.right = NewBinaryTreeNode(5)
				root.right.left = NewBinaryTreeNode(6)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsPerfect()

				Expect(actual).To(Equal(false))
			})
		})
	})

	Context("IsBalanced", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return true for binary tree with no root", func() {
				binaryTree := NewBinaryTree[int](nil)

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with no children", func() {
				root := NewBinaryTreeNode(1)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children", func() {
				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})

			It("should return true for binary tree with root with 2 children, & 4 grandchildren", func() {

				root := NewBinaryTreeNode(1)
				root.left = NewBinaryTreeNode(2)
				root.right = NewBinaryTreeNode(3)
				root.left.left = NewBinaryTreeNode(4)
				root.left.right = NewBinaryTreeNode(5)

				binaryTree := NewBinaryTree[int](root)

				actual := binaryTree.IsBalanced()

				Expect(actual).To(Equal(true))
			})
		})
	})

	Context("LeafSimilar", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return false for tree 1 having no root and tree 2 having a root", func() {
				root := NewBinaryTreeNode(1)
				tree1 := NewBinaryTree[int](root)

				tree2 := NewBinaryTree[int](nil)

				actual := tree1.LeafSimilar(tree2)
				Expect(actual).To(Equal(false))
			})

			It("should return true for tree1=3,5,1,6,2,9,8,null,null,7,4 and tree2=3,5,1,6,7,4,2,null,null,null,null,null,null,9,8", func() {
				left1 := NewBinaryTreeNode(5, Left(NewBinaryTreeNode(6)), Right(NewBinaryTreeNode(2, Left(NewBinaryTreeNode(7)), Right(NewBinaryTreeNode(4)))))
				right1 := NewBinaryTreeNode(1, Left(NewBinaryTreeNode(9)), Right(NewBinaryTreeNode(8)))

				root1 := NewBinaryTreeNode(3, Left(left1), Right(right1))
				tree1 := NewBinaryTree(root1)

				left2 := NewBinaryTreeNode(5, Left(NewBinaryTreeNode(6)), Right(NewBinaryTreeNode(7)))
				right2 := NewBinaryTreeNode(1, Left(NewBinaryTreeNode(4)), Right(NewBinaryTreeNode(2, Left(NewBinaryTreeNode(9)), Right(NewBinaryTreeNode(8)))))

				root2 := NewBinaryTreeNode(3, Left(left2), Right(right2))
				tree2 := NewBinaryTree(root2)

				actual := tree1.LeafSimilar(tree2)

				Expect(actual).To(Equal(true))
			})

			It("should return false for tree1=1,2,3 and tree2=1,3,2", func() {
				root1 := NewBinaryTreeNode(1, Left(NewBinaryTreeNode(2)), Right(NewBinaryTreeNode(3)))
				tree1 := NewBinaryTree(root1)

				root2 := NewBinaryTreeNode(1, Left(NewBinaryTreeNode(3)), Right(NewBinaryTreeNode(2)))
				tree2 := NewBinaryTree(root2)

				actual := tree1.LeafSimilar(tree2)

				Expect(actual).To(Equal(false))
			})
		})
	})

	Context("CountGoodNodes", func() {
		When("Binary Tree contains nodes with data of type int", func() {
			It("should return 0 for tree having no root", func() {
				tree := NewBinaryTree[int](nil)

				actual := tree.CountGoodNodes()
				Expect(actual).To(Equal(0))
			})

			It("should return 1 for tree having root, but no children", func() {
				root := NewBinaryTreeNode(1)
				tree := NewBinaryTree(root)

				actual := tree.CountGoodNodes()
				Expect(actual).To(Equal(1))
			})

			It("should return 4 for tree=(3,1,4,3,null,1,5)", func() {
				root := NewBinaryTreeNode(3, Left(NewBinaryTreeNode(1, Left(NewBinaryTreeNode(3)))),
					Right(NewBinaryTreeNode(4, Left(NewBinaryTreeNode(1)), Right(NewBinaryTreeNode(5)))))

				tree := NewBinaryTree(root)

				actual := tree.CountGoodNodes()

				Expect(actual).To(Equal(4))
			})

			It("should return 3 for tree=(3,3,null,4,2)", func() {
				root := NewBinaryTreeNode(3, Left(NewBinaryTreeNode(3, Right(NewBinaryTreeNode(2)), Left(NewBinaryTreeNode(4)))))
				tree := NewBinaryTree(root)

				actual := tree.CountGoodNodes()

				Expect(actual).To(Equal(3))
			})
		})
	})
})
