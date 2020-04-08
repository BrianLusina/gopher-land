package funcs

import (
	"fmt"
)

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubstrExpr   = MathExpr("substract")
	MultiplyExpr = MathExpr("multiply")
)

func main() {
	func() {
		println("Hello")
	}()

	a := func(name string) {
		fmt.Println("Hello %s", name)
	}

	a("You")
}

func exp(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return Add
	case SubstrExpr:
		return Divide
	}
}
