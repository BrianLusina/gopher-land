package indexedcapitalization

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(st string, arr []int, exp string) {
	var ans = Capitalize(st, arr)
	Expect(ans).To(Equal(exp))
}

func uniq(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func contain_s(arr []int, e int) bool {
	for _, n := range arr {
		if n == e {
			return true
		}
	}
	return false
}

func ci90(st string, arr []int) string {
	var res string
	for i, ch := range st {
		if contain_s(arr, i) {
			res += strings.ToUpper(string(ch))
		} else {
			res += string(ch)
		}
	}
	return res
}

func TestIndexedCapitalization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Indexed Capitaliztion Test Suite")
}

var _ = Describe("Indexed Capitalization tests", func() {
	It("It should work for basic tests", func() {
		dotest("abcdef", []int{1, 2, 5}, "aBCdeF")
		dotest("abcdef", []int{1, 2, 5, 100}, "aBCdeF")
		dotest("codewars", []int{1, 3, 5, 50}, "cOdEwArs")
		dotest("abracadabra", []int{2, 6, 9, 10}, "abRacaDabRA")
		dotest("codewarriors", []int{5}, "codewArriors")
		dotest("indexinglessons", []int{0}, "Indexinglessons")
	})

	It("It Should Work For Random Cases", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		alph := "abcdefghijklmnopqrstuvwxyz"
		for i := 0; i < 100; i++ {
			st := ""
			size := 10 + rand.Intn(20)
			for j := 0; j <= size; j++ {
				g := rand.Intn(26)
				st += string(alph[g])
			}
			arrLen := 1 + rand.Intn(size*2)
			arr := make([]int, arrLen)
			for i := range arr {
				arr[i] = rand.Intn(100)
			}
			arr = uniq(arr)
			sort.Ints(arr)
			exp := ci90(st, arr)
			dotest(st, arr, exp)
		}
	})
})
