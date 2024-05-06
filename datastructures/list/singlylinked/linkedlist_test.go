package singlylinkedlist

import (
	"fmt"
	"gopherland/datastructures/list"
	"gopherland/pkg/utils"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("Is Palindrome", func(t *testing.T) {
		type testCase[T any] struct {
			data     []T
			expected bool
		}

		var testCases = []testCase[any]{
			{
				data:     []any{"r", "a", "c", "e", "c", "a", "r"},
				expected: true,
			},
		}

		t.Run("using a stack", func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("should return %v for %v", tc.expected, tc.data), func(t *testing.T) {
					linkedList := New[any]()
					for _, d := range tc.data {
						linkedList.Append(d)
					}
					actual := linkedList.IsPalindrome()
					if actual != tc.expected {
						t.Fail()
						t.Fatalf("expected isPalindrome() for data=%v to be %v, got %v", tc.data, tc.expected, actual)
					}
				})
			}
		})

		t.Run("using two pointers", func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("should return %v for %v", tc.expected, tc.data), func(t *testing.T) {
					linkedList := New[any]()
					for _, d := range tc.data {
						linkedList.Append(d)
					}
					actual := linkedList.IsPalindromeTwoPointers()
					if actual != tc.expected {
						t.Fail()
						t.Fatalf("expected IsPalindromeTwoPointers() for data=%v to be %v, got %v", tc.data, tc.expected, actual)
					}
				})
			}
		})
	})
}

func BenchmarkLinkedList(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	b.Run("Is Palindrome", func(b *testing.B) {
		type testCase[T any] struct {
			data     []T
			expected bool
		}

		var testCases = []testCase[any]{
			{
				data:     []any{"r", "a", "c", "e", "c", "a", "r"},
				expected: true,
			},
		}

		b.Run("using a stack", func(b *testing.B) {
			for _, tc := range testCases {
				b.Run(fmt.Sprintf("should return %v for %v", tc.expected, tc.data), func(b *testing.B) {
					linkedList := New[any]()
					for _, d := range tc.data {
						linkedList.Append(d)
					}
					actual := linkedList.IsPalindrome()
					if actual != tc.expected {
						b.Fail()
						b.Fatalf("expected isPalindrome() for data=%v to be %v, got %v", tc.data, tc.expected, actual)
					}
				})
			}
		})

		b.Run("using two pointers", func(b *testing.B) {
			for _, tc := range testCases {
				b.Run(fmt.Sprintf("should return %v for %v", tc.expected, tc.data), func(b *testing.B) {
					linkedList := New[any]()
					for _, d := range tc.data {
						linkedList.Append(d)
					}
					actual := linkedList.IsPalindromeTwoPointers()
					if actual != tc.expected {
						b.Fail()
						b.Fatalf("expected isPalindrome() for data=%v to be %v, got %v", tc.data, tc.expected, actual)
					}
				})
			}
		})
	})
}

func TestSinglyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SinglyLinkedList Suite")
}

var _ = Describe("SinglyLinkedList", func() {
	t := GinkgoT()
	Context("Prepend ->Append -> AddAtPosition -> GetNthNode -> DeleteAtPosition -> GetNthNode operations", func() {

		It("should handle \"Prepend(1)\" -> \"Append(3)\" -> \"AddAtPosition(1,2)\" -> \"GetNthNode(1)\" -> \"DeleteAtPosition(1)\" -> \"GetNthNode(1)\"", func() {
			sll := New[int]()
			sll.Prepend(1)
			sll.Append(3)
			// linked list: 1 -> 3

			sll.AddAtPosition(1, 2)
			// linked list is now: 1 -> 2 -> 3

			actual2, _ := sll.GetNthNode(1)
			Expect(actual2.Data).To(Equal(2))

			node, err := sll.DeleteAtPosition(1)
			Expect(err).To(BeNil())
			Expect(node.Data).To(Equal(2))
			// linked list is now: 1 -> 3

			actual3, _ := sll.GetNthNode(1)
			Expect(actual3.Data).To(Equal(3))
		})

		It("Prepend(7) -> Prepend(2) -> Prepend(1) -> AddAtPosition(3, 0) -> DeleteAtPosition(2) -> Prepend(6) -> Append(4) -> GetNtNode(4) -> Prepend(4) -> AddAtPosition(5, 0) -> Prepend(6) ", func() {
			sll := New[int]()

			sll.Prepend(7)
			Expect(sll.Head.Data).To(Equal(7))
			// linked list: 7 -> nil

			sll.Prepend(2)
			Expect(sll.Head.Data).To(Equal(2))
			// linked list: 2 -> 7 -> nil

			sll.Prepend(1)
			Expect(sll.Head.Data).To(Equal(1))
			// linked list: 1 -> 2 -> 7 -> nil

			sll.AddAtPosition(3, 0)
			Expect(sll.Length()).To(Equal(4))
			// linked list: 1 -> 2 -> 7 -> 0 -> nil

			node, err := sll.DeleteAtPosition(2)
			Expect(err).To(BeNil())
			Expect(node.Data).To(Equal(7))
			// linked list: 1 -> 2 -> 0 -> nil

			sll.Prepend(6)
			Expect(sll.Head.Data).To(Equal(6))
			// linked list: 6 -> 1 -> 2 -> 0 -> nil

			sll.Append(4)
			// linked list: 6 -> 1 -> 2 -> 0 -> 4 -> nil

			actual4, _ := sll.GetNthNode(4)
			Expect(actual4.Data).To(Equal(4))

			sll.Prepend(4)
			Expect(sll.Head.Data).To(Equal(4))
			// linked list: 4 -> 6 -> 1 -> 2 -> 0 -> 4 -> nil

			sll.AddAtPosition(5, 0)
			// linked list: 4 -> 6 -> 1 -> 2 -> 0 -> 5 -> 4 -> nil

			sll.Prepend(6)
			Expect(sll.Head.Data).To(Equal(6))
			// linked list: 6 -> 4 -> 6 -> 1 -> 2 -> 0 -> 5 -> 4 -> nil
		})

		It("Prepend(7) -> Append(7) -> Prepend(9) -> Append(8) -> Prepend(6) -> Prepend(0) -> GetNthNode(5)-> Prepend(0) -> GetNthNode(2) -> GetNthNode(5) -> Append(4)", func() {
			sll := New[int]()

			sll.Prepend(7)
			Expect(sll.Head.Data).To(Equal(7))
			// linked list: 7 -> nil

			sll.Append(7)
			Expect(sll.Length()).To(Equal(2))
			// linked list: 7 -> 7 -> nil

			sll.Prepend(9)
			Expect(sll.Head.Data).To(Equal(9))
			Expect(sll.Length()).To(Equal(3))
			// linked list: 9 -> 7 -> 7 -> nil

			sll.Append(8)
			Expect(sll.Length()).To(Equal(4))
			// linked list: 9 -> 7 -> 7 -> 8 -> nil

			sll.Prepend(6)
			Expect(sll.Head.Data).To(Equal(6))
			Expect(sll.Length()).To(Equal(5))
			// linked list: 6 -> 9 -> 7 -> 7 -> 8 -> nil

			sll.Prepend(0)
			Expect(sll.Head.Data).To(Equal(0))
			Expect(sll.Length()).To(Equal(6))
			// linked list: 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> nil

			actual5, _ := sll.GetNthNode(5)
			Expect(actual5.Data).To(Equal(8))
			// linked list: 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> nil

			sll.Prepend(0)
			Expect(sll.Head.Data).To(Equal(0))
			Expect(sll.Length()).To(Equal(7))
			// linked list: 0 -> 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> nil

			actual2, _ := sll.GetNthNode(2)
			Expect(actual2.Data).To(Equal(6))
			// linked list: 0 -> 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> nil

			actual5, _ = sll.GetNthNode(5)
			Expect(actual5.Data).To(Equal(7))
			// linked list: 0 -> 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> nil

			sll.Append(4)
			Expect(sll.Length()).To(Equal(8))
			// linked list: 0 -> 0 -> 6 -> 9 -> 7 -> 7 -> 8 -> 4 -> nil
		})

		It("Prepend(2) -> DeleteAtPosition(1) -> Prepend(2) -> Prepend(7) -> Prepend(3) -> Prepend(2) -> Prepend(5)-> Append(5) -> GetNthNode(5) -> DeleteAtPosition(6) -> DeleteAtPosition(4)", func() {
			sll := New[int]()

			sll.Prepend(2)
			Expect(sll.Head.Data).To(Equal(2))
			// linked list: 2 -> nil

			nodeAtPosition1, err := sll.DeleteAtPosition(1)
			Expect(nodeAtPosition1).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(sll.Length()).To(Equal(1))
			// linked list: 2 -> nil

			sll.Prepend(2)
			Expect(sll.Head.Data).To(Equal(2))
			Expect(sll.Length()).To(Equal(2))
			// linked list: 2 -> 2 -> nil

			sll.Prepend(7)
			Expect(sll.Head.Data).To(Equal(7))
			Expect(sll.Length()).To(Equal(3))
			// linked list: 7 -> 2 -> 2 -> nil

			sll.Prepend(3)
			Expect(sll.Head.Data).To(Equal(3))
			Expect(sll.Length()).To(Equal(4))
			// linked list: 3 -> 7 -> 2 -> 2 -> nil

			sll.Prepend(2)
			Expect(sll.Head.Data).To(Equal(2))
			Expect(sll.Length()).To(Equal(5))
			// linked list: 2 -> 3 -> 7 -> 2 -> 2 -> nil

			sll.Prepend(5)
			Expect(sll.Head.Data).To(Equal(5))
			Expect(sll.Length()).To(Equal(6))
			// linked list: 5 -> 2 -> 3 -> 7 -> 2 -> 2 -> nil

			sll.Append(5)
			Expect(sll.Length()).To(Equal(7))
			// linked list: 5 -> 2 -> 3 -> 7 -> 2 -> 2 -> 5 -> nil

			actual5, _ := sll.GetNthNode(5)
			Expect(actual5.Data).To(Equal(2))
			// linked list: 5 -> 2 -> 3 -> 7 -> 2 -> 2 -> 5 -> nil

			nodeAtPosition6, err := sll.DeleteAtPosition(6)
			Expect(err).To(BeNil())
			Expect(nodeAtPosition6.Data).To(Equal(5))
			Expect(sll.Length()).To(Equal(6))
			// linked list: 5 -> 2 -> 3 -> 7 -> 2 -> 2 -> nil

			nodeAtPosition4, err := sll.DeleteAtPosition(4)
			Expect(err).To(BeNil())
			Expect(nodeAtPosition4.Data).To(Equal(2))
			Expect(sll.Length()).To(Equal(5))
			// linked list: 5 -> 2 -> 3 -> 7 -> 2 -> nil
		})

		It("Prepend(4) -> GetNthNode(1) -> Prepend(1) -> Prepend(5) -> DeleteAtPosition(3) -> Prepend(7) -> GetNthNode(3)-> GetNthNode(3) -> GetNthNode(3) -> Prepend(1) -> DeleteAtPosition(4)", func() {
			sll := New[int]()

			sll.Prepend(4)
			Expect(sll.Head.Data).To(Equal(4))
			// linked list: 4 -> nil

			nodeAtPosition1, err := sll.GetNthNode(1)
			Expect(nodeAtPosition1).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(sll.Length()).To(Equal(1))
			// linked list: 2 -> nil

			sll.Prepend(1)
			Expect(sll.Head.Data).To(Equal(1))
			Expect(sll.Length()).To(Equal(2))
			// linked list: 1 -> 4 -> nil

			sll.Prepend(5)
			Expect(sll.Head.Data).To(Equal(5))
			Expect(sll.Length()).To(Equal(3))
			// linked list: 5 -> 1 -> 4 -> nil

			nodeAtPosition3, err := sll.DeleteAtPosition(3)
			Expect(nodeAtPosition3).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(sll.Length()).To(Equal(3))
			// linked list: 5 -> 1 -> 4 -> nil

			sll.Prepend(7)
			Expect(sll.Head.Data).To(Equal(7))
			Expect(sll.Length()).To(Equal(4))
			// linked list: 7 -> 5 -> 1 -> 4 -> nil

			actual3, _ := sll.GetNthNode(3)
			Expect(actual3.Data).To(Equal(4))
			// linked list: 7 -> 5 -> 1 -> 4 -> nil

			actual3, _ = sll.GetNthNode(3)
			Expect(actual3.Data).To(Equal(4))
			// linked list: 7 -> 5 -> 1 -> 4 -> nil

			actual3, _ = sll.GetNthNode(3)
			Expect(actual3.Data).To(Equal(4))
			// linked list: 7 -> 5 -> 1 -> 4 -> nil

			sll.Prepend(1)
			Expect(sll.Head.Data).To(Equal(1))
			Expect(sll.Length()).To(Equal(5))
			// linked list: 1 -> 7 -> 5 -> 1 -> 4 -> nil

			nodeAtPosition4, err := sll.DeleteAtPosition(4)
			Expect(nodeAtPosition4).ToNot(BeNil())
			Expect(nodeAtPosition4.Data).To(Equal(4))
			Expect(err).To(BeNil())
			Expect(sll.Length()).To(Equal(4))
			// linked list: 1 -> 7 -> 5 -> 1 -> nil
		})

		It("Prepend(5) -> AddAtPosition(1, 2) -> GetNthNode(1) -> Prepend(6) -> Append(2) -> GetNthNode(3)-> Append(1) -> GetNthNode(5) -> Prepend(2) -> GetNthNode(2) -> Prepend(6)", func() {
			sll := New[int]()

			sll.Prepend(5)
			Expect(sll.Head.Data).To(Equal(5))
			Expect(sll.Length()).To(Equal(1))
			// linked list: 5 -> nil

			sll.AddAtPosition(1, 2)
			Expect(sll.Length()).To(Equal(2))
			// linked list: 5 -> 2 -> nil

			firstNode, err1 := sll.GetNthNode(1)
			Expect(err1).To(BeNil())
			Expect(firstNode.Data).To(Equal(2))
			Expect(sll.Head.Data).To(Equal(5))
			Expect(sll.Length()).To(Equal(2))
			// linked list: 5 -> 2 -> nil

			sll.Prepend(6)
			Expect(sll.Head.Data).To(Equal(6))
			Expect(sll.Length()).To(Equal(3))
			// linked list: 6 -> 5 -> 2 -> nil

			sll.Append(2)
			Expect(sll.Length()).To(Equal(4))
			// linked list: 6 -> 5 -> 2 -> 2 -> nil

			thirdNode, err2 := sll.GetNthNode(3)
			Expect(err2).To(BeNil())
			Expect(thirdNode.Data).To(Equal(2))
			Expect(sll.Length()).To(Equal(4))
			// linked list: 6 -> 5 -> 2 -> 2 -> nil

			sll.Append(1)
			Expect(sll.Length()).To(Equal(5))
			// linked list: 6 -> 5 -> 2 -> 2 -> 1 -> nil

			fifthNode, err3 := sll.GetNthNode(5)
			Expect(err3).ToNot(BeNil())
			Expect(fifthNode).To(BeNil())
			// linked list: 6 -> 5 -> 2 -> 2 -> 1 -> nil

			sll.Prepend(2)
			Expect(sll.Head.Data).To(Equal(2))
			Expect(sll.Length()).To(Equal(6))
			// linked list: 2 -> 6 -> 5 -> 2 -> 2 -> 1 -> nil

			secondNode, err4 := sll.GetNthNode(2)
			Expect(secondNode.Data).To(Equal(5))
			Expect(err4).To(BeNil())
			Expect(sll.Length()).To(Equal(6))
			// linked list: 2 -> 6 -> 5 -> 2 -> 2 -> 1 -> nil

			sll.Prepend(6)
			Expect(sll.Head.Data).To(Equal(6))
			Expect(sll.Length()).To(Equal(7))
			// linked list: 6 -> 2 -> 6 -> 5 -> 2 -> 2 -> 1 -> nil
		})
	})

	Context("GetMiddleNode", func() {
		It("should return nil for an empty list", func() {
			sll := New[int]()
			val, err := sll.GetMiddleNode()

			Expect(err).ToNot(BeNil())
			Expect(val).To(BeNil())

			Expect(err).To(Equal(list.ErrEmptyList))
		})

		It("should return 3 for linked list of 1 -> 2 -> 3 -> 4 -> 5", func() {
			sll := New[int]()
			sll.Append(1)
			sll.Append(2)
			sll.Append(3)
			sll.Append(4)
			sll.Append(5)

			val, err := sll.GetMiddleNode()

			Expect(err).To(BeNil())
			Expect(val).ToNot(BeNil())
			Expect(val.Data).To(Equal(3))
		})

		It("should return 7 for linked list of 2->4->6->7->5->1", func() {
			sll := New[int]()
			sll.Append(2)
			sll.Append(4)
			sll.Append(6)
			sll.Append(7)
			sll.Append(5)
			sll.Append(1)

			val, err := sll.GetMiddleNode()

			Expect(err).To(BeNil())
			Expect(val).ToNot(BeNil())
			Expect(val.Data).To(Equal(7))
		})
	})

	Context("Reverse", func() {
		It("should reverse a non empty linked list 1 -> 2 -> 3 -> 4 -> 5", func() {
			sll := New[int]()
			sll.Append(1)
			sll.Append(2)
			sll.Append(3)
			sll.Append(4)
			sll.Append(5)

			sll.Reverse()

			Expect(sll.Head.Data).To(Equal(5))
		})

		It("should reverse linked list of 2->4->6->7->5->1", func() {
			sll := New[int]()
			sll.Append(2)
			sll.Append(4)
			sll.Append(6)
			sll.Append(7)
			sll.Append(5)
			sll.Append(1)

			sll.Reverse()

			Expect(sll.Head.Data).To(Equal(1))
		})
	})

	Context("KthToLastNode", func() {
		When("LinkedList is a->b->c->d", func() {
			sll := New[string]()
			sll.Append("a")
			sll.Append("b")
			sll.Append("c")
			sll.Append("d")

			It("should return d for k = 1", func() {
				expected := list.NewNode("d")
				actual, err := sll.KthToLastNode(1)
				assert.NoError(t, err)

				Expect(actual).To(Equal(expected))
			})

			It("should return c for k = 2", func() {
				expected := list.NewNode("c")
				actual, err := sll.KthToLastNode(2)
				assert.NoError(t, err)

				Expect(actual.Data).To(Equal(expected.Data))
			})

			It("should return b for k = 3", func() {
				expected := list.NewNode("b")
				actual, err := sll.KthToLastNode(3)
				assert.NoError(t, err)

				Expect(actual.Data).To(Equal(expected.Data))
			})

			It("should return a for k = 4", func() {
				expected := list.NewNode("a")
				actual, err := sll.KthToLastNode(4)
				assert.NoError(t, err)

				Expect(actual.Data).To(Equal(expected.Data))
			})

			It("should return an error for k = 5", func() {
				actual, err := sll.KthToLastNode(5)
				assert.Error(t, err)
				assert.Nil(t, actual)
			})
		})
	})

	Context("DeleteMiddleNode", func() {

		Context("Using 2 Passes", func() {
			When("linked list is [1,3,4,7,1,2,6]", func() {
				values := []int{1, 3, 4, 7, 1, 2, 6}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(7)

				It("should return 7 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})

			When("linked list is [1,2,3,4]", func() {
				values := []int{1, 2, 3, 4}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(3)

				It("should return 3 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})

			When("linked list is [2, 1]", func() {
				values := []int{2, 1}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(1)

				It("should return 1 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})
		})

		Context("Using 2 Pointers", func() {
			When("linked list is [1,3,4,7,1,2,6]", func() {
				values := []int{1, 3, 4, 7, 1, 2, 6}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(7)

				It("should return 7 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle2Pointers()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})

			When("linked list is [1,2,3,4]", func() {
				values := []int{1, 2, 3, 4}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(3)

				It("should return 3 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle2Pointers()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})

			When("linked list is [2, 1]", func() {
				values := []int{2, 1}
				linkedList := New[int]()
				for _, v := range values {
					linkedList.Append(v)
				}

				expected := list.NewNode(1)

				It("should return 1 as the middle node and delete it from linked list", func() {
					actual := linkedList.DeleteMiddle2Pointers()

					assert.Equal(t, expected.Data, actual.Data)
				})
			})
		})
	})

	Context("OddEventList", func() {
		runTest := func(data, expected []int) {
			linkedList := New[int]()
			for _, v := range data {
				linkedList.Append(v)
			}

			actualHead := linkedList.OddEvenList()
			actualNodes := []int{}

			for actualHead != nil {
				actualNodes = append(actualNodes, actualHead.Data)
				actualHead = actualHead.Next
			}

			zipped, err := utils.Zip[int, int](actualNodes, expected)
			assert.NoError(t, err)

			for _, zipPair := range zipped {
				assert.Equal(t, zipPair.A, zipPair.B)
			}
		}

		When("linked list is [1, 2, 3, 4, 5]", func() {
			data := []int{1, 2, 3, 4, 5}
			expected := []int{1, 3, 5, 2, 4}

			It("should return [1, 3, 5, 2, 4] as the new list", func() {
				runTest(data, expected)
			})
		})

		When("linked list is [2, 1, 3, 5, 6, 4, 7]", func() {
			data := []int{2, 1, 3, 5, 6, 4, 7}
			expected := []int{2, 3, 6, 7, 1, 5, 4}

			It("should return [2, 3, 6, 7, 1, 5, 4] as the new list", func() {
				runTest(data, expected)
			})
		})
	})

	Context("Swap", func() {
		type testCase[T comparable] struct {
			data     []T
			expected []T
			nodeOne  T
			nodeTwo  T
			runTest  func(data, expected []T, nodeOne, nodeTwo T)
		}

		testCases := []testCase[string]{
			{
				data:     []string{"A", "B", "C", "D"},
				expected: []string{"A", "C", "B", "D"},
				nodeOne:  "B",
				nodeTwo:  "C",
				runTest: func(data, expected []string, nodeOne, nodeTwo string) {
					linkedList := New[string]()
					for _, v := range data {
						linkedList.Append(v)
					}

					linkedList.SwapNodes(nodeOne, nodeTwo)
					actualNodes := []string{}
					actualHead := linkedList.Head

					for actualHead != nil {
						actualNodes = append(actualNodes, actualHead.Data)
						actualHead = actualHead.Next
					}

					zipped, err := utils.Zip(actualNodes, expected)
					assert.NoError(t, err)

					for _, zipPair := range zipped {
						assert.Equal(t, zipPair.A, zipPair.B)
					}
				},
			},
		}

		for _, tc := range testCases {
			When(fmt.Sprintf("linked list is %v", tc.data), func() {
				It(fmt.Sprintf("should return %v as the new list", tc.expected), func() {
					tc.runTest(tc.data, tc.expected, tc.nodeOne, tc.nodeTwo)
				})
			})
		}
	})

	Context("Remove Duplicates", func() {
		type testCase[T comparable] struct {
			data     []T
			expected []T
			runTest  func(data, expected []T)
		}

		testCases := []testCase[string]{
			{
				data:     []string{"A", "B", "C", "D", "A"},
				expected: []string{"A", "B", "C", "D"},
				runTest: func(data, expected []string) {
					linkedList := New[string]()
					for _, v := range data {
						linkedList.Append(v)
					}

					actualHead := linkedList.RemoveDuplicates()
					actualNodes := []string{}

					for actualHead != nil {
						actualNodes = append(actualNodes, actualHead.Data)
						actualHead = actualHead.Next
					}

					assert.ElementsMatch(t, expected, actualNodes)
				},
			},
		}

		for _, tc := range testCases {
			When(fmt.Sprintf("linked list is %v", tc.data), func() {
				It(fmt.Sprintf("should return %v as the new list", tc.expected), func() {
					tc.runTest(tc.data, tc.expected)
				})
			})
		}
	})

})
