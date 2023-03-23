package car

import "fmt"

type Car struct {
	wheels Wheels
	speed  Speed
	color  Color
}

func (c Car) Drive() error {
	fmt.Printf("Car is driving")
	return nil
}

func (c Car) Stop() error {
	fmt.Printf("Car is stopping")
	return nil
}
