package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't widthdraw")

// NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// withdraw x mount from you account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// ChangeOwner of the accout
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "`s account.\nHas: ", a.Balance())
}
