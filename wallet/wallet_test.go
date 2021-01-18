package wallet

import (
	"testing" 
	_"fmt"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet  Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Wanted %s, got %s", want, got)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}	

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("widthdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		got := wallet.Widthdraw(Bitcoin(10))
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Wanted %s, got %s", want, got)
		}
		assertBalance(t, wallet, Bitcoin(10))
	})


	
}