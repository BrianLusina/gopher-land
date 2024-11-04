package goroutines

import "fmt"

// running this will cause fatal error: all goroutines are asleep - deadlock!
func run_deadlock_goroutine() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Println(v, v2)
}
