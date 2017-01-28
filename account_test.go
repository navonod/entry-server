package entry

import "testing"

func TestAccountCredit(t *testing.T) {
	var a Account
	a.credit(1.25)
	if a.balance != 1.25 {
		t.Errorf("Account.credit test failed. Expected %f, found %f", 1.25, a.balance)
	}
	a.credit(3)
	if a.balance != 4.25 {
		t.Errorf("Account.credit twice test failed. Expected %f, found %f", 4.25, a.balance)
	}
}

func TestAccountDebit(t *testing.T) {
	var a Account
	a.balance = 5.25
	a.debit(3.25)
	if a.balance != 2 {
		t.Errorf("Account.debit test failed. Expected %f, found %f", 2.0, a.balance)
	}
}
