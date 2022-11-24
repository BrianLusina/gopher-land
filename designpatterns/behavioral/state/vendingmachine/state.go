package vendingmachine

// State is an interface that all concrete states implement
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
