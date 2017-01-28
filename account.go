package entry

type Account struct {
	id      int32
	name    string
	balance float32
	txs     []*Tx
}

func newAccount(name string, balance float32) Account {
	return Account{name: name, balance: balance, txs: make([]*Tx, 0)}
}

func (a *Account) credit(amount float32) {
	a.balance += amount
}

func (a *Account) debit(amount float32) {
	a.balance -= amount
}
