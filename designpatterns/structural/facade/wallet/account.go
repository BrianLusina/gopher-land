package wallet

import "fmt"

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (acc *Account) checkAccount(accountName string) error {
	if acc.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	return nil
}
