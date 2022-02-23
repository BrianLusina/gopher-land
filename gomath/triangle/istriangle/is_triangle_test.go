package istriangle

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsTriangle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsTriangle Test Suite")
}

var _ = Describe("IsTriagnel Tests", func() {
	It("works for some examples", func() {
		Expect(IsTriangle(1, 2, 2)).To(Equal(true))
		Expect(IsTriangle(7, 2, 2)).To(Equal(false))
		Expect(IsTriangle(1, 2, 3)).To(Equal(false))
		Expect(IsTriangle(3, 1, 2)).To(Equal(false))
		Expect(IsTriangle(5, 1, 2)).To(Equal(false))
		Expect(IsTriangle(1, 2, 5)).To(Equal(false))
		Expect(IsTriangle(2, 5, 1)).To(Equal(false))
		Expect(IsTriangle(4, 2, 3)).To(Equal(true))
		Expect(IsTriangle(5, 1, 5)).To(Equal(true))
		Expect(IsTriangle(2, 2, 2)).To(Equal(true))
		Expect(IsTriangle(-1, 2, 3)).To(Equal(false))
		Expect(IsTriangle(1, -2, 3)).To(Equal(false))
		Expect(IsTriangle(1, 2, -3)).To(Equal(false))
		Expect(IsTriangle(0, 2, 3.)).To(Equal(false))
	})
	It("works for random examples", func() {
		for i := 0; i < 40; i++ {
			a, b, c := rand.Intn(12)-2, rand.Intn(12)-2, rand.Intn(12)-2
			Expect(IsTriangle(a, b, c)).To(Equal(a+b > c && b+c > a && c+a > b))
		}
	})
})
