package concurrency

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1 + <-right }

func run_chaining_goroutines() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // start the chaining
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, approx. 1.5 s
}
