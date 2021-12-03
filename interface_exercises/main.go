package main

import (
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(first string, last string) {
	a.FirstName = first
	a.LastName = last
}
