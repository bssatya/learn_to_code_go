package main

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance += ammount
}

func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if w.balance < ammount {
		return errors.New("oh No")
	}

	w.balance -= ammount
	return nil
}

func (w *Wallet) Balance() (balance Bitcoin) {
	balance = w.balance
	return
}
