package entry

import "time"

type Tx struct {
	id          int32
	date        time.Time
	fromAccount *Account
	toAccount   *Account
	description string
	amount      float32
}

func newTx(fromAccount *Account, toAccount *Account, description string, amount float32) Tx {
	result := Tx{}
	result.description = description
	result.fromAccount = fromAccount
	result.toAccount = toAccount
	result.amount = amount
	return result
}

func (t *Tx) execute() *Tx {
	t.fromAccount.debit(t.amount)
	t.toAccount.credit(t.amount)
	t.date = time.Now()

	// NB (from https://gobyexample.com/slices):
	// Note that we need to accept a return value from append as we may get a new slice value.
	s1 := append(t.fromAccount.txs, t)
	t.fromAccount.txs = s1
	s2 := append(t.toAccount.txs, t)
	t.toAccount.txs = s2

	return t
}
