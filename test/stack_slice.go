package main

import "fmt"

type stack []interface{}

func (s *stack) empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *stack) push(data interface{}) {
	*s = append(*s, data)
}

func (s *stack) pop() interface{} {
	if s.empty() {
		return nil
	}

	top := len(*s) - 1
	data := (*s)[top]
	*s = (*s)[:top]
	return data
}

func main() {
	var s stack
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
