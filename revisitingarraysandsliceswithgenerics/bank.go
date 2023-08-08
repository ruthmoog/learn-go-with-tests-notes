package main

type Transaction struct {
	From string
	Sum  int
	To   string
}

type Account struct {
	Name    string
	Balance int
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(account Account, transaction Transaction) Account {
	if transaction.From == account.Name {
		account.Balance -= transaction.Sum
	}
	if transaction.To == account.Name {
		account.Balance += transaction.Sum
	}
	return account
}

func NewTransaction(from, to Account, amount int) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: amount}
}
