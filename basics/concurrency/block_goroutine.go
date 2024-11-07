// Illustrates where the goroutine pump sends integers in an infinite loop on the channel. However, because there is
// no receiver, the only output is the number 0
package concurrency

import (
	"fmt"
	"time"
)

func run_block_goroutines() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
	time.Sleep(1e9)
}

// suck only receives integers from the channel ch
func suck(ch <-chan int) {
	// use a for-range loop to read from a channel
	for v := range ch {
		fmt.Println(v)
	}
}

// pump sends integers to the channel ch
func pump(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
