package main

import (
	"fmt"

	"github.com/kyh0703/golang/example/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	accounts.Deposit(10)
	fmt.Println(accounts.Balance())
}
