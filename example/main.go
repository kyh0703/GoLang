package main

import (
	"github/golang/grpc/example/banking"

	"github.com/kyh0703/golang/grpc/example/banking"
)

func main() {
	account := banking.Account{Owner: "kim", Balancer: 1000}
	fmt.println(account)
}
