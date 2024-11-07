package synccond

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	cond    *sync.Cond
	balance int
}

func run_sync_cond() {
	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	listener := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go listener(10)
	go listener(15)

	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}
