package main

import (
	"fmt"

	"github.com/sashin92/goBasic/myBank/accounts"
)

func main() {
	account := accounts.NewAccount("sashin")
	account.Deposit(50)
	fmt.Println(account)
	fmt.Println("==============")
	account.Withdraw(30)
	fmt.Println(account)
	fmt.Println("==============")
	err := account.Withdraw(30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("==============")
	fmt.Println(account)
}
