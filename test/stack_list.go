package main

import "fmt"

type Node struct {
	data interface{}
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Empty() bool {
	if s.count == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(node *Node) {
	s.nodes = append(s.nodes, node)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	top := len(s.nodes) - 1
	data := s.nodes[top]
	s.nodes = s.nodes[:top]
	return data
}

func main() {
	s := NewStack()
	s.Push(&Node{data: 1})
	s.Push(&Node{data: 2})
	s.Push(&Node{data: 3})
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
