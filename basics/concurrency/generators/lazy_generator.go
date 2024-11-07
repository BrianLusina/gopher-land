package main

import "fmt"

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInts() int {
	return <-resume
}

func lazyGenerator() {
	resume = integers()
	fmt.Println(generateInts()) // 0
	fmt.Println(generateInts()) // 1
	fmt.Println(generateInts()) // 2
	fmt.Println(generateInts()) // 3
}
