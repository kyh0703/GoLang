package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
