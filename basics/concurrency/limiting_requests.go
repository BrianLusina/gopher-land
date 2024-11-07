package concurrency

const (
	AvailableMemory         = 10 << 20 // 10 MB
	AverageMemoryPerRequest = 10 << 10 // 10 KB
	MaxRequests             = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MaxRequests)

type MyRequest struct {
	a, b      int
	replyChan chan int
}

func process(r *MyRequest) {
	// Do something
	// may take a long time and use up memory & CPU
	r.replyChan <- r.a + r.b
}

func handle(r *MyRequest) {
	process(r)
	<-sem // signal done: enable next request to start by making 1 empty place in buffer
}

func MyServer(queue chan *MyRequest) {
	for {
		sem <- 1 // blocks when channel is full (1000 requests are active

		// so wait here until there is capacity to process a request
		// doesn't matter what we put in it
		request := <-queue
		go handle(request)
	}
}

func run_limiting_requests() {
	queue := make(chan *MyRequest)
	go MyServer(queue)
}
