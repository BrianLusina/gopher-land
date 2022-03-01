// simply used to demonstrate how to use the timer package to time function
package timer

import (
	"fmt"
	"time"
)

func calculation() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func timeMe() {
	start := time.Now()
	calculation()
	end := time.Now()
	fmt.Println("Time taken: ", end.Sub(start))
}
