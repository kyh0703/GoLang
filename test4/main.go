package main

import (
	"fmt"
	"time"
)

type Queue[T any] struct {
	value *T
}

func (q *Queue[T]) Pop() T {
	arr := *q.value
	nArr := arr[1:]
	q.value = &nArr
	return arr[0]
}

func main() {
	timeT := time.Duration()
	fmt.Println(timeT)
}
