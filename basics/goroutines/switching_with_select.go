package main

import (
	"fmt"
	"runtime"
	"time"
)

func pumpDoubles(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pumpPlus5(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suckAll(ch1 <-chan int, ch2 <-chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received %d from channel 1 \n", v)
		case v := <-ch2:
			fmt.Printf("Received %d from channel 2 \n", v)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pumpDoubles(ch1)
	go pumpPlus5(ch2)
	suckAll(ch1, ch2)
	time.Sleep(1e9)
}
