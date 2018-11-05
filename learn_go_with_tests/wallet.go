package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(ammount int) {
	w.balance += ammount
}

func (w *Wallet) Balance() (balance int) {
	balance = w.balance
	return
}
