package goroutines

import "fmt"

func produce(c chan func()) {
	f := <-c
	f()
	c <- nil
}

func run_anonymous_closure() {
	c := make(chan func())
	go produce(c)
	x := 0
	c <- func() { x++ }
	fmt.Println(x)
}
