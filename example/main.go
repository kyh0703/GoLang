package main

import (
	"fmt"

	"github.com/kyh0703/golang/example/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
