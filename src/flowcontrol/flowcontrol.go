package main

import (
	"fmt"
	"math"
)

// usage of if statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// usage of if statement with a short statement
func pow(x, y, lim float64) float64 {

	// variable in the statement is only in scope until the end of the for loop
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}


func main() {
	sum := 0

	for i:=0; i< 10; i++ {
		sum += 1
	}

	fmt.Println(sum)

	// init and post statements are optional
	var p = 1
	for ;p < 1000;  {
		p += 1000
	}

	println(p)

	println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
