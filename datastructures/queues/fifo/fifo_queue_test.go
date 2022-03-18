package fifo

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFifoQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FifoQueue Test Suite")
}

var _ = Describe("FifoQueue", func() {
	fifoQueue := NewFifoQueue()

	It("Should add 1 to fifoQueue", func() {
		fifoQueue.Enqueue(1)

		Expect(fifoQueue.Peek()).To(Equal(1))
		Expect(fifoQueue.items).To(Equal([]interface{}{1}))
	})

	It("Should add 2 to the fifoQueue", func() {
		fifoQueue.Enqueue(2)

		Expect(fifoQueue.Peek()).To(Equal(1))
		Expect(fifoQueue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should peek front element in the fifoQueue without removing it", func() {
		Expect(fifoQueue.Peek()).To(Equal(1))
		Expect(fifoQueue.Items()).To(Equal([]interface{}{1, 2}))
	})

	It("Should remove element from the front of the Queue", func() {
		Expect(fifoQueue.Dequeue()).To(Equal(1))
		Expect(fifoQueue.Items()).To(Equal([]interface{}{2}))
	})

	It("Should return false for empty() for the current fifoQueue", func() {
		Expect(fifoQueue.isEmpty()).To(Equal(false))
	})
})
