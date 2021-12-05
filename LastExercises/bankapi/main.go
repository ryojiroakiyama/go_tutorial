package main

import (
	"encoding/json"
	"fmt"
	bank "go_tutorial/LastExercises/bankcore"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, exist := accounts[number]
		if !exist {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			//fmt.Fprintf(w, bank.Statement(account))
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, exist := accounts[number]
		if !exist {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, exist := accounts[number]
		if !exist {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func send(w http.ResponseWriter, req *http.Request) {
	sourcenumberqs := req.URL.Query().Get("sourcenumber")
	remitternumberqs := req.URL.Query().Get("remitternumber")
	amountqs := req.URL.Query().Get("amount")

	if sourcenumberqs == "" || remitternumberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if sourcenumber, err := strconv.ParseFloat(sourcenumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid source account number!")
	} else if remitternumber, err := strconv.ParseFloat(remitternumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid remitter account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		if sourceaccount, exist := accounts[sourcenumber]; !exist {
			fmt.Fprintf(w, "Source account with number %v can't be found!", sourcenumber)
		} else if remitteraccount, exist := accounts[remitternumber]; !exist {
			fmt.Fprintf(w, "Remitter account with number %v can't be found!", remitternumber)
		} else {
			err := sourceaccount.Send(amount, remitteraccount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, sourceaccount.Statement())
				fmt.Println()
				fmt.Fprintf(w, remitteraccount.Statement())
			}
		}
	}
}

func testnormal() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}
	accounts[1002] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Ben",
			Address: "Zushi, Kanagawa",
			Phone:   "(222) 555 5555",
		},
		Number: 1002,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/send", send)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type Custom struct {
	bank.Account
}

func (c *Custom) Statement() string {
	j, _ := json.Marshal(c.Account)
	return string(j)
}

func test_use_another_statement() {
	accounts[1001] = &Custom{
		bank.Account{
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}
	//accounts[1002] = &bank.Account{
	//	Customer: bank.Customer{
	//		Name:    "Ben",
	//		Address: "Zushi, Kanagawa",
	//		Phone:   "(222) 555 5555",
	//	},
	//	Number: 1002,
	//}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	//http.HandleFunc("/send", send)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func main() {
	//testnormal()
	test_use_another_statement()
}
