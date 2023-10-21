package bank_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/reduce/bank"
)

func TestBank(t *testing.T) {
	var (
		riya  = bank.NewAccount("Riya", 100)
		chris = bank.NewAccount("Chris", 75)
		adil  = bank.NewAccount("Adil", 200)
	)

	transactions := []bank.Transaction{
		bank.NewTransaction(chris, riya, 100),
		bank.NewTransaction(adil, chris, 25),
	}

	newBalanceFor := func(account bank.Account) float64 {
		return bank.NewBalanceFor(account, transactions).Balance
	}

	assertEqual(t, newBalanceFor(riya), 200)
	assertEqual(t, newBalanceFor(chris), 0)
	assertEqual(t, newBalanceFor(adil), 175)
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}
