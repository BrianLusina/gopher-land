package types

import "fmt"

func main() {
	var a [2]string
	a[0] = "Yoh"
	a[1] = "Punk"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	// take a slice from index 1 to 5, taking the second element, but excluding
	// the 4th
	var s []int = primes[1:5]
	fmt.Println(s)
}
