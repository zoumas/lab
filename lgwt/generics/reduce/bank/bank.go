package bank

func Reduce[A, B any](f func(x B, y A) B, xs []A, identity B) B {
	a := identity
	for _, x := range xs {
		a = f(a, x)
	}
	return a
}

type Account struct {
	Name    string
	Balance float64
}

func NewAccount(name string, balance float64) Account {
	return Account{name, balance}
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  sum,
	}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(applyTransaction, transactions, account)
}

func applyTransaction(a Account, t Transaction) Account {
	switch a.Name {
	case t.From:
		a.Balance -= t.Sum
	case t.To:
		a.Balance += t.Sum
	}
	return a
}
