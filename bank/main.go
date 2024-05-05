package main

import (
	"bank/banking"
	"fmt"
)

func main() {
	account := banking.NewAccount("nico")

	fmt.Println(*account)

	account.Deposit(50)
	fmt.Println(account.Balance())

	err := account.Withdraw(60)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("sashin")
	fmt.Println(account.Balance(), account.Owner())

	fmt.Println(*account)
}
