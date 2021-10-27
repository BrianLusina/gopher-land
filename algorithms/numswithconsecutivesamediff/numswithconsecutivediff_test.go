package numswithconsecutivesamediff

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNumsWithConsecutiveDiff(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NumsWithConsecutiveDiff Suite")
}

var _ = Describe("NumsWithConsecutivDiff", func() {
	It("should return [181,292,707,818,929] for n =3 and k = 7", func() {
		n := 3
		k := 7
		expected := []int{181, 292, 707, 818, 929}
		actual := numsSameConsecDiff(n, k)
		Expect(actual).Should(Equal(expected))
	})

	It("should return [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98] for n = 2 and k = 1", func() {
		n := 2
		k := 1
		expected := []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98}
		actual := numsSameConsecDiff(n, k)
		Expect(actual).Should(Equal(expected))
	})

	It("should return [11,22,33,44,55,66,77,88,99] for n = 2 and k = 0", func() {
		n := 2
		k := 0
		expected := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
		actual := numsSameConsecDiff(n, k)
		Expect(actual).Should(Equal(expected))
	})

	It("should return [13,20,24,31,35,42,46,53,57,64,68,75,79,86,97] for n = 2 and k = 2", func() {
		n := 2
		k := 2
		expected := []int{13, 20, 24, 31, 35, 42, 46, 53, 57, 64, 68, 75, 79, 86, 97}
		actual := numsSameConsecDiff(n, k)
		Expect(actual).Should(Equal(expected))
	})
})
