package factories

import "fmt"

type filter func(int) bool
type sliceSplit func([]int) ([]int, []int)

// isOdd checks if a number is odd
func isOdd(integer int) bool {
	return integer%2 != 0
}

// greaterThanN checks if an integer is greater than n
func isGreaterThan(integer, n int) bool {
	return integer > n
}

// filterFactory takes in a filter function and returns a function of type sliceSplit
func filterFactory(f filter) sliceSplit {
	return func(s []int) ([]int, []int) {
		yes, no := []int{}, []int{}
		for _, n := range s {
			if f(n) {
				yes = append(yes, n)
			} else {
				no = append(no, n)
			}
		}
		return yes, no
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)
	odd_even_function := filterFactory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)
	//separate those that are bigger than 4 and those that are not.
	// bigger, smaller := filterFactory(isGreaterThan(4))(s)
	// fmt.Println("Bigger than 4: ", bigger)
	// fmt.Println("Smaller than 4: ", smaller)
}
