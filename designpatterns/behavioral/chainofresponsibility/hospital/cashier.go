package hospital

import "fmt"

// Cashier is a concrete handler
type Cashier struct {
	next Department
}

func (d *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
		return
	}

	fmt.Println("Cashier getting money from patient")
}

func (d *Cashier) setNext(next Department) {
	d.next = next
}
