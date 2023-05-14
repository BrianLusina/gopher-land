package circular

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCircularQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Circular Queue Test Suite with slice")
}

var _ = Describe("CircularQueue", func() {
	Context("Using a slice as the underlying implementation", func() {

		When("handling data of type int on a circular queue of size 5", func() {
			queue := NewQueue[int](5)

			It("Should add 1 to queue", func() {
				queue.Enqueue(1)

				Expect(queue.Peek()).To(Equal(1))
				Expect(queue.Size()).To(Equal(1))
			})

			It("Should add 2 to the queue", func() {
				queue.Enqueue(2)

				Expect(queue.Peek()).To(Equal(1))
				Expect(queue.Size()).To(Equal(2))
			})

			It("Should peek front element in the queue without removing it", func() {
				Expect(queue.Peek()).To(Equal(1))
				Expect(queue.Size()).To(Equal(2))
			})

			It("Should remove element from the front of the Queue", func() {
				Expect(queue.Dequeue()).To(Equal(1))
				Expect(queue.Size()).To(Equal(1))
			})

			It("Should return false for empty() for the current queue", func() {
				Expect(queue.IsEmpty()).To(Equal(false))
			})

		})
	})

	Context("Using a linked list as the underlying implementation", func() {})

})
