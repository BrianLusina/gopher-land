package stack

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaxStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxStack Suite")
}

var _ = Describe("MinStack", func() {
	maxStack := NewMaxStack(10)

	It("Should add -2 to stack", func() {
		maxStack.Push(-2)

		Expect(maxStack.Peek()).To(Equal(-2))
	})

	It("Should add 0 to stack", func() {
		maxStack.Push(0)

		Expect(maxStack.Peek()).To(Equal(0))
	})

	It("Should add -3 to stack", func() {
		maxStack.Push(-3)

		Expect(maxStack.Peek()).To(Equal(-3))
	})

	It("Should getMax & return 0 without removing it from the stack", func() {
		actual := maxStack.GetMax()

		Expect(actual).To(Equal(0))
	})

	It("Should pop() top item from stack", func() {
		actual := maxStack.Pop()

		Expect(actual).To(Equal(-3))
		Expect(maxStack.Peek()).To(Equal(0))
	})

	It("Should peek() top item from stack without removing it", func() {
		actual := maxStack.Peek()

		Expect(actual).To(Equal(0))
	})

	It("Should getMax() from current stack & return -2 without removing it from Stack", func() {
		actual := maxStack.GetMax()

		Expect(actual).To(Equal(0))
	})
})
