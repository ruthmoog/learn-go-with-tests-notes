package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Ruth",
			To:   "Chris",
			Sum:  1000000,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Ruth"), -1000000)
	AssertEqual(t, BalanceFor(transactions, "Chris"), 1000000)
}
