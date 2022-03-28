package stack

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMinStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MinStack Suite")
}

var _ = Describe("MinStack", func() {

	Describe("with values 0, -2, -3", func() {

		minStack := NewMinStack(10)

		It("Should add -2 to stack", func() {
			err := minStack.Push(-2)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(-2))
		})

		It("Should add 0 to stack", func() {
			err := minStack.Push(0)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(0))
		})

		It("Should add -3 to stack", func() {
			err := minStack.Push(-3)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(-3))
		})

		It("Should getMin & return -3 without removing it from the stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(-3))
		})

		It("Should pop() top item from stack", func() {
			_, err := minStack.Pop()

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(0))
		})

		It("Should peek() top item from stack without removing it", func() {
			actual, err := minStack.Peek()

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(0))
		})

		It("Should getMin() from current stack & return -2 without removing it from Stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(-2))
		})
	})

	Describe("with values 0, 2, 0, 3", func() {

		minStack := NewMinStack(10)

		It("Should add 2 to stack", func() {
			err := minStack.Push(2)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(2))
		})

		It("Should add 0 to stack", func() {
			err := minStack.Push(0)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(0))
		})

		It("Should add 3 to stack", func() {
			err := minStack.Push(3)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(3))
		})

		It("Should add another 0 to stack", func() {
			err := minStack.Push(0)

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(0))
		})

		It("Should getMin & return 0 without removing it from the stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(0))
		})

		It("Should pop() top item from stack", func() {
			_, err := minStack.Pop()

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(3))
		})

		It("Should getMin() from current stack & return 0 without removing it from Stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(0))
		})

		It("Should pop() top item from stack", func() {
			_, err := minStack.Pop()

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(0))
		})

		It("Should getMin() from current stack & return 0 without removing it from Stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(0))
		})

		It("Should pop() top item from stack", func() {
			_, err := minStack.Pop()

			Expect(err).To(BeNil())
			Expect(minStack.Peek()).To(Equal(2))
		})

		It("Should getMin() from current stack & return 2 without removing it from Stack", func() {
			actual := minStack.GetMin()

			Expect(actual).To(Equal(2))
		})
	})
})
