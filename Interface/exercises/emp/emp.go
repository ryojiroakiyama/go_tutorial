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

// answer example
//package store

//import (
//    "errors"
//    "fmt"
//)

//type Account struct {
//    FirstName string
//    LastName  string
//}

//type Employee struct {
//    Account
//    Credits float64
//}

//func (a *Account) ChangeName(newname string) {
//    a.FirstName = newname
//}

//func (e Employee) String() string {
//    return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
//}

//func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
//    return &Employee{Account{firstName, lastName}, credits}, nil
//}

//func (e *Employee) AddCredits(amount float64) (float64, error) {
//    if amount > 0.0 {
//        e.Credits += amount
//        return e.Credits, nil
//    }
//    return 0.0, errors.New("Invalid credit amount.")
//}

//func (e *Employee) RemoveCredits(amount float64) (float64, error) {
//    if amount > 0.0 {
//        if amount <= e.Credits {
//            e.Credits -= amount
//            return e.Credits, nil
//        }
//        return 0.0, errors.New("You can't remove more credits than the account has.")
//    }
//    return 0.0, errors.New("You can't remove negative numbers.")
//}

//func (e *Employee) CheckCredits() float64 {
//    return e.Credits
//}
