package customer

import (
	"fmt"
	"log"
	"testing"
)

func TestCustomer(t *testing.T) {
	c := customer{age: 33, name: "Jonny", data: &data{balance: 100.}}
	c.add(50.0)
	fmt.Printf("Balance: %.2f\n", c.data.balance)

	if err := c.validate(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}
}
