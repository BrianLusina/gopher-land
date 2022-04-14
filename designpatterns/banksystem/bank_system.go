package banksystem

// Bank represents a bank
type Bank struct {
	// accounts is a mapping of account ids to BankAccount
	accounts map[int]*BankAccount
}

// NewBank creates a new Bank given accounts with balances
func NewBank(accountsWithBalance []int64) Bank {
	accounts := make(map[int]*BankAccount)
	for accountId, balance := range accountsWithBalance {
		if balance < 0 {
			panic("Negative balance is not allowed")
		}
		accounts[accountId+1] = NewBankAccount(accountId, balance)
	}
	return Bank{accounts: accounts}
}

func (bank Bank) getAccount(accountId int) *BankAccount {
	return bank.accounts[accountId]
}

func (bank Bank) accountExists(accountId int) bool {
	account, ok := bank.accounts[accountId]
	return ok && account != nil
}

// Transfer transfers money from one account to another
func (bank *Bank) Transfer(fromAccountId int, toAccountId int, money int64) bool {
	if !bank.accountExists(fromAccountId) || !bank.accountExists(toAccountId) {
		return false
	}

	fromAccount := bank.getAccount(fromAccountId)
	toAccount := bank.getAccount(toAccountId)

	if fromAccount.GetBalance() < money {
		return false
	}

	err := fromAccount.Withdraw(money)
	if err != nil {
		return false
	}

	toAccount.Deposit(money)
	return true
}

// Deposit deposits money into an account given an accountId. Returns false if account is not found
func (bank *Bank) Deposit(accountId int, money int64) bool {
	if bank.accountExists(accountId) {
		account := bank.getAccount(accountId)

		account.Deposit(money)
		return true
	}
	return false
}

func (bank *Bank) Withdraw(accountId int, money int64) bool {
	if bank.accountExists(accountId) {
		account := bank.getAccount(accountId)

		err := account.Withdraw(money)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
