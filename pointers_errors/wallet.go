package pointers_errors

import (
	"errors"
	"fmt"
)
var ErrInsufficientFunds string = "cannot withdraw, insufficient funds"

type Bitcoin int 

type Wallet struct {
	balance Bitcoin
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}