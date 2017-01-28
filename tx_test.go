package entry

import (
	"flag"
	"os"
	"testing"
)

var accountFrom *Account
var accountTo *Account

func TestMain(m *testing.M) {
	flag.Parse()
	// NB assigning accountFrom = newAccount("from", 350.0) fails, as the returned value has no address
	tmpFrom := newAccount("from", 350.0)
	accountFrom = &tmpFrom
	tmpTo := newAccount("to", 700.0)
	accountTo = &tmpTo
	os.Exit(m.Run())
}

func TestTxCreation(t *testing.T) {
	tx := newTx(accountFrom, accountTo, "test Tx", 350.0)
	if tx.amount != 350.0 {
		t.Errorf("Failed Tx Creation test. Expected amount = %f, found %f", 350.0, tx.amount)
	}
}

func TestTxExecution(t *testing.T) {
	tx := newTx(accountFrom, accountTo, "test Tx", 350.0)
	tx.execute()
	if accountFrom.balance != 0.0 {
		t.Errorf("Failed Tx Exectution test. Expected From Account Balance to equal %f, found %f", 0.0, accountFrom.balance)
	}
	if accountTo.balance != 1050.0 {
		t.Errorf("Failed Tx Exectution test. Expected To Account Balance to equal %f, found %f", 1050.0, accountTo.balance)
	}
	if len(accountFrom.txs) != 1 {
		t.Errorf("Failed Tx Exectution test. Expected From Account txs length to equal %d, found %d", 1, len(accountFrom.txs))
	}
	if len(accountTo.txs) != 1 {
		t.Errorf("Failed Tx Exectution test. Expected To Account txs length to equal %d, found %d", 1, len(accountFrom.txs))
	}
	if accountFrom.txs[0] != accountTo.txs[0] {
		t.Errorf("Failed Tx Exectution test. Expected To Account tx & From Account tx to share memory address.")
	}
}
