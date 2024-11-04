package goroutines

func processChannel(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

func run_pipe_filter() {
	receiveOnlyCh := make(<-chan int)
	sendOnlyCh := make(chan<- int)
	go processChannel(receiveOnlyCh, sendOnlyCh)
}
