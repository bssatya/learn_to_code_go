package main

import (
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

func (w *Wallet) Balance() (balance Bitcoin) {
	balance = w.balance
	return
}
