package hospital

import "fmt"

// Medical is a concrete handler
type Medical struct {
	next Department
}

func (d *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		d.next.execute((p))
		return
	}

	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	d.next.execute(p)
}

func (d *Medical) setNext(next Department) {
	d.next = next
}
