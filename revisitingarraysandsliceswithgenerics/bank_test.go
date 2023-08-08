package main

import (
	"testing"
)

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

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}

func Find[A any](items []A, function func(A) bool) (value A, found bool) {
	for _, v := range items {
		if function(v) {
			return v, true
		}
	}
	return
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
