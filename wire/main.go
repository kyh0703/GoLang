package main

import "fmt"

// interface
type Fooer interface {
    Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
    return string(*b)
}

// proivider
func provideMyFooer() *MyFooer {
    b := new(MyFooer)
    *b = "Hello, World!"
    return b
}

type Bar string

func provideBar(f Fooer) string {
    // f will be a *MyFooer.
    return f.Foo()
}

func main() {
	fmt.Println("test")
}