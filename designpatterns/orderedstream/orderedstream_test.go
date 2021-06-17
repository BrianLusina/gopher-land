package orderedstream_test

import (
	"testing"

	"github.com/BrianLusina/goffer-land/designpatterns/orderedstream"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrderedStream(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OrderedStream Suite")
}

var _ = Describe("OrderedStream with n = 5", func() {
	orderedStream := orderedstream.Constructor(5)

	It("insert 3, 'ccccc'", func() {
		actual := orderedStream.Insert(3, "ccccc")
		Expect(actual).To(Equal([]string{}))
	})

	XIt("insert 1, 'aaaaa'", func() {
		actual := orderedStream.Insert(3, "aaaaa")
		Expect(actual).To(Equal([]string{"aaaaa"}))
	})

	XIt("insert 2, 'bbbbb'", func() {
		actual := orderedStream.Insert(2, "bbbbb")
		Expect(actual).To(Equal([]string{"bbbbb", "ccccc"}))
	})

	It("insert 5, 'eeeee'", func() {
		actual := orderedStream.Insert(5, "eeeee")
		Expect(actual).To(Equal([]string{}))
	})

	XIt("insert 4, 'ddddd'", func() {
		actual := orderedStream.Insert(4, "ddddd")
		Expect(actual).To(Equal([]string{"ddddd", "eeeee"}))
	})
})
