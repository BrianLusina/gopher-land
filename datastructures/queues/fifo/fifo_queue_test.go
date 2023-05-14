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
	When("int queue is of max size of 5", func() {
		fifoQueue := NewFifoQueue[int](5)

		It("Should add 1 to fifoQueue", func() {
			fifoQueue.Enqueue(1)

			Expect(fifoQueue.Peek()).To(Equal(1))
			Expect(fifoQueue.Size()).To(Equal(1))
		})

		It("Should add 2 to the fifoQueue", func() {
			fifoQueue.Enqueue(2)

			Expect(fifoQueue.Peek()).To(Equal(1))
			Expect(fifoQueue.Size()).To(Equal(2))
		})

		It("Should peek front element in the fifoQueue without removing it", func() {
			Expect(fifoQueue.Peek()).To(Equal(1))
		})

		It("Should remove element from the front of the Queue", func() {
			Expect(fifoQueue.Dequeue()).To(Equal(1))
			Expect(fifoQueue.Size()).To(Equal(1))
		})

		It("Should return false for IsEmpty() for the current fifoQueue", func() {
			Expect(fifoQueue.IsEmpty()).To(Equal(false))
		})

		It("Should return size of 0 for the current fifoQueue after dequeuing last item", func() {
			Expect(fifoQueue.Dequeue()).To(Equal(2))
			Expect(fifoQueue.Size()).To(Equal(0))
		})
	})
})
