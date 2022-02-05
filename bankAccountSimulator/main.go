package main

import (
	"fmt"
	"log"

	"github.com/traceur0/goProject/bankAccountSimulator/accounts"
)


func main() {
	account := accounts.NewAccount("Joe")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err) // log.fatalln(err)
	} // You must design the Error situation in Golang
	fmt.Println(account)
}