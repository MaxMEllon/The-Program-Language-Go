package bank_test

import (
	"github.com/maxmellon/The-Program-Language-Go/ch09/ex01"
	"testing"
)

func TestWithdraw(t *testing.T) {
	bank.Deposit(500)
	ok := bank.Withdraw(100)
	if !ok {
		t.Errorf("Result is %v, but %v", ok, true)
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Errorf("Result is %v, but %v", ok, true)
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Errorf("Result is %v, but %v", ok, true)
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Errorf("Result is %v, but %v", ok, true)
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Errorf("Result is %v, but %v", ok, true)
	}
	ok = bank.Withdraw(100)
	if ok {
		t.Errorf("Result is %v, but %v", ok, false)
	}
	if bank.Balance() != 0 {
		t.Errorf("Result is %d, but %d", bank.Balance(), 0)
	}
}
