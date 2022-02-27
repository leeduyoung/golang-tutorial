package main

import (
	"fmt"
	"sample/learngo/banking/accounts"
)

func main() {
	account := accounts.NewAccount("kaye")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(account.Balance())

}
