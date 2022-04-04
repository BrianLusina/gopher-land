package printererrors

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestBinarySearchTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PrinterError Suite")
}

var _ = Describe("PrinterError", func() {
	It("should work on fixed tests", func() {
		Expect(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("3/56"))
		Expect(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")).To(Equal("6/60"))
		Expect(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu")).To(Equal("11/65"))
	})
})

var _ = Describe("Extensive Test", func() {
	myPrinterError := func(s string) string {
		err := 0
		for _, r := range s {
			if r < 'a' || r > 'm' {
				err++
			}
		}
		return strconv.Itoa(err) + "/" + strconv.Itoa(len(s))
	}

	randInt := func(min, max int) int {
		return min + rand.Intn(max-min) + 1
	}

	generateRandomStr := func(size int) string {
		rs := make([]rune, size)
		for i := range rs {
			rs[i] = rune(randInt('a', 'z'))
		}
		return string(rs[:])
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 200; i++ {
		It(fmt.Sprintf("%d random tests", i), func() {
			r := randInt(50, 500)
			str := generateRandomStr(r)
			Expect(PrinterError(str)).To(Equal(myPrinterError(str)))
		})
	}
})
