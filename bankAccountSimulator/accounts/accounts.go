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

var errNotEnoughCash = errors.New("not enough cash")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (receiver *Account) Deposit(amount int) {
	fmt.Println("Amount of Deposit :", amount)
	receiver.balance += amount
}

// Balance of your account
func (receiver Account) Balance() int {
	return receiver.balance
}

// Withdraw x amount of your account
func (receiver *Account) Withdraw(amount int) error {
	if receiver.balance < amount { 
		return errNotEnoughCash
	}
	receiver.balance -= amount
	return nil
}

// ChangeOwner of the account
func (receiver *Account) ChangeOwner(newOwner string) {
	receiver.owner = newOwner
}

func (receiver Account) Owner() string {
	return receiver.owner
}