package account

import "sync"

// Account defines a bank account.
type Account struct {
	amount int64
	isOpen bool
	mux    sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount, isOpen: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.amount, a.isOpen
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.isOpen {
		if (a.amount + amount) < 0 {
			return a.amount, false
		}
		a.amount += amount
		return a.amount, true
	}
	return a.amount, false
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.isOpen {
		return 0, false
	}
	a.isOpen = false
	amount := a.amount
	a.amount = 0
	return amount, true
}
