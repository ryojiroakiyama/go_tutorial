package bank

import (
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if amount > a.Balance {
		return errors.New("the amount to withdraw greater than balance")
	}

	a.Balance -= amount
	return nil
}

// Send ...
func (a *Account) Send(amount float64, remittee *Account) error {
	err := a.Withdraw(amount)
	if err != nil {
		return fmt.Errorf("fail to withdraw from remittance source: %v", err)
	}
	err = remittee.Deposit(amount)
	if err != nil {
		a.Balance += amount
		return fmt.Errorf("fail to deposit to remittee: %v", err)
	}
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
