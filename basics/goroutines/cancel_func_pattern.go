package main

import "fmt"

func countToWithCancel(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan bool)
	cancel := func() {
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

func main() {
	ch, cancel := countToWithCancel(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
