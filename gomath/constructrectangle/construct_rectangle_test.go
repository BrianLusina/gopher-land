package constructrectangle

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConstructRectangle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConstructRectangle Suite")
}

var _ = Describe("ConstructRectangle", func() {
	It("should return [2, 2] for area of 4", func() {
		area := 4
		expected := []int{2, 2}
		actual := constructRectangle(area)
		Expect(actual).To(Equal(expected))
	})

	It("should return [37, 1] for area of 37", func() {
		area := 37
		expected := []int{37, 1}
		actual := constructRectangle(area)
		Expect(actual).To(Equal(expected))
	})

	It("should return [427, 286] for area of 122122", func() {
		area := 122122
		expected := []int{427, 286}
		actual := constructRectangle(area)
		Expect(actual).To(Equal(expected))
	})
})
