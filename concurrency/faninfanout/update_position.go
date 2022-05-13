package faninfanout

import (
	"fmt"
	"math/rand"
	"time"
)

func updatePosition(name string) <-chan string {
	positionCh := make(chan string)

	go func() {
		for i := 0; ; i++ {
			positionCh <- fmt.Sprintf("%s: %d", name, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return positionCh
}

func fanIn(chan1, chan2 <-chan string) <-chan string {
	outCh := make(chan string)

	go func() {
		for {
			outCh <- <-chan1
		}
	}()

	go func() {
		for {
			outCh <- <-chan2
		}
	}()

	// or below
	// go func(){
	// 	for {
	// 		select {
	// 		case s := <-chan1:
	// 			outCh <- s
	// 		case s := <-chan2:
	// 			outCh <- s
	// 		}
	// 	}
	// }()

	return outCh
}

// func main() {
// 	positionsChannel := fanIn(updatePosition("Legolas :"), updatePosition("Gandalf :"))

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-positionsChannel)
// 	}

// 	fmt.Println("Done with getting updates on positions.")
// }
