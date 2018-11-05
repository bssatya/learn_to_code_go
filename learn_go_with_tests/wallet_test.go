package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func (t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got (%s), want (%s)", got, want)
		}
	})
	
	t.Run("Withdraw", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Got (%s), Want (%s)", got, want)
		}

	})

}
