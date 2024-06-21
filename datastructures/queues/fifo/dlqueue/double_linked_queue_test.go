package dlqueue

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinkedListQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedListQueue Test Suite")
}

var _ = Describe("LinkedListQueue", func() {
	linkedListQueue := NewFifoDLLQueue[int]()

	It("Should add 1 to linkedListQueue", func() {
		linkedListQueue.Enqueue(1)

		Expect(linkedListQueue.Peek()).To(Equal(1))
		Expect(linkedListQueue.Items()).To(Equal([]interface{}{1}))
	})

	It("Should add 2 to the linkedListQueue", func() {
		linkedListQueue.Enqueue(2)

		Expect(linkedListQueue.Peek()).To(Equal(1))
		Expect(linkedListQueue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should peek front element in the linkedListQueue without removing it", func() {
		Expect(linkedListQueue.Peek()).To(Equal(1))
		Expect(linkedListQueue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should remove element from the front of the Queue", func() {
		Expect(linkedListQueue.Dequeue()).To(Equal(1))
		Expect(linkedListQueue.Items()).To(Equal([]interface{}{2}))
	})

	It("Should return false for empty() for the current linkedListQueue", func() {
		Expect(linkedListQueue.isEmpty()).To(Equal(false))
	})
})
