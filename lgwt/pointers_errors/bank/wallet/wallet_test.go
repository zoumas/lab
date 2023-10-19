package wallet_test

import (
	"fmt"
	"testing"

	"github.com/zoumas/lab/lgwt/pointers_errors/bank/bitcoin"
	"github.com/zoumas/lab/lgwt/pointers_errors/bank/wallet"
)

func TestWallet_Deposit(t *testing.T) {
	w := wallet.Wallet{}

	w.Deposit(bitcoin.Bitcoin(10))

	want := bitcoin.Bitcoin(10)
	assertBalance(t, w, want)
}

func ExampleWallet_Deposit() {
	w := wallet.Wallet{}

	w.Deposit(10)

	fmt.Printf("%s\n", w.Balance())
	// Output: 10 BTC
}

func TestWallet_Withdraw(t *testing.T) {
	t.Run("with suffient funds", func(t *testing.T) {
		w := wallet.NewWallet(20)

		err := w.Withdraw(bitcoin.Bitcoin(10))
		assertNoError(t, err)

		want := bitcoin.Bitcoin(10)
		assertBalance(t, *w, want)
	})

	t.Run("with insufficient funds", func(t *testing.T) {
		initialBalance := bitcoin.Bitcoin(20)
		w := wallet.NewWallet(initialBalance)

		err := w.Withdraw(100)
		assertError(t, err, wallet.ErrInsufficientFunds)

		assertBalance(t, *w, initialBalance)
	})
}

func ExampleWallet_Withdraw() {
	w := wallet.NewWallet(20)

	w.Withdraw(10)
	fmt.Println(w.Balance())

	err := w.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 10 BTC
	// cannot withdraw, insufficient funds
}

func assertBalance(t testing.TB, w wallet.Wallet, want bitcoin.Bitcoin) {
	t.Helper()

	if got := w.Balance(); got != want {
		t.Errorf("\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("\nwanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("\nunexpected error:\n%q", err)
	}
}
