package bank

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
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

func Reduce[A, B any](collection []A, acc func(B, A) B, identity B) B {
	result := identity
	for _, v := range collection {
		result = acc(result, v)
	}
	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(balance float64, t Transaction) float64 {
		switch name {
		case t.From:
			balance -= t.Sum
		case t.To:
			balance += t.Sum
		}
		return balance
	}

	return Reduce(transactions, adjustBalance, 0.0)
}
