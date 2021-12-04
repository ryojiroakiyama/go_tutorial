package bank

import "testing"

func TestAccount(t *testing.T) {
	accout := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  10001,
		Balance: 0,
	}

	if accout.Name == "" {
		t.Error("can't create an Account object")
	}
}
