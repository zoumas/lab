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
	fmt.Println("Initial balance:", w.Balance())

	w.Deposit(10)

	fmt.Println("Balance after depositing 10 BTC:", w.Balance())
	// Output:
	// Initial balance: 0 BTC
	// Balance after depositing 10 BTC: 10 BTC
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
	fmt.Println("Initial balance:", w.Balance())

	w.Withdraw(10)
	fmt.Println("Withdrawing 10 BTC leaves the wallet with:", w.Balance())

	err := w.Withdraw(20)
	if err != nil {
		fmt.Println("Trying to withdraw 20 BTC:", err)
	}
	// Output:
	// Initial balance: 20 BTC
	// Withdrawing 10 BTC leaves the wallet with: 10 BTC
	// Trying to withdraw 20 BTC: cannot withdraw, insufficient funds
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
