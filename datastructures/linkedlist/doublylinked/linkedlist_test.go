package doublylinkedlist

import (
	"fmt"
	"gopherland/datastructures/linkedlist"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDoublyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DoublyLinkedList Suite")
}

var _ = Describe("DoublyLinkedList", func() {

	Describe("Delete operations", func() {
		Context("when deleting by node data", func() {

			It("should successfully delete a node by its data", func() {
				dll := NewLinkedList[int]()

				dll.Append(1)
				dll.Append(2)
				dll.Append(3)
				dll.Append(4)
				dll.Append(5)

				Expect(dll.Size()).To(Equal(5))

				dll.DeleteNodeByData(3)

				Expect(dll.Size()).To(Equal(4))

				Expect(fmt.Sprintf("%v", dll)).To(Equal("1 <-> 2 <-> 4 <-> 5 <-> nil"))
			})
		})

	})

	Describe("Prepend operations", func() {

		Context("when adding a new node by it's data to a non empty linked list", func() {

			It("should successfully add a new node to the head", func() {
				dll := NewLinkedList[int]()

				dll.Append(1)
				dll.Append(2)
				dll.Append(3)
				dll.Append(4)
				dll.Append(5)

				Expect(dll.Size()).To(Equal(5))

				dll.Prepend(0)

				Expect(dll.Size()).To(Equal(6))

				Expect(fmt.Sprintf("%v", dll)).To(Equal("0 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <-> nil"))
			})
		})

		Context("when adding a new node by it's data to an empty linked list", func() {

			It("should successfully add a new node to the head when the node has no head", func() {
				dll := NewLinkedList[int]()

				Expect(dll.Size()).To(Equal(0))

				dll.Prepend(1)

				Expect(dll.Size()).To(Equal(1))
				Expect(dll.Head.Data).To(Equal(1))
			})
		})
	})

	Context("GetMiddleNode", func() {
		It("should return nil for an empty list", func() {
			dll := NewLinkedList[int]()
			val, err := dll.GetMiddleNode()

			Expect(err).ToNot(BeNil())
			Expect(val).To(BeNil())

			Expect(err).To(Equal(linkedlist.ErrEmptyList))
		})

		It("should return 3 for linked list of 1 -> 2 -> 3 -> 4 -> 5", func() {
			dll := NewLinkedList[int]()
			dll.Append(1)
			dll.Append(2)
			dll.Append(3)
			dll.Append(4)
			dll.Append(5)

			val, err := dll.GetMiddleNode()

			Expect(err).To(BeNil())
			Expect(val).ToNot(BeNil())
			Expect(val.Data).To(Equal(3))
		})

		It("should return 7 for linked list of 2->4->6->7->5->1", func() {
			dll := NewLinkedList[int]()
			dll.Append(2)
			dll.Append(4)
			dll.Append(6)
			dll.Append(7)
			dll.Append(5)
			dll.Append(1)

			val, err := dll.GetMiddleNode()

			Expect(err).To(BeNil())
			Expect(val).ToNot(BeNil())
			Expect(val.Data).To(Equal(7))
		})
	})

	Context("Reverse", func() {
		It("should reverse a non empty linked list 1 -> 2 -> 3 -> 4 -> 5", func() {
			sll := NewLinkedList[int]()
			sll.Append(1)
			sll.Append(2)
			sll.Append(3)
			sll.Append(4)
			sll.Append(5)

			sll.Reverse()

			Expect(sll.Head.Data).To(Equal(5))
		})

		It("should reverse linked list of 2->4->6->7->5->1", func() {
			sll := NewLinkedList[int]()
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

	Context("Rotate", func() {
		It("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 when k=0", func() {
			dll := NewLinkedList[int]()
			dll.Append(10)
			dll.Append(20)
			dll.Append(30)
			dll.Append(40)
			dll.Append(50)
			dll.Append(60)

			dll.Rotate(0)

			Expect(dll.Head.Data).To(Equal(10))
		})

		It("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=7 positions", func() {
			dll := NewLinkedList[int]()
			dll.Append(10)
			dll.Append(20)
			dll.Append(30)
			dll.Append(40)
			dll.Append(50)
			dll.Append(60)

			dll.Rotate(7)

			Expect(dll.Head.Data).To(Equal(10))
		})

		It("should rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=4 positions", func() {
			dll := NewLinkedList[int]()
			dll.Append(10)
			dll.Append(20)
			dll.Append(30)
			dll.Append(40)
			dll.Append(50)
			dll.Append(60)

			dll.Rotate(4)

			Expect(dll.Head.Data).To(Equal(50))
		})
	})
})
