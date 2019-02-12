package main

import (
	"fmt"
)

// the type of variable comes last
var c, python, kotlin bool

// variables with initializers with the type declared
var i, j int = 1, 2

func main() {
	// declaring variables with the type inferred
	var js, cpp, java = true, false, "no!"

	var x int
	fmt.Println(x, c, python, kotlin)

	fmt.Println(i, j, js, cpp, java)

	// short variable declarations,
	// only available within functions
	k := 9
	fmt.Println(k)
}
