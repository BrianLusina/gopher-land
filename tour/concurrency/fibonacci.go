package concurrency

import "fmt"

func fibonacci(limit int, channel chan int) {
	a, b := 0, 1
	for i := 0; i < limit; i++ {
		channel <- a
		a, b = b, a+b
	}
	close(channel)
}

func sum(a []int, channel chan int) {
	sum := 0
	for _, value := range a {
		sum += value
	}
	channel <- sum
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	channel := make(chan int)

	go sum(a[:len(a)/2], channel)
	go sum(a[len(a)/2:], channel)

	x, y := <-channel, <-channel

	fmt.Println(x, y, x+y)

	fibonacciChannel := make(chan int, 10)
	go fibonacci(cap(fibonacciChannel), fibonacciChannel)
	fmt.Println("fibonacci sequence")
	for x := range fibonacciChannel {
		fmt.Println(x)
	}
}

