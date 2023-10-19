// package wallet provides types and functions for working with a client's Bitcoin wallet.
package wallet

import (
	"errors"

	"github.com/zoumas/lab/lgwt/pointers_errors/bank/bitcoin"
)

// Wallet represents a client's Bitcoin wallet.
type Wallet struct {
	balance bitcoin.Bitcoin
}

// Wallet Constructor. Sets a Wallet's initial balance.
func NewWallet(initialBalance bitcoin.Bitcoin) *Wallet {
	return &Wallet{
		balance: initialBalance,
	}
}

// Deposit adds amount to the Wallet's balance.
func (w *Wallet) Deposit(amount bitcoin.Bitcoin) {
	w.balance += amount
}

// Balance returns the wallet's balance.
func (w Wallet) Balance() bitcoin.Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw removes amount from the wallet's balance.
// Returns an error if the amount is greater than the wallet's current balance.
func (w *Wallet) Withdraw(amount bitcoin.Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
