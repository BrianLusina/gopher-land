package circleofnumbers

import (
	"math/rand"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("CircleOfNumbers(10, 2)", func() { Expect(CircleOfNumbers(10, 2)).To(Equal(7)) })
	It("CircleOfNumbers(10, 7)", func() { Expect(CircleOfNumbers(10, 7)).To(Equal(2)) })
	It("CircleOfNumbers(4, 1)", func() { Expect(CircleOfNumbers(4, 1)).To(Equal(3)) })
	It("CircleOfNumbers(6, 3)", func() { Expect(CircleOfNumbers(6, 3)).To(Equal(0)) })
	It("CircleOfNumbers(20, 0)", func() { Expect(CircleOfNumbers(20, 0)).To(Equal(10)) })
})

func Sol(n int, firstNumber int) int {
	return (firstNumber + n/2) % n
}

var _ = Describe("Random Tests", func() {
	rand.Seed(time.Now().UnixNano())
	for t := 0; t < 100; t++ {
		n := (rand.Intn(499) + 2) * 2
		f := rand.Intn(n)
		It("CircleOfNumbers("+strconv.Itoa(n)+", "+strconv.Itoa(f)+")", func() {
			Expect(CircleOfNumbers(n, f)).To(Equal(Sol(n, f)))
		})
	}
})
