package charcodecalc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func referenceSolution(s string) (res int) {
	for _, r := range s {
		for n := int(r); n != 0; n /= 10 {
			if n%10 == 7 {
				res += 6
			}
		}
	}
	return
}

func TestCharCodeCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Char Code Calculation Suite")
}

var _ = Describe("Char Code Calculation", func() {
	It("input of abcdef should return 6", func() {
		input := "abcdef"
		expected := 6
		actual := Calc(input)

		Expect(actual).To(Equal(expected))
	})

	It("input of ifkhchlhfd should return 6", func() {
		input := "ifkhchlhfd"
		expected := 6
		actual := Calc(input)

		Expect(actual).To(Equal(expected))
	})

	It("input of aaaaaddddr should return 30", func() {
		input := "aaaaaddddr"
		expected := 30
		actual := Calc(input)

		Expect(actual).To(Equal(expected))
	})

	It("input of jfmgklf8hglbe should return 6", func() {
		input := "jfmgklf8hglbe"
		expected := 6
		actual := Calc(input)

		Expect(actual).To(Equal(expected))
	})

	It("input of jaam should return 12", func() {
		input := "jaam"
		expected := 12
		actual := Calc(input)

		Expect(actual).To(Equal(expected))
	})

	for i := 0; i < 100; i++ {
		It(fmt.Sprintf("Random test %d", i), func() {
			arr := make([]rune, 1+rand.Intn(8))
			for j := range arr {
				arr[j] = rune(32 + rand.Intn(95))
			}
			input := string(arr)
			expected := referenceSolution(input)
			actual := Calc(input)

			Expect(actual).To(Equal(expected))
		})
	}
})
