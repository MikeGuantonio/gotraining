package wallet

import (
	"fmt"
	"errors"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Widthdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("Insufficent Funds")
	}
	w.balance -= amount
	return nil
}