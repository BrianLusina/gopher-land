package clock

import (
	"fmt"
	"gopherland/gomath/utils"
)

// Define the Clock type here.
type Clock struct {
	h, m int
}

func New(h, m int) Clock {
	i := (h*60 + m)
	n := 24 * 60
	total := utils.Modulo(i, n)
	hour := total / 60
	minute := total % 60
	return Clock{h: hour, m: minute}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
