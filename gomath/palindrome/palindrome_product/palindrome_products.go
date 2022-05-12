package palindrome

import (
	"fmt"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax, %d > %d", fmin, fmax)
	}

	pmin := Product{}
	pmax := Product{}

	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			p := x * y
			if !isPalindrome(p) {
				continue
			}
			compare := func(current *Product, better bool) {
				switch {
				case current.Factorizations == nil || better:
					*current = Product{p, [][2]int{{x, y}}}
				case p == current.Product:
					current.Factorizations =
						append(current.Factorizations, [2]int{x, y})
				}
			}
			compare(&pmin, p < pmin.Product)
			compare(&pmax, p > pmax.Product)
		}
	}

	if len(pmin.Factorizations) == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes in range [%d, %d]", fmin, fmax)
	}

	return pmin, pmax, nil
}

func isPalindrome(x int) bool {
	var result string
	s := strconv.Itoa(x)
	for _, v := range s {
		result = string(v) + result
	}
	return result == s
}
