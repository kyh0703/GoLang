package main

type Node struct {
	data interface{}
}

type Queue struct {
	nodes []*Node
	count int
}

func (q *Queue) Empty() bool {
	return false
}

func (q *Queue) Push(node *Node) {

}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	return nil
}

func main() {

}
