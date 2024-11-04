package goroutines

import "fmt"

func run_random_ones_zeroes() {
	ch := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-ch, " ")
		}
	}()

	// producer:
	for i := 0; i <= 100000; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		}
	}
}
