package main

import (
<<<<<<< HEAD
	"github.com/kyh0703/golang/example/mydict"
=======
	"fmt"

	mydict "github.com/kyh0703/golang/example/mydict"
>>>>>>> 753cd406a8b860cf1d1ec9777313b5132652b062
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
<<<<<<< HEAD
	err := dictionary.Update(baseWord, "Second")
	if err != nil {

	}
=======
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
>>>>>>> 753cd406a8b860cf1d1ec9777313b5132652b062
}
