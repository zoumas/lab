package bank_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/arrays_slices/bank"
	"github.com/zoumas/lab/lgwt/generics/assert"
)

func TestBank(t *testing.T) {
	var (
		riya  = bank.Account{Name: "Riya", Balance: 100}
		chris = bank.Account{Name: "Chris", Balance: 75}
		adil  = bank.Account{Name: "Adil", Balance: 200}

		transactions = []bank.Transaction{
			bank.NewTransaction(chris, riya, 100),
			bank.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account bank.Account) float64 {
		return bank.NewBalanceFor(account, transactions).Balance
	}

	assert.Equal(t, newBalanceFor(riya), 200)
	assert.Equal(t, newBalanceFor(chris), 0)
	assert.Equal(t, newBalanceFor(adil), 175)
}
