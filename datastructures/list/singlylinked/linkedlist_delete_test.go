package singlylinkedlist

import (
	"gopherland/datastructures/list"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSinglyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SinglyLinkedList Delete Suite")
}

var _ = Describe("SinglyLinkedList Delete", func() {

	Context("Delete Node By Position", func() {

		When("linked list is [1,3,4,7,1,2,6], delete node at position 0", func() {
			values := []int{1, 3, 4, 7, 1, 2, 6}
			linkedList := New[int]()
			for _, v := range values {
				linkedList.Append(v)
			}

			expected := list.NewNode(1)

			It("should return 1 as the deleted node and delete it from linked list", func() {
				actual, actualErr := linkedList.DeleteNodeByPosition(0)
				assert.NoError(GinkgoT(), actualErr)
				assert.Equal(GinkgoT(), expected.Data, actual.Data)
			})
		})

		When("linked list is [1,2,3,4], delete node at position 3", func() {
			values := []int{1, 2, 3, 4}
			linkedList := New[int]()
			for _, v := range values {
				linkedList.Append(v)
			}

			expected := list.NewNode(4)

			It("should return 3 as the deleted node and delete it from linked list", func() {
				actual, actualErr := linkedList.DeleteNodeByPosition(3)
				assert.NoError(GinkgoT(), actualErr)
				assert.Equal(GinkgoT(), expected.Data, actual.Data)
			})
		})

		When("linked list is [2, 1], delete node at position 2", func() {
			values := []int{2, 1}
			linkedList := New[int]()
			for _, v := range values {
				linkedList.Append(v)
			}

			It("should return 1 as the middle node and delete it from linked list", func() {
				actual, actualErr := linkedList.DeleteNodeByPosition(2)
				assert.Error(GinkgoT(), actualErr)
				assert.Nil(GinkgoT(), actual)
			})
		})

	})
})
