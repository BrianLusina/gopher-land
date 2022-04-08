package main

func processChannel(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

func main() {
	receiveOnlyCh := make(<-chan int)
	sendOnlyCh := make(chan<- int)
	go processChannel(receiveOnlyCh, sendOnlyCh)
}
