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
		wallet.Widthdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{
			balance: startingBalance,
		}

		wallet.Deposit(Bitcoin(20))

		err := wallet.Widthdraw(Bitcoin(100))

		if err == nil {
			t.Errorf("Wanted an error but no error present")
		}
	})


	
}