package banksystem

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type BankAccount struct {
	accountId int
	balance   int64
}

func NewBankAccount(accountNumber int, balance int64) *BankAccount {
	return &BankAccount{accountNumber, balance}
}

func (account *BankAccount) GetAccountId() int {
	return account.accountId
}

func (account *BankAccount) GetBalance() int64 {
	return account.balance
}

func (account *BankAccount) Deposit(amount int64) {
	account.balance += amount
}

func (account *BankAccount) Withdraw(amount int64) error {
	if account.balance-amount < 0 {
		return ErrInsufficientFunds
	}
	account.balance -= amount
	return nil
}
