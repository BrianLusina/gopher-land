package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func HeavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	println("Done with heavy function 1")
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	println("Done with heavy function 2")
}

func run_using_wait_groups() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go HeavyFunction1(wg)
	go HeavyFunction2(wg)
	wg.Wait()
	fmt.Printf("All Finished!")
}
