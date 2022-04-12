package main

import (
	"fmt"
)

// sendData sends data to the channel
func sendDataSync(ch chan<- string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

// getData receives data from the channel
func getDataSync(ch <-chan string) {
	for {
		input, isOpen := <-ch
		if !isOpen {
			break
		}
		fmt.Printf("%s\n ", input)
	}
}

func main() {
	ch := make(chan string)
	go sendDataSync(ch)
	getDataSync(ch)
}
