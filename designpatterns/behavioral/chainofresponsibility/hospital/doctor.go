package hospital

import "fmt"

// Doctor is a concrete handler
type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckupDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute((p))
		return
	}

	fmt.Println("Doctor checking patient")
	p.doctorCheckupDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
