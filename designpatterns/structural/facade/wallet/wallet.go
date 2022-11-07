package wallet

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return new(Wallet)
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
}

func (w *Wallet) debitBalance(amount int) error {
	if amount > w.balance {
		return fmt.Errorf("Balance is insufficient")
	}
	w.balance -= amount
	return nil
}
