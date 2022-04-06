package primegap

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPrimeGap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrimeGap Suite")
}

var _ = Describe("PrimeGap", func() {

	It("should return nil for gap less than 2", func() {
		gap := 1
		start := 100
		end := 110
		actual := Gap(gap, start, end)
		Expect(actual).To(BeNil())
	})

	It("should return nil for start being greater than end", func() {
		gap := 2
		start := 300
		end := 110
		actual := Gap(gap, start, end)
		Expect(actual).To(BeNil())
	})

	It("should return (101, 103) for gap of 2 with start 100 and end 110", func() {
		gap := 2
		start := 100
		end := 110
		expected := []int{101, 103}
		actual := Gap(gap, start, end)
		Expect(actual).To(Equal(expected))
	})

	It("should return (103, 107) for gap of 4 with start 100 and end 110", func() {
		gap := 4
		start := 100
		end := 110
		expected := []int{103, 107}
		actual := Gap(gap, start, end)
		Expect(actual).To(Equal(expected))
	})

	It("should return nil for gap of 6 with start 100 and end 110", func() {
		gap := 6
		start := 100
		end := 110
		actual := Gap(gap, start, end)
		Expect(actual).To(BeNil())
	})

	It("should return (359, 367) for gap of 8 with start 300 and end 400", func() {
		gap := 8
		start := 300
		end := 400
		expected := []int{359, 367}
		actual := Gap(gap, start, end)
		Expect(actual).To(Equal(expected))
	})

	It("should return (337, 347) for gap of 10 with start 300 and end 400", func() {
		gap := 10
		start := 300
		end := 400
		expected := []int{337, 347}
		actual := Gap(gap, start, end)
		Expect(actual).To(Equal(expected))
	})
})
