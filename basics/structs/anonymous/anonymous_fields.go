package main

import "fmt"

type inner struct {
	a int
	b int
}

type outer struct {
	c int
	d float32

	int
	inner
}

func main() {
	outer1 := new(outer)
	outer1.c = 1
	outer1.d = 5.6
	outer1.int = 4
	outer1.a = 2
	outer1.b = 3

	fmt.Printf("outer1.c is: %d\n", outer1.c)
	fmt.Printf("outer1.d is: %f\n", outer1.d)
	fmt.Printf("outer1.int is: %d\n", outer1.int)
	fmt.Printf("outer1.a is: %d\n", outer1.a)
	fmt.Printf("outer1.b is: %d\n", outer1.b)

	// with a struct-literal:
	outer2 := outer{6, 7.5, 60, inner{5, 10}}
	fmt.Println("outer2 is: ", outer2)
}
