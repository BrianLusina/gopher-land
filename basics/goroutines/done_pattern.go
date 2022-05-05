package goroutines

type searcher func(string) []string

func searchData(s string, searchers []searcher) []string {
	done := make(chan bool)
	result := make(chan []string)

	for _, search := range searchers {
		go func(seek searcher) {
			select {
			case result <- seek(s):
			case <-done:
			}
		}(search)
	}
	r := <-result
	close(done)
	return r
}
