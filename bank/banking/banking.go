package banking

import (
	"errors"
	"fmt"
)

// Account is Account information
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("충분한 돈이 없어요")

// NewAccount is return new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit is deposit to account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance is return account balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from Account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of Account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
