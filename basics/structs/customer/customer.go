package customer

import (
	"errors"
	"gopherland/pkg/errdefs"
)

type customer struct {
	data *data
	age  int
	name string
}

type data struct {
	balance float64
}

func (c customer) add(v float64) {
	c.data.balance += v
}

func (c customer) validate() error {
	var m *errdefs.MultiError

	if c.age < 0 {
		m = &errdefs.MultiError{}
		m.Add(errors.New("age is negative"))
	}

	if c.name == "" {
		if m == nil {
			m = &errdefs.MultiError{}
		}

		m.Add(errors.New("name is nil"))
	}

	if m != nil {
		return m
	}

	return nil
}
