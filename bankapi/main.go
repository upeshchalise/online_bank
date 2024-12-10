package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
	// fmt.Println(bank.Hello())
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found", number)
		} else {
			fmt.Fprintln(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if amountqs == "" {
		fmt.Fprintf(w, "Amount number is missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number")
	} else {
		account, ok := accounts[number]

		if !ok {
			fmt.Println(w, "Account with number %v can't be found", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Println(w, "%v", err)
			} else {
				fmt.Fprintln(w, account.Statement())
			}

		}

	}
}

func withdraw(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if amountqs == "" {
		fmt.Fprintf(w, "Amount number is missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number")
	} else {
		account, ok := accounts[number]

		if !ok {
			fmt.Println(w, "Account with number %v can't be found", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Println(w, "%v", err)
			} else {
				fmt.Fprintln(w, account.Statement())
			}
		}
	}

}
