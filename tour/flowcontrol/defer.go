package main

import "fmt"

//  Defer
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func stackEmUp() {
	fmt.Println("Counting numbers and stars")

	for i := 0; i < 10; i++ {
		defer println(i)
	}

	fmt.Println("I'm done, but I will be executed first")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	stackEmUp()
}

