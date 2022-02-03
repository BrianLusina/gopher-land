package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h, m int
}

func modulo(n, m int) int {
	return ((n % m) + m) % m
}

func New(h, m int) Clock {
	i := (h*60 + m)
	n := 24 * 60
	total := modulo(i, n)
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
