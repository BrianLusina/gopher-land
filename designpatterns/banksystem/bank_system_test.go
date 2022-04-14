package banksystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithFiveAccounts(t *testing.T) {
	accountsWithBalances := []int64{10, 100, 20, 50, 30}
	bank := NewBank(accountsWithBalances)

	// withdraw from account 3 with initial balance of 20, current balance is 10
	actualOne := bank.Withdraw(3, 10)
	assert.True(t, actualOne)

	// transfer 20 from account 5(balance of 30) to account 1(balance of 10)
	actualTwo := bank.Transfer(5, 1, 20)
	assert.True(t, actualTwo)

	// deposit 20 to account 5 with current balance of 10, new balance is 30
	actualThree := bank.Deposit(5, 20)
	assert.True(t, actualThree)

	// transfer 15 from account 3(current balance 10) to account 4(current balance 50)
	actualFour := bank.Transfer(3, 4, 15)
	assert.False(t, actualFour)

	// withdraw 50 from account 10 should return false. There is no such account
	actualFive := bank.Withdraw(10, 50)
	assert.False(t, actualFive)
}
