package wallet

import (
	"testing" 
	"fmt"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}	

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10) 

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	if got != want {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}