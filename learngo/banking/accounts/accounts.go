package accounts

import (
	"encoding/json"
	"errors"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var ErrNoMoney = errors.New("can't withdraw")

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Desposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return ErrNoMoney
	}

	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	jsonData, err := json.Marshal(struct {
		Owner   string
		Balance int
	}{
		Owner:   a.owner,
		Balance: a.balance,
	})

	if err != nil {
		return err.Error()
	}

	return string(jsonData)
}
