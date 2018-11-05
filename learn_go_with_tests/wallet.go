package main

type Bitcoin int

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
