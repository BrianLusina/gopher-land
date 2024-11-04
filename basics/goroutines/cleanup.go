package goroutines

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func run_cleanup() {
	for i := range countTo(10) {
		fmt.Println(i)
	}
}
