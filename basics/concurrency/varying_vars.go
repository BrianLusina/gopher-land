package concurrency

import "fmt"

func run_varying_vars() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(val int) {
			ch <- val * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
