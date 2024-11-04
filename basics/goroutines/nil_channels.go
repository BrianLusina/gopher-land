package goroutines

// merge demonstrates the merging of two channels into a single channel using the concept of nil channels
func merge(ch1, ch2 chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v1, isOpen := <-ch1:
				if !isOpen {
					ch1 = nil
					break
				}
				ch <- v1
			case v2, isOpen := <-ch2:
				if !isOpen {
					ch2 = nil
					break
				}
				ch <- v2
			}
		}
		close(ch)
	}()

	return ch
}
