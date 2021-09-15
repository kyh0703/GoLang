package main

import "fmt"

type queue []interface{}

func (s *queue) empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *queue) push(data interface{}) {
	*s = append(*s, data)
}

func (s *queue) pop() interface{} {
	if s.empty() {
		return nil
	}
	data := (*s)[0]
	*s = (*s)[1:]
	return data
}

func main() {
	var s queue
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
