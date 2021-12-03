package emp

import "fmt"

type Account struct {
	firstName string
	lastName  string
}

func (a *Account) ChangeName(first string, last string) {
	a.firstName = first
	a.lastName = last
}

type Employee struct {
	Account
	credit float64
}

func (e *Employee) AddCredits(add float64) {
	e.credit += add
}

func (e *Employee) RemoveCredits() {
	e.credit = 0
}

func (e Employee) CheckCredits() float64 {
	return e.credit
}

func (a Account) String() string {
	return fmt.Sprintf("%v %v", a.firstName, a.lastName)
}
