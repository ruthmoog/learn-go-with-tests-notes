package pointersanderrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("Wanted %d but got %d", want, got)
	}
}

func Hello() {
	fmt.Println("hello")
}
