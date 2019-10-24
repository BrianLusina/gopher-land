package src

import "fmt"

/**
 * function declaring types
 */
func add(x int, y int) int {
	return x + y
}

// function declaring same type for parameters
func subtract(a, b int) int {
	return a - b
}

// function returning multiple values
func swap(m, n string) (string, string) {
	return n, m
}

// has named return values
func split(sum int) (x, y int)  {
	x = sum * 4/ 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(1, 4))
	fmt.Println(subtract(6, 8))

	a, b := swap("Hello", "Hooman")

	fmt.Println(a, b)

	fmt.Println(split(17))
}
