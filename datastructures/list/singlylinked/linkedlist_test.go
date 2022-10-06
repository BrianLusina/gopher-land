package singlylinkedlist

import (
	"gopherland/datastructures/linkedlist"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSinglyLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SinglyLinkedList Suite")
}

var _ = Describe("SinglyLinkedList", func() {
	It("should handle \"Prepend(1)\" -> \"Append(3)\" -> \"AddAtPosition(1,2)\" -> \"GetNthNode(1)\" -> \"DeleteAtPosition(1)\" -> \"GetNthNode(1)\"", func() {
		sll := NewLinkedList[int]()
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
		sll := NewLinkedList[int]()

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
		sll := NewLinkedList[int]()

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
		sll := NewLinkedList[int]()

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
		sll := NewLinkedList[int]()

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
		sll := NewLinkedList[int]()

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

	Context("GetMiddleNode", func() {
		It("should return nil for an empty list", func() {
			sll := NewLinkedList[int]()
			val, err := sll.GetMiddleNode()

			Expect(err).ToNot(BeNil())
			Expect(val).To(BeNil())

			Expect(err).To(Equal(linkedlist.ErrEmptyList))
		})

		It("should return 3 for linked list of 1 -> 2 -> 3 -> 4 -> 5", func() {
			sll := NewLinkedList[int]()
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
			sll := NewLinkedList[int]()
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
			sll := NewLinkedList[int]()
			sll.Append(10)
			sll.Append(20)
			sll.Append(30)
			sll.Append(40)
			sll.Append(50)
			sll.Append(60)

			sll.Rotate(0)

			Expect(sll.Head.Data).To(Equal(10))
		})

		It("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=7 positions", func() {
			sll := NewLinkedList[int]()
			sll.Append(10)
			sll.Append(20)
			sll.Append(30)
			sll.Append(40)
			sll.Append(50)
			sll.Append(60)

			sll.Rotate(7)

			Expect(sll.Head.Data).To(Equal(10))
		})

		It("should rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=4 positions", func() {
			sll := NewLinkedList[int]()
			sll.Append(10)
			sll.Append(20)
			sll.Append(30)
			sll.Append(40)
			sll.Append(50)
			sll.Append(60)

			sll.Rotate(4)

			Expect(sll.Head.Data).To(Equal(50))
		})
	})
})
