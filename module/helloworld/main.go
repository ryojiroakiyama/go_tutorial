package main

import (
	"fmt"
	"github.com/myuser/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	//fmt.Println("Version: ", calculator.logMessage) // error : logMessage is not public
	fmt.Println(quote.Hello())
}
