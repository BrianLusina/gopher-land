package stack

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestLifoStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LifoStack Test Suite")
}

var _ = Describe("LifoStack", func() {
	stack := NewLifoStack[int](10)

	It("Should add -2 to stack", func() {
		err := stack.Push(-2)

		if err != nil {
			Expect(err).To(BeNil())
		}

		Expect(stack.Peek()).To(Equal(-2))
		Expect(stack.items).To(Equal([]interface{}{-2}))
	})

	It("Should correctly check on its size", func() {
		Expect(stack.isEmpty()).To(Equal(false))
		Expect(stack.isFull()).To(Equal(false))
		Expect(stack.size()).To(Equal(1))
	})

	It("Should add -3 to stack", func() {
		_ = stack.Push(-3)

		Expect(stack.Peek()).To(Equal(-3))
		Expect(stack.items).To(Equal([]interface{}{-2, -3}))
	})

	It("Should correctly check on its size after adding -3", func() {
		Expect(stack.isEmpty()).To(Equal(false))
		Expect(stack.isFull()).To(Equal(false))
		Expect(stack.size()).To(Equal(2))
	})

	It("Should add 0 to stack", func() {
		_ = stack.Push(0)

		Expect(stack.Peek()).To(Equal(0))
		Expect(stack.items).To(Equal([]interface{}{-2, -3, 0}))
	})

	It("Should correctly check on its size after adding 0", func() {
		Expect(stack.isEmpty()).To(Equal(false))
		Expect(stack.isFull()).To(Equal(false))
		Expect(stack.size()).To(Equal(3))
	})

	It("Should peek and return 0 without removing item", func() {
		actual, err := stack.Peek()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(0))
		Expect(stack.items).To(Equal([]interface{}{-2, -3, 0}))
		Expect(stack.size()).To(Equal(3))
	})

	It("Should pop item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(0))
		Expect(stack.items).To(Equal([]interface{}{-2, -3}))
		Expect(stack.size()).To(Equal(2))
	})

	It("Should pop another item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(-3))
		Expect(stack.items).To(Equal([]interface{}{-2}))
		Expect(stack.size()).To(Equal(1))
	})

	It("Should pop yet another item & remove it & return it reducing size", func() {
		actual, err := stack.Pop()

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(-2))
		Expect(stack.items).To(Equal([]interface{}{}))
		Expect(stack.size()).To(Equal(0))
	})

	It("Should fail to peek item as its now empty", func() {
		actual, err := stack.Peek()

		assert.Error(GinkgoT(), err)
		Expect(actual).To(BeNil())
	})
})
