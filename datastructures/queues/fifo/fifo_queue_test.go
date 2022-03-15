package fifo

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Test Suite")
}

var _ = Describe("Queue", func() {
	queue := NewQueue()

	It("Should add 1 to queue", func() {
		queue.Enqueue(1)

		Expect(queue.Peek()).To(Equal(1))
		Expect(queue.Items()).To(Equal([]interface{}{1}))
	})

	It("Should add 2 to the queue", func() {
		queue.Enqueue(2)

		Expect(queue.Peek()).To(Equal(1))
		Expect(queue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should peek front element in the queue without removing it", func() {
		Expect(queue.Peek()).To(Equal(1))
		Expect(queue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should remove element from the front of the Queue", func() {
		Expect(queue.Dequeue()).To(Equal(1))
		Expect(queue.Items()).To(Equal([]interface{}{2}))
	})

	It("Should return false for empty() for the current queue", func() {
		Expect(queue.Empty()).To(Equal(false))
	})

})
