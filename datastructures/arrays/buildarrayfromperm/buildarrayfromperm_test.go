package buildarrayfromperm

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuildArrayFromPerm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuildArrayFromPerm Suite")
}

var _ = Describe("BuildArrayFromPerm", func() {

	It("test one with 0, 2, 1, 5, 3, 4", func() {
		nums := []int{0, 2, 1, 5, 3, 4}
		expected := []int{0, 1, 2, 4, 5, 3}
		actual := buildArray(nums)

		Expect(actual).To(Equal(expected))
	})

	It("test two with 5, 0, 1, 2, 3, 4", func() {
		nums := []int{5, 0, 1, 2, 3, 4}
		expected := []int{4, 5, 0, 1, 2, 3}
		actual := buildArray(nums)

		Expect(actual).To(Equal(expected))
	})
})
