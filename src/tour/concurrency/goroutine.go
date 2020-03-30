package concurrency

import (
	"fmt"
	"time"
)

// Welcome message prints a simple welcome message
func Welcome() {
	fmt.Println("Welcome welcome Bots and Droids!")
}

// main runs a goroutine
func main() {
	// run a different goroutine
	go Welcome()

	// run another goroutine
	go func() {
		fmt.Println("Hello Universe!")
	}()

	// we will wait until goroutines finish execution
	time.Sleep(time.Millisecond*200)
}