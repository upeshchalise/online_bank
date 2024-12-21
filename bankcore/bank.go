package bank

import (
	"errors"
	"fmt"
)

// func Hello() string {
// 	return "Hey! I am working"
// }

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {

	if amount < 0 {
		return errors.New("withdraw amount should be greater than zero")
	}

	if amount > a.Balance {
		return errors.New("no enough balance")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(amount float64, destination *Account) error {

	if amount < 0 {
		return errors.New("deposited amount should be greater than zero")
	}

	if amount > a.Balance {
		return errors.New("no enough balance")
	}

	a.Withdraw(amount)
	// destination.Balance += amount
	destination.Deposit(amount)
	return nil
}

type Bank interface {
	Statement() string
}

func Statement(b Bank) string {
	return b.Statement()
}
