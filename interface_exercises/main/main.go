package main

import (
	"fmt"
	"github.com/myuser/emp"
)

func main() {
	//e := emp.Employee{emp.Account{"Ryojiro", "Akiyama"}, 0}
	e := emp.Employee{}
	fmt.Println("Account:", e)
	e.ChangeName("Jiro", "MAAAAA")
	fmt.Println("Account:", e)
	fmt.Println("Credits:", e.CheckCredits())
	e.AddCredits(50)
	fmt.Println("Credits:", e.CheckCredits())
	e.RemoveCredits()
	fmt.Println("Credits:", e.CheckCredits())

}
