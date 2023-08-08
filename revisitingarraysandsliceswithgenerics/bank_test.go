package main

import "testing"

func TestBadBank(t *testing.T) {
	var (
		ruth   = Account{Name: "Ruth", Balance: 0}
		chris  = Account{Name: "Chris", Balance: 3000}
		pepper = Account{Name: "Pepper", Balance: 4}

		transactions = []Transaction{
			NewTransaction(ruth, chris, 100),
			NewTransaction(chris, pepper, 300),
		}
	)

	newBalanceFor := func(account Account) int {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(ruth), -100)
	AssertEqual(t, newBalanceFor(chris), 2800)
	AssertEqual(t, newBalanceFor(pepper), 304)
}
