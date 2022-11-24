package vendingmachine

import "fmt"

// NoItemState is a concrete state, indicating there are no items in the VendingMachine
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
