package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

// type Stringer interface {
// 	String() string
// }

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if w.balance < value {
		return errors.New("insufficient balance")
	}
	w.balance -= value
	return nil
}
