package goroutines

import (
	"fmt"
	"time"
)

func run_channel_ticker() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(" .")
			time.Sleep(5e7)
		}
	}
}
