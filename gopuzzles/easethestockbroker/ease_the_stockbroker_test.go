package easethestockbroker

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEaseTheStockBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EaseTheStockBroker Test Suite")
}

func dotest(lst string, exp string) {
	var ans = BalanceStatement(lst)
	Expect(ans).To(Equal(exp))
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}

func generateInputs(k int) string {
	ticker := []string{"GOOG", "APPL", "CSCO", "CAP", "CLH16.NYM", "ZYGN", "JPMC", "BoAML", "CITI", "MYSPACE"}
	qty := []string{"12", "24.0", "50", "100", "45", "67", "56", "78", "90", "15"}
	price := []string{"5.5", "89.3", "45.5", "12.8", "123.0", "210", "450", "34.8", "55.5", "160.45"}
	verb := []string{"B", "S", "B", "S", "B", "S", "B", "S", "P", "C"}
	res := ""
	shuffle(ticker)
	shuffle(qty)
	shuffle(price)
	shuffle(verb)
	for i := 0; i < k; i++ {
		res += fmt.Sprintf("%s %s %s %s, ", ticker[i], qty[i], price[i], verb[i])
	}
	return res[0 : len(res)-2]
}

var _ = Describe("Tests BalanceStatement", func() {

	It("should handle GOOG 300 542.0 B, AAPL 50 145.0 B, CSCO 250.0 29 B, GOOG 200 580.0 S", func() {
		input := "GOOG 300 542.0 B, AAPL 50 145.0 B, CSCO 250.0 29 B, GOOG 200 580.0 S"
		expected := "Buy: 169850 Sell: 116000; Badly formed 1: CSCO 250.0 29 B ;"
		actual := BalanceStatement(input)
		Expect(actual).To(Equal(expected))
	})

	It("should handle GOOG 90 160.45 B, JPMC 67 12.8 S, MYSPACE 24.0 210 B, CITI 50 450 B, CSCO 100 55.5 S", func() {
		input := "GOOG 90 160.45 B, JPMC 67 12.8 S, MYSPACE 24.0 210 B, CITI 50 450 B, CSCO 100 55.5 S"
		expected := "Buy: 14440 Sell: 6408; Badly formed 2: MYSPACE 24.0 210 B ;CITI 50 450 B ;"
		actual := BalanceStatement(input)
		Expect(actual).To(Equal(expected))
	})

	It("should handle NGA 1300 2.66 B, CLH15.NYM 50 56.32 B, OWW 1000 11.623 B, OGG 20 580.1 B", func() {
		input := "ZNGA 1300 2.66 B, CLH15.NYM 50 56.32 B, OWW 1000 11.623 B, OGG 20 580.1 B"
		expected := "Buy: 29499 Sell: 0"
		actual := BalanceStatement(input)
		Expect(actual).To(Equal(expected))
	})

	It("should handle empty order", func() {
		input := ""
		expected := "Buy: 0 Sell: 0"
		dotest(input, expected)
	})

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 40; i++ {
		k := random(2, 6)
		input := generateInputs(k)
		It(fmt.Sprintf("should handle %s", input), func() {
			expected := BalanceStatementV2(input)
			actual := BalanceStatement(input)
			Expect(actual).To(Equal(expected))
		})
	}
})
