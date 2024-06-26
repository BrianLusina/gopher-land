package dynamicstack

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestLifoStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dynamic Stack Test Suite")
}

var _ = Describe("Dynamic Stack", func() {
	stack := New[int]()

	It("Should add -2 to stack", func() {
		err := stack.Push(-2)

		if err != nil {
			Expect(err).To(BeNil())
		}

		Expect(stack.Peek()).To(Equal(-2))
		Expect(stack.Items()).To(Equal([]int{-2}))
	})

	It("Should correctly check on its size", func() {
		Expect(stack.IsEmpty()).To(Equal(false))
		Expect(stack.Size()).To(Equal(1))
	})

	It("Should add -3 to stack", func() {
		_ = stack.Push(-3)

		Expect(stack.Peek()).To(Equal(-3))
		Expect(stack.Items()).To(Equal([]int{-2, -3}))
	})

	It("Should correctly check on its size after adding -3", func() {
		Expect(stack.IsEmpty()).To(Equal(false))
		Expect(stack.Size()).To(Equal(2))
	})

	It("Should add 0 to stack", func() {
		_ = stack.Push(0)

		Expect(stack.Peek()).To(Equal(0))
		Expect(stack.Items()).To(Equal([]int{-2, -3, 0}))
	})

	It("Should correctly check on its size after adding 0", func() {
		Expect(stack.IsEmpty()).To(Equal(false))
		Expect(stack.Size()).To(Equal(3))
	})

	It("Should peek and return 0 without removing item", func() {
		actual, err := stack.Peek()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(0))
		Expect(stack.Items()).To(Equal([]int{-2, -3, 0}))
		Expect(stack.Size()).To(Equal(3))
	})

	It("Should pop item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(0))
		Expect(stack.Items()).To(Equal([]int{-2, -3}))
		Expect(stack.Size()).To(Equal(2))
	})

	It("Should pop another item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(-3))
		Expect(stack.Items()).To(Equal([]int{-2}))
		Expect(stack.Size()).To(Equal(1))
	})

	It("Should pop yet another item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(-2))
		Expect(stack.Items()).To(Equal([]int{}))
		Expect(stack.Size()).To(Equal(0))
	})

	It("Should fail to peek item as its now empty", func() {
		actual, err := stack.Peek()

		assert.Error(GinkgoT(), err)
		Expect(actual).To(Equal(0))
	})
})
